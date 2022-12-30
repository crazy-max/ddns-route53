package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awsr53 "github.com/aws/aws-sdk-go-v2/service/route53"
	awsr53types "github.com/aws/aws-sdk-go-v2/service/route53/types"
)

type Client struct {
	ctx          context.Context
	client       *awsr53.Client
	hostedZoneID string
}

// New initializes a new route53 client
func New(ctx context.Context, accessKey, secretKey, hostedZoneID string) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		// Route53 uses a global endpoint and route53domains
		// currently only has a single regional endpoint in us-east-1
		// http://docs.aws.amazon.com/general/latest/gr/rande.html#r53_region
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		return nil, err
	}
	return &Client{
		ctx:          ctx,
		client:       awsr53.NewFromConfig(cfg),
		hostedZoneID: hostedZoneID,
	}, nil
}

func (r *Client) Update(ctx context.Context, changes []awsr53types.Change, comment string) (*awsr53.ChangeResourceRecordSetsOutput, error) {
	return r.client.ChangeResourceRecordSets(ctx, &awsr53.ChangeResourceRecordSetsInput{
		ChangeBatch: &awsr53types.ChangeBatch{
			Comment: aws.String(comment),
			Changes: changes,
		},
		HostedZoneId: aws.String(r.hostedZoneID),
	})
}
