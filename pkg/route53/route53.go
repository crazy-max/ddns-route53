package route53

import (
	"context"
	"net"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awsr53 "github.com/aws/aws-sdk-go-v2/service/route53"
	awsr53types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/pkg/errors"
)

type Client struct {
	ctx          context.Context
	client       *awsr53.Client
	hostedZoneID string
}

// New initializes a new route53 client
func New(ctx context.Context, accessKey, secretKey, hostedZoneID string, maxRetries int, maxBackoffDelay time.Duration) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		// Route53 uses a global endpoint and route53domains
		// currently only has a single regional endpoint in us-east-1
		// http://docs.aws.amazon.com/general/latest/gr/rande.html#r53_region
		config.WithRegion("us-east-1"),
		config.WithRetryer(func() aws.Retryer {
			r := retry.AddWithMaxAttempts(retry.NewStandard(), maxRetries)
			r = retry.AddWithMaxBackoffDelay(r, maxBackoffDelay)
			return r
		}),
	)
	if err != nil {
		return nil, err
	}
	if len(accessKey) > 0 && len(secretKey) > 0 {
		cfg.Credentials = credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")
	}
	return &Client{
		ctx:          ctx,
		client:       awsr53.NewFromConfig(cfg),
		hostedZoneID: hostedZoneID,
	}, nil
}

func (r *Client) Update(changes []awsr53types.Change, comment string) (*awsr53.ChangeResourceRecordSetsOutput, error) {
	return r.client.ChangeResourceRecordSets(r.ctx, &awsr53.ChangeResourceRecordSetsInput{
		ChangeBatch: &awsr53types.ChangeBatch{
			Comment: aws.String(comment),
			Changes: changes,
		},
		HostedZoneId: aws.String(r.hostedZoneID),
	})
}

func (r *Client) RecordIP(records []awsr53types.ResourceRecordSet, name *string, recordType awsr53types.RRType) (net.IP, error) {
	var record awsr53types.ResourceRecordSet
	for _, rec := range records {
		if rec.Type != recordType || *rec.Name != *name {
			continue
		}
		record = rec
		break
	}
	if len(record.ResourceRecords) == 0 {
		return nil, nil
	}
	return net.ParseIP(*record.ResourceRecords[0].Value), nil
}

func (r *Client) ListRecords() ([]awsr53types.ResourceRecordSet, error) {
	var records []awsr53types.ResourceRecordSet
	req := &awsr53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(r.hostedZoneID),
	}
	for {
		resp, err := r.client.ListResourceRecordSets(r.ctx, req)
		if err != nil {
			return nil, errors.Wrap(err, "failed to fetch records")
		}
		for _, set := range resp.ResourceRecordSets {
			records = append(records, set)
		}
		if !resp.IsTruncated {
			break
		}
		req.StartRecordIdentifier = resp.NextRecordIdentifier
		req.StartRecordName = resp.NextRecordName
		req.StartRecordType = resp.NextRecordType
	}
	return records, nil
}
