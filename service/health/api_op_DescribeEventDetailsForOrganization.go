// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/health/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns detailed information about one or more specified events for one or more
// Amazon Web Services accounts in your organization. This information includes
// standard event data (such as the Amazon Web Services Region and service), an
// event description, and (depending on the event) possible metadata. This
// operation doesn't return affected entities, such as the resources related to the
// event. To return affected entities, use the
// DescribeAffectedEntitiesForOrganization (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntitiesForOrganization.html)
// operation. Before you can call this operation, you must first enable Health to
// work with Organizations. To do this, call the
// EnableHealthServiceAccessForOrganization (https://docs.aws.amazon.com/health/latest/APIReference/API_EnableHealthServiceAccessForOrganization.html)
// operation from your organization's management account. When you call the
// DescribeEventDetailsForOrganization operation, specify the
// organizationEventDetailFilters object in the request. Depending on the Health
// event type, note the following differences:
//   - To return event details for a public event, you must specify a null value
//     for the awsAccountId parameter. If you specify an account ID for a public
//     event, Health returns an error message because public events aren't specific to
//     an account.
//   - To return event details for an event that is specific to an account in your
//     organization, you must specify the awsAccountId parameter in the request. If
//     you don't specify an account ID, Health returns an error message because the
//     event is specific to an account in your organization.
//
// For more information, see Event (https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html)
// . This operation doesn't support resource-level permissions. You can't use this
// operation to allow or deny access to specific Health events. For more
// information, see Resource- and action-based conditions (https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html#resource-action-based-conditions)
// in the Health User Guide.
func (c *Client) DescribeEventDetailsForOrganization(ctx context.Context, params *DescribeEventDetailsForOrganizationInput, optFns ...func(*Options)) (*DescribeEventDetailsForOrganizationOutput, error) {
	if params == nil {
		params = &DescribeEventDetailsForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventDetailsForOrganization", params, optFns, c.addOperationDescribeEventDetailsForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventDetailsForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventDetailsForOrganizationInput struct {

	// A set of JSON elements that includes the awsAccountId and the eventArn .
	//
	// This member is required.
	OrganizationEventDetailFilters []types.EventAccountFilter

	// The locale (language) to return information in. English (en) is the default and
	// the only supported value at this time.
	Locale *string

	noSmithyDocumentSerde
}

type DescribeEventDetailsForOrganizationOutput struct {

	// Error messages for any events that could not be retrieved.
	FailedSet []types.OrganizationEventDetailsErrorItem

	// Information about the events that could be retrieved.
	SuccessfulSet []types.OrganizationEventDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEventDetailsForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventDetailsForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventDetailsForOrganization{}, middleware.After)
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
	if err = addOpDescribeEventDetailsForOrganizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventDetailsForOrganization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEventDetailsForOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "health",
		OperationName: "DescribeEventDetailsForOrganization",
	}
}
