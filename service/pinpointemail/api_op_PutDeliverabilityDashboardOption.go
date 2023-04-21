// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enable or disable the Deliverability dashboard for your Amazon Pinpoint
// account. When you enable the Deliverability dashboard, you gain access to
// reputation, deliverability, and other metrics for the domains that you use to
// send email using Amazon Pinpoint. You also gain the ability to perform
// predictive inbox placement tests. When you use the Deliverability dashboard, you
// pay a monthly subscription charge, in addition to any other fees that you accrue
// by using Amazon Pinpoint. For more information about the features and cost of a
// Deliverability dashboard subscription, see Amazon Pinpoint Pricing (http://aws.amazon.com/pinpoint/pricing/)
// .
func (c *Client) PutDeliverabilityDashboardOption(ctx context.Context, params *PutDeliverabilityDashboardOptionInput, optFns ...func(*Options)) (*PutDeliverabilityDashboardOptionOutput, error) {
	if params == nil {
		params = &PutDeliverabilityDashboardOptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutDeliverabilityDashboardOption", params, optFns, c.addOperationPutDeliverabilityDashboardOptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutDeliverabilityDashboardOptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Enable or disable the Deliverability dashboard for your Amazon Pinpoint
// account. When you enable the Deliverability dashboard, you gain access to
// reputation, deliverability, and other metrics for the domains that you use to
// send email using Amazon Pinpoint. You also gain the ability to perform
// predictive inbox placement tests. When you use the Deliverability dashboard, you
// pay a monthly subscription charge, in addition to any other fees that you accrue
// by using Amazon Pinpoint. For more information about the features and cost of a
// Deliverability dashboard subscription, see Amazon Pinpoint Pricing (http://aws.amazon.com/pinpoint/pricing/)
// .
type PutDeliverabilityDashboardOptionInput struct {

	// Specifies whether to enable the Deliverability dashboard for your Amazon
	// Pinpoint account. To enable the dashboard, set this value to true .
	//
	// This member is required.
	DashboardEnabled bool

	// An array of objects, one for each verified domain that you use to send email
	// and enabled the Deliverability dashboard for.
	SubscribedDomains []types.DomainDeliverabilityTrackingOption

	noSmithyDocumentSerde
}

// A response that indicates whether the Deliverability dashboard is enabled for
// your Amazon Pinpoint account.
type PutDeliverabilityDashboardOptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutDeliverabilityDashboardOptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutDeliverabilityDashboardOption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutDeliverabilityDashboardOption{}, middleware.After)
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
	if err = addOpPutDeliverabilityDashboardOptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutDeliverabilityDashboardOption(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutDeliverabilityDashboardOption(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutDeliverabilityDashboardOption",
	}
}
