// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the specified notification subscriptions.
func (c *Client) DescribeNotificationSubscriptions(ctx context.Context, params *DescribeNotificationSubscriptionsInput, optFns ...func(*Options)) (*DescribeNotificationSubscriptionsOutput, error) {
	if params == nil {
		params = &DescribeNotificationSubscriptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNotificationSubscriptions", params, optFns, c.addOperationDescribeNotificationSubscriptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNotificationSubscriptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNotificationSubscriptionsInput struct {

	// The ID of the organization.
	//
	// This member is required.
	OrganizationId *string

	// The maximum number of items to return with this call.
	Limit *int32

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string

	noSmithyDocumentSerde
}

type DescribeNotificationSubscriptionsOutput struct {

	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	Marker *string

	// The subscriptions.
	Subscriptions []types.Subscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeNotificationSubscriptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeNotificationSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeNotificationSubscriptions{}, middleware.After)
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
	if err = addOpDescribeNotificationSubscriptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNotificationSubscriptions(options.Region), middleware.Before); err != nil {
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

// DescribeNotificationSubscriptionsAPIClient is a client that implements the
// DescribeNotificationSubscriptions operation.
type DescribeNotificationSubscriptionsAPIClient interface {
	DescribeNotificationSubscriptions(context.Context, *DescribeNotificationSubscriptionsInput, ...func(*Options)) (*DescribeNotificationSubscriptionsOutput, error)
}

var _ DescribeNotificationSubscriptionsAPIClient = (*Client)(nil)

// DescribeNotificationSubscriptionsPaginatorOptions is the paginator options for
// DescribeNotificationSubscriptions
type DescribeNotificationSubscriptionsPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeNotificationSubscriptionsPaginator is a paginator for
// DescribeNotificationSubscriptions
type DescribeNotificationSubscriptionsPaginator struct {
	options   DescribeNotificationSubscriptionsPaginatorOptions
	client    DescribeNotificationSubscriptionsAPIClient
	params    *DescribeNotificationSubscriptionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeNotificationSubscriptionsPaginator returns a new
// DescribeNotificationSubscriptionsPaginator
func NewDescribeNotificationSubscriptionsPaginator(client DescribeNotificationSubscriptionsAPIClient, params *DescribeNotificationSubscriptionsInput, optFns ...func(*DescribeNotificationSubscriptionsPaginatorOptions)) *DescribeNotificationSubscriptionsPaginator {
	if params == nil {
		params = &DescribeNotificationSubscriptionsInput{}
	}

	options := DescribeNotificationSubscriptionsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeNotificationSubscriptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeNotificationSubscriptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeNotificationSubscriptions page.
func (p *DescribeNotificationSubscriptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeNotificationSubscriptionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeNotificationSubscriptions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeNotificationSubscriptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "DescribeNotificationSubscriptions",
	}
}
