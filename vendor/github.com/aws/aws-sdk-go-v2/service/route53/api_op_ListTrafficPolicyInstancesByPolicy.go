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

// Gets information about the traffic policy instances that you created by using a
// specify traffic policy version.
//
// After you submit a CreateTrafficPolicyInstance or an UpdateTrafficPolicyInstance
// request, there's a brief delay while Amazon Route 53 creates the resource record
// sets that are specified in the traffic policy definition. For more information,
// see the State response element.
//
// Route 53 returns a maximum of 100 items in each response. If you have a lot of
// traffic policy instances, you can use the MaxItems parameter to list them in
// groups of up to 100.
func (c *Client) ListTrafficPolicyInstancesByPolicy(ctx context.Context, params *ListTrafficPolicyInstancesByPolicyInput, optFns ...func(*Options)) (*ListTrafficPolicyInstancesByPolicyOutput, error) {
	if params == nil {
		params = &ListTrafficPolicyInstancesByPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrafficPolicyInstancesByPolicy", params, optFns, c.addOperationListTrafficPolicyInstancesByPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrafficPolicyInstancesByPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains the information about the request to list your
// traffic policy instances.
type ListTrafficPolicyInstancesByPolicyInput struct {

	// The ID of the traffic policy for which you want to list traffic policy
	// instances.
	//
	// This member is required.
	TrafficPolicyId *string

	// The version of the traffic policy for which you want to list traffic policy
	// instances. The version must be associated with the traffic policy that is
	// specified by TrafficPolicyId .
	//
	// This member is required.
	TrafficPolicyVersion *int32

	// If the value of IsTruncated in the previous response was true , you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstancesByPolicy request.
	//
	// For the value of hostedzoneid , specify the value of HostedZoneIdMarker from
	// the previous response, which is the hosted zone ID of the first traffic policy
	// instance that Amazon Route 53 will return if you submit another request.
	//
	// If the value of IsTruncated in the previous response was false , there are no
	// more traffic policy instances to get.
	HostedZoneIdMarker *string

	// The maximum number of traffic policy instances to be included in the response
	// body for this request. If you have more than MaxItems traffic policy instances,
	// the value of the IsTruncated element in the response is true , and the values of
	// HostedZoneIdMarker , TrafficPolicyInstanceNameMarker , and
	// TrafficPolicyInstanceTypeMarker represent the first traffic policy instance that
	// Amazon Route 53 will return if you submit another request.
	MaxItems *int32

	// If the value of IsTruncated in the previous response was true , you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstancesByPolicy request.
	//
	// For the value of trafficpolicyinstancename , specify the value of
	// TrafficPolicyInstanceNameMarker from the previous response, which is the name of
	// the first traffic policy instance that Amazon Route 53 will return if you submit
	// another request.
	//
	// If the value of IsTruncated in the previous response was false , there are no
	// more traffic policy instances to get.
	TrafficPolicyInstanceNameMarker *string

	// If the value of IsTruncated in the previous response was true , you have more
	// traffic policy instances. To get more traffic policy instances, submit another
	// ListTrafficPolicyInstancesByPolicy request.
	//
	// For the value of trafficpolicyinstancetype , specify the value of
	// TrafficPolicyInstanceTypeMarker from the previous response, which is the name of
	// the first traffic policy instance that Amazon Route 53 will return if you submit
	// another request.
	//
	// If the value of IsTruncated in the previous response was false , there are no
	// more traffic policy instances to get.
	TrafficPolicyInstanceTypeMarker types.RRType

	noSmithyDocumentSerde
}

// A complex type that contains the response information for the request.
type ListTrafficPolicyInstancesByPolicyOutput struct {

	// A flag that indicates whether there are more traffic policy instances to be
	// listed. If the response was truncated, you can get the next group of traffic
	// policy instances by calling ListTrafficPolicyInstancesByPolicy again and
	// specifying the values of the HostedZoneIdMarker ,
	// TrafficPolicyInstanceNameMarker , and TrafficPolicyInstanceTypeMarker elements
	// in the corresponding request parameters.
	//
	// This member is required.
	IsTruncated bool

	// The value that you specified for the MaxItems parameter in the call to
	// ListTrafficPolicyInstancesByPolicy that produced the current response.
	//
	// This member is required.
	MaxItems *int32

	// A list that contains one TrafficPolicyInstance element for each traffic policy
	// instance that matches the elements in the request.
	//
	// This member is required.
	TrafficPolicyInstances []types.TrafficPolicyInstance

	// If IsTruncated is true , HostedZoneIdMarker is the ID of the hosted zone of the
	// first traffic policy instance in the next group of traffic policy instances.
	HostedZoneIdMarker *string

	// If IsTruncated is true , TrafficPolicyInstanceNameMarker is the name of the
	// first traffic policy instance in the next group of MaxItems traffic policy
	// instances.
	TrafficPolicyInstanceNameMarker *string

	// If IsTruncated is true , TrafficPolicyInstanceTypeMarker is the DNS type of the
	// resource record sets that are associated with the first traffic policy instance
	// in the next group of MaxItems traffic policy instances.
	TrafficPolicyInstanceTypeMarker types.RRType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTrafficPolicyInstancesByPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListTrafficPolicyInstancesByPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListTrafficPolicyInstancesByPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTrafficPolicyInstancesByPolicy"); err != nil {
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
	if err = addSpanRetryLoop(stack, options); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpListTrafficPolicyInstancesByPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrafficPolicyInstancesByPolicy(options.Region), middleware.Before); err != nil {
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
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListTrafficPolicyInstancesByPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTrafficPolicyInstancesByPolicy",
	}
}
