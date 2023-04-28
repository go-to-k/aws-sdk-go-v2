// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets detailed information about a specified number of requests--a sample--that
// WAF randomly selects from among the first 5,000 requests that your Amazon Web
// Services resource received during a time range that you choose. You can specify
// a sample size of up to 500 requests, and you can specify any time range in the
// previous three hours. GetSampledRequests returns a time range, which is usually
// the time range that you specified. However, if your resource (such as a
// CloudFront distribution) received 5,000 requests before the specified time range
// elapsed, GetSampledRequests returns an updated time range. This new time range
// indicates the actual period during which WAF selected the requests in the
// sample.
func (c *Client) GetSampledRequests(ctx context.Context, params *GetSampledRequestsInput, optFns ...func(*Options)) (*GetSampledRequestsOutput, error) {
	if params == nil {
		params = &GetSampledRequestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSampledRequests", params, optFns, c.addOperationGetSampledRequestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSampledRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSampledRequestsInput struct {

	// The number of requests that you want WAF to return from among the first 5,000
	// requests that your Amazon Web Services resource received during the time range.
	// If your resource received fewer requests than the value of MaxItems ,
	// GetSampledRequests returns information about all of them.
	//
	// This member is required.
	MaxItems int64

	// The metric name assigned to the Rule or RuleGroup dimension for which you want
	// a sample of requests.
	//
	// This member is required.
	RuleMetricName *string

	// Specifies whether this is for an Amazon CloudFront distribution or for a
	// regional application. A regional application can be an Application Load Balancer
	// (ALB), an Amazon API Gateway REST API, an AppSync GraphQL API, an Amazon Cognito
	// user pool, an App Runner service, or an Amazon Web Services Verified Access
	// instance. To work with CloudFront, you must also specify the Region US East (N.
	// Virginia) as follows:
	//   - CLI - Specify the Region when you use the CloudFront scope:
	//   --scope=CLOUDFRONT --region=us-east-1 .
	//   - API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// The start date and time and the end date and time of the range for which you
	// want GetSampledRequests to return a sample of requests. You must specify the
	// times in Coordinated Universal Time (UTC) format. UTC format includes the
	// special designator, Z . For example, "2016-09-27T14:50Z" . You can specify any
	// time range in the previous three hours. If you specify a start time that's
	// earlier than three hours ago, WAF sets it to three hours ago.
	//
	// This member is required.
	TimeWindow *types.TimeWindow

	// The Amazon resource name (ARN) of the WebACL for which you want a sample of
	// requests.
	//
	// This member is required.
	WebAclArn *string

	noSmithyDocumentSerde
}

type GetSampledRequestsOutput struct {

	// The total number of requests from which GetSampledRequests got a sample of
	// MaxItems requests. If PopulationSize is less than MaxItems , the sample includes
	// every request that your Amazon Web Services resource received during the
	// specified time range.
	PopulationSize int64

	// A complex type that contains detailed information about each of the requests in
	// the sample.
	SampledRequests []types.SampledHTTPRequest

	// Usually, TimeWindow is the time range that you specified in the
	// GetSampledRequests request. However, if your Amazon Web Services resource
	// received more than 5,000 requests during the time range that you specified in
	// the request, GetSampledRequests returns the time range for the first 5,000
	// requests. Times are in Coordinated Universal Time (UTC) format.
	TimeWindow *types.TimeWindow

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSampledRequestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSampledRequests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSampledRequests{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetSampledRequestsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSampledRequests(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opGetSampledRequests(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "GetSampledRequests",
	}
}
