package app

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/crazy-max/ddns-route53/v2/internal/config"
	"github.com/crazy-max/ddns-route53/v2/internal/model"
	"github.com/crazy-max/ddns-route53/v2/pkg/identme"
	"github.com/crazy-max/ddns-route53/v2/pkg/utl"
	"github.com/hako/durafmt"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

// DDNSRoute53 represents an active ddns-route53 object
type DDNSRoute53 struct {
	meta     model.Meta
	cfg      *config.Config
	cron     *cron.Cron
	r53      *route53.Route53
	im       *identme.Client
	jobID    cron.EntryID
	lastIPv4 net.IP
	lastIPv6 net.IP
	locker   uint32
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

	creds := credentials.NewChainCredentials(
		[]credentials.Provider{
			&credentials.EnvProvider{},
			&credentials.StaticProvider{
				Value: credentials.Value{
					AccessKeyID:     accessKeyID,
					SecretAccessKey: secretAccessKey,
					SessionToken:    "",
				}},
		},
	)

	// AWS SDK session
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	return &DDNSRoute53{
		meta: meta,
		cfg: cfg,
		cron: cron.New(cron.WithParser(cron.NewParser(
			cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor),
		)),
		r53: route53.New(sess, &aws.Config{Credentials: creds}),
		im: identme.NewClient(
			meta.UserAgent,
			cfg.Cli.MaxRetries,
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
		durafmt.ParseShort(c.cron.Entry(c.jobID).Next.Sub(time.Now())).String(),
		c.cron.Entry(c.jobID).Next)

	select {}
}

// Run runs ddns-route53 process
func (c *DDNSRoute53) Run() {
	var err error
	if !atomic.CompareAndSwapUint32(&c.locker, 0, 1) {
		log.Warn().Msg("Already running")
		return
	}
	defer atomic.StoreUint32(&c.locker, 0)
	if c.jobID > 0 {
		defer log.Info().Msgf("Next run in %s (%s)",
			durafmt.ParseShort(c.cron.Entry(c.jobID).Next.Sub(time.Now())).String(),
			c.cron.Entry(c.jobID).Next)
	}

	var wanIPv4 net.IP
	if *c.cfg.Route53.HandleIPv4 {
		wanIPv4, err = c.im.IPv4()
		if err != nil {
			log.Error().Err(err).Msg("Cannot retrieve WAN IPv4 address")
		} else {
			log.Info().Msgf("Current WAN IPv4: %s", wanIPv4)
		}
	}

	var wanIPv6 net.IP
	if *c.cfg.Route53.HandleIPv6 {
		wanIPv6, err = c.im.IPv6()
		if err != nil {
			log.Error().Err(err).Msg("Cannot retrieve WAN IPv6 address")
		} else {
			log.Info().Msgf("Current WAN IPv6: %s", wanIPv6)
		}
	}

	if wanIPv4 == nil && wanIPv6 == nil {
		return
	}

	// Skip if last IP is identical or empty
	if wanIPv4.Equal(c.lastIPv4) && wanIPv6.Equal(c.lastIPv6) {
		log.Info().Msg("WAN IPv4/IPv6 addresses have not changed since last update. Skipping...")
		return
	}

	// Create Route53 changes
	var r53Changes []*route53.Change
	for _, rs := range c.cfg.Route53.RecordsSet {
		if wanIPv4 == nil && rs.Type == route53.RRTypeA {
			log.Error().Msgf("No WAN IPv4 address available to update %s record", rs.Name)
			continue
		} else if wanIPv6 == nil && rs.Type == route53.RRTypeAaaa {
			log.Error().Msgf("No WAN IPv6 address available to update %s record", rs.Name)
			continue
		}
		recordValue := aws.String(wanIPv4.String())
		if rs.Type == route53.RRTypeAaaa {
			recordValue = aws.String(wanIPv6.String())
		}
		r53Changes = append(r53Changes, &route53.Change{
			Action: aws.String("UPSERT"),
			ResourceRecordSet: &route53.ResourceRecordSet{
				Name:            aws.String(rs.Name),
				Type:            aws.String(rs.Type),
				TTL:             aws.Int64(rs.TTL),
				ResourceRecords: []*route53.ResourceRecord{{Value: recordValue}},
			},
		})
	}

	// Check changes
	if len(r53Changes) == 0 {
		log.Warn().Msgf("No Route53 record set to update. Skipping...")
		return
	}

	// Create resource records
	resRS := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Comment: aws.String(fmt.Sprintf("Updated by %s %s at %s",
				c.meta.Name,
				c.meta.Version,
				time.Now().In(c.loc).Format("2006-01-02 15:04:05"),
			)),
			Changes: r53Changes,
		},
		HostedZoneId: aws.String(c.cfg.Route53.HostedZoneID),
	}

	// Update records
	resp, err := c.r53.ChangeResourceRecordSets(resRS)
	if err != nil {
		log.Error().Err(err).Msg("Cannot update records set")
	}
	log.Info().Interface("changes", resp).Msgf("%d records set updated", len(r53Changes))

	// Update last IPv4/IPv6
	c.lastIPv4 = wanIPv4
	c.lastIPv6 = wanIPv6
}

// Close closes ddns-route53
func (c *DDNSRoute53) Close() {
	if c.cron != nil {
		c.cron.Stop()
	}
}
