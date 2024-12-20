// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of location objects and their CIDR blocks.
func (c *Client) ListCidrBlocks(ctx context.Context, params *ListCidrBlocksInput, optFns ...func(*Options)) (*ListCidrBlocksOutput, error) {
	if params == nil {
		params = &ListCidrBlocksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCidrBlocks", params, optFns, c.addOperationListCidrBlocksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCidrBlocksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCidrBlocksInput struct {

	// The UUID of the CIDR collection.
	//
	// This member is required.
	CollectionId *string

	// The name of the CIDR collection location.
	LocationName *string

	// Maximum number of results you want returned.
	MaxResults *int32

	// An opaque pagination token to indicate where the service is to begin
	// enumerating results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCidrBlocksOutput struct {

	// A complex type that contains information about the CIDR blocks.
	CidrBlocks []types.CidrBlockSummary

	// An opaque pagination token to indicate where the service is to begin
	// enumerating results. If no value is provided, the listing of results starts from
	// the beginning.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCidrBlocksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListCidrBlocks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListCidrBlocks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCidrBlocks"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListCidrBlocksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCidrBlocks(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListCidrBlocksAPIClient is a client that implements the ListCidrBlocks
// operation.
type ListCidrBlocksAPIClient interface {
	ListCidrBlocks(context.Context, *ListCidrBlocksInput, ...func(*Options)) (*ListCidrBlocksOutput, error)
}

var _ ListCidrBlocksAPIClient = (*Client)(nil)

// ListCidrBlocksPaginatorOptions is the paginator options for ListCidrBlocks
type ListCidrBlocksPaginatorOptions struct {
	// Maximum number of results you want returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCidrBlocksPaginator is a paginator for ListCidrBlocks
type ListCidrBlocksPaginator struct {
	options   ListCidrBlocksPaginatorOptions
	client    ListCidrBlocksAPIClient
	params    *ListCidrBlocksInput
	nextToken *string
	firstPage bool
}

// NewListCidrBlocksPaginator returns a new ListCidrBlocksPaginator
func NewListCidrBlocksPaginator(client ListCidrBlocksAPIClient, params *ListCidrBlocksInput, optFns ...func(*ListCidrBlocksPaginatorOptions)) *ListCidrBlocksPaginator {
	if params == nil {
		params = &ListCidrBlocksInput{}
	}

	options := ListCidrBlocksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCidrBlocksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCidrBlocksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCidrBlocks page.
func (p *ListCidrBlocksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCidrBlocksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListCidrBlocks(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListCidrBlocks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCidrBlocks",
	}
}
