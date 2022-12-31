package app

import (
	"context"
	"fmt"
	"net"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsr53types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/crazy-max/ddns-route53/v2/internal/config"
	"github.com/crazy-max/ddns-route53/v2/internal/model"
	"github.com/crazy-max/ddns-route53/v2/pkg/route53"
	"github.com/crazy-max/ddns-route53/v2/pkg/utl"
	"github.com/crazy-max/ddns-route53/v2/pkg/wanip"
	"github.com/hako/durafmt"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

// DDNSRoute53 represents an active ddns-route53 object
type DDNSRoute53 struct {
	meta      model.Meta
	cfg       *config.Config
	cron      *cron.Cron
	r53       *route53.Client
	im        *wanip.Client
	jobID     cron.EntryID
	currentIP ip
	locker    uint32
}

type ip struct {
	V4 net.IP
	V6 net.IP
}

// New creates new ddns-route53 instance
func New(meta model.Meta, cfg *config.Config) (*DDNSRoute53, error) {
	var err error
	var accessKeyID string
	var secretAccessKey string

	if cfg.Credentials != nil {
		accessKeyID, err = utl.GetSecret(cfg.Credentials.AccessKeyID, cfg.Credentials.AccessKeyIDFile)
		if err != nil {
			log.Warn().Err(err).Msg("Cannot retrieve access key ID")
		}
		secretAccessKey, err = utl.GetSecret(cfg.Credentials.SecretAccessKey, cfg.Credentials.SecretAccessKeyFile)
		if err != nil {
			log.Warn().Err(err).Msg("Cannot retrieve secret access key")
		}
	}

	r53, err := route53.New(context.TODO(), accessKeyID, secretAccessKey, cfg.Route53.HostedZoneID)
	if err != nil {
		return nil, err
	}

	return &DDNSRoute53{
		meta: meta,
		cfg:  cfg,
		cron: cron.New(cron.WithParser(cron.NewParser(
			cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor),
		)),
		r53: r53,
		im:  wanip.NewClient(meta.UserAgent, cfg.Cli.MaxRetries),
	}, nil
}

// Start starts ddns-route53
func (c *DDNSRoute53) Start() error {
	var err error

	// Run on startup
	c.Run()

	// Check scheduler enabled
	if c.cfg.Cli.Schedule == "" {
		return nil
	}

	// Init scheduler
	c.jobID, err = c.cron.AddJob(c.cfg.Cli.Schedule, c)
	if err != nil {
		return err
	}
	log.Info().Msgf("Cron initialized with schedule %s", c.cfg.Cli.Schedule)

	// Start scheduler
	c.cron.Start()
	log.Info().Msgf("Next run in %s (%s)",
		durafmt.Parse(time.Until(c.cron.Entry(c.jobID).Next)).LimitFirstN(2).String(),
		c.cron.Entry(c.jobID).Next)

	select {}
}

// Run runs ddns-route53 process
func (c *DDNSRoute53) Run() {
	var err error
	var wanIP ip

	if !atomic.CompareAndSwapUint32(&c.locker, 0, 1) {
		log.Warn().Msg("Already running")
		return
	}
	defer atomic.StoreUint32(&c.locker, 0)
	if c.jobID > 0 {
		defer log.Info().Msgf("Next run in %s (%s)",
			durafmt.Parse(time.Until(c.cron.Entry(c.jobID).Next)).LimitFirstN(2).String(),
			c.cron.Entry(c.jobID).Next)
	}

	if *c.cfg.Route53.HandleIPv4 {
		var wanErrs wanip.Errors
		wanIP.V4, wanErrs = c.im.IPv4()
		wanLogger := log.Error()
		if wanIP.V4 != nil {
			wanLogger = log.Debug()
			log.Info().Msgf("Current WAN IPv4: %s", wanIP.V4)
		}
		if wanErrs != nil {
			for _, wanIPErr := range wanErrs {
				wanLogger.Err(wanIPErr.Err).Str("provider-url", wanIPErr.ProviderURL).Msg("Cannot retrieve WAN IPv4 address")
			}
		}
	}

	if *c.cfg.Route53.HandleIPv6 {
		var wanErrs wanip.Errors
		wanIP.V6, wanErrs = c.im.IPv6()
		wanLogger := log.Error()
		if wanIP.V4 != nil {
			wanLogger = log.Debug()
			log.Info().Msgf("Current WAN IPv6: %s", wanIP.V6)
		}
		if wanErrs != nil {
			for _, wanIPErr := range wanErrs {
				wanLogger.Err(wanIPErr.Err).Str("provider-url", wanIPErr.ProviderURL).Msg("Cannot retrieve WAN IPv6 address")
			}
		}
	}

	if wanIP.V4 == nil && wanIP.V6 == nil {
		return
	}

	// Skip if current IP is identical or empty
	if wanIP.V4.Equal(c.currentIP.V4) && wanIP.V6.Equal(c.currentIP.V6) {
		log.Info().Msg("WAN IPv4/IPv6 addresses have not changed since last update. Skipping...")
		return
	}

	// Create Route53 changes
	var r53Changes []awsr53types.Change
	for _, rs := range c.cfg.Route53.RecordsSet {
		if wanIP.V4 == nil && rs.Type == awsr53types.RRTypeA {
			log.Error().Msgf("No WAN IPv4 address available to update %s record", rs.Name)
			continue
		} else if wanIP.V6 == nil && rs.Type == awsr53types.RRTypeAaaa {
			log.Error().Msgf("No WAN IPv6 address available to update %s record", rs.Name)
			continue
		}
		recordValue := aws.String(wanIP.V4.String())
		if rs.Type == awsr53types.RRTypeAaaa {
			recordValue = aws.String(wanIP.V6.String())
		}
		r53Changes = append(r53Changes, awsr53types.Change{
			Action: awsr53types.ChangeActionUpsert,
			ResourceRecordSet: &awsr53types.ResourceRecordSet{
				Name:            aws.String(rs.Name),
				Type:            rs.Type,
				TTL:             aws.Int64(rs.TTL),
				ResourceRecords: []awsr53types.ResourceRecord{{Value: recordValue}},
			},
		})
	}

	if len(r53Changes) == 0 {
		log.Warn().Msgf("No Route53 record set to update. Skipping...")
		return
	}

	// Update Route53 records set
	resp, err := c.r53.Update(context.Background(), r53Changes, fmt.Sprintf("Updated by %s %s at %s",
		c.meta.Name,
		c.meta.Version,
		time.Now().Format("2006-01-02 15:04:05"),
	))
	if err != nil {
		log.Error().Err(err).Msg("Cannot update records set")
		return
	}

	log.Info().Interface("changes", resp).Msgf("%d records set updated", len(r53Changes))
	c.currentIP = wanIP
}

// Close closes ddns-route53
func (c *DDNSRoute53) Close() {
	if c.cron != nil {
		c.cron.Stop()
	}
}
