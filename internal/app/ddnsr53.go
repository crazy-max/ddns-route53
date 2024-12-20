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
	"github.com/golang-module/carbon/v2"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

// DDNSRoute53 represents an active ddns-route53 object
type DDNSRoute53 struct {
	meta   model.Meta
	cfg    *config.Config
	cron   *cron.Cron
	r53    *route53.Client
	wip    *wanip.Client
	jobID  cron.EntryID
	locker uint32
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

	r53, err := route53.New(context.TODO(), accessKeyID, secretAccessKey, cfg.Route53.HostedZoneID, cfg.Cli.MaxRetries, cfg.Cli.MaxBackoffDelay)
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
		wip: wanip.New(
			wanip.WithInterfaceName(cfg.Cli.Ifname),
			wanip.WithUserAgent(meta.UserAgent),
			wanip.WithMaxRetries(cfg.Cli.MaxRetries),
		),
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
		carbon.CreateFromStdTime(c.cron.Entry(c.jobID).Next).DiffAbsInString(),
		c.cron.Entry(c.jobID).Next)

	select {}
}

// Run runs ddns-route53 process
func (c *DDNSRoute53) Run() {
	var wanIPv4, wanIPv6 net.IP

	if !atomic.CompareAndSwapUint32(&c.locker, 0, 1) {
		log.Warn().Msg("Already running")
		return
	}
	defer atomic.StoreUint32(&c.locker, 0)
	if c.jobID > 0 {
		defer log.Info().Msgf("Next run in %s (%s)",
			carbon.CreateFromStdTime(c.cron.Entry(c.jobID).Next).DiffAbsInString(),
			c.cron.Entry(c.jobID).Next)
	}

	if *c.cfg.Route53.HandleIPv4 {
		var wanErrs wanip.Errors
		wanIPv4, wanErrs = c.wip.IPv4()
		wanLogger := log.Error()
		if wanIPv4 != nil {
			wanLogger = log.Debug()
			log.Info().Msgf("Current WAN IPv4: %s", wanIPv4)
		}
		for _, wanIPErr := range wanErrs {
			wanLogger.Err(wanIPErr.Err).Str("provider-url", wanIPErr.ProviderURL).Msg("Cannot retrieve WAN IPv4 address")
		}
	}

	if *c.cfg.Route53.HandleIPv6 {
		var wanErrs wanip.Errors
		wanIPv6, wanErrs = c.wip.IPv6()
		wanLogger := log.Error()
		if wanIPv6 != nil {
			wanLogger = log.Debug()
			log.Info().Msgf("Current WAN IPv6: %s", wanIPv6)
		}
		for _, wanIPErr := range wanErrs {
			wanLogger.Err(wanIPErr.Err).Str("provider-url", wanIPErr.ProviderURL).Msg("Cannot retrieve WAN IPv6 address")
		}
	}

	if wanIPv4 == nil && wanIPv6 == nil {
		return
	}

	records, err := c.r53.ListRecords()
	if err != nil {
		log.Warn().Err(err).Msg("Cannot list records")
	}

	var r53Changes []awsr53types.Change
	for _, rs := range c.cfg.Route53.RecordsSet {
		recordValue := new(string)
		if rs.Type == awsr53types.RRTypeA {
			if wanIPv4 == nil {
				log.Error().Msgf("No WAN IPv4 address available to update %s record set", rs.Name)
				continue
			}
			if recordCurrentIP, err := c.r53.RecordIP(records, aws.String(rs.Name), rs.Type); err == nil && wanIPv4.Equal(recordCurrentIP) {
				log.Info().Msgf("WAN IPv4 has not changed for %s record set", rs.Name)
				continue
			}
			recordValue = aws.String(wanIPv4.String())
		} else if rs.Type == awsr53types.RRTypeAaaa {
			if wanIPv6 == nil {
				log.Error().Msgf("No WAN IPv6 address available to update %s record set", rs.Name)
				continue
			}
			if recordCurrentIP, err := c.r53.RecordIP(records, aws.String(rs.Name), rs.Type); err == nil && wanIPv6.Equal(recordCurrentIP) {
				log.Info().Msgf("WAN IPv6 has not changed for %s record set", rs.Name)
				continue
			}
			recordValue = aws.String(wanIPv6.String())
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
	resp, err := c.r53.Update(r53Changes, fmt.Sprintf("Updated by %s %s at %s",
		c.meta.Name,
		c.meta.Version,
		time.Now().Format("2006-01-02 15:04:05"),
	))
	if err != nil {
		log.Error().Err(err).Msg("Cannot update records set")
		return
	}

	log.Info().Interface("changes", resp).Msgf("%d record(s) set updated", len(r53Changes))
}

// Close closes ddns-route53
func (c *DDNSRoute53) Close() {
	if c.cron != nil {
		c.cron.Stop()
	}
}
