// Code generated by smithy-go-codegen DO NOT EDIT.

package privatenetworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/privatenetworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists orders. Add filters to your request to return a more specific list of
// results. Use filters to match the Amazon Resource Name (ARN) of the network site
// or the status of the order. If you specify multiple filters, filters are joined
// with an OR, and the request returns results that match all of the specified
// filters.
func (c *Client) ListOrders(ctx context.Context, params *ListOrdersInput, optFns ...func(*Options)) (*ListOrdersOutput, error) {
	if params == nil {
		params = &ListOrdersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOrders", params, optFns, c.addOperationListOrdersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOrdersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOrdersInput struct {

	// The Amazon Resource Name (ARN) of the network.
	//
	// This member is required.
	NetworkArn *string

	// The filters.
	//   - NETWORK_SITE - The Amazon Resource Name (ARN) of the network site.
	//   - STATUS - The status ( ACKNOWLEDGING | ACKNOWLEDGED | UNACKNOWLEDGED ).
	// Filter values are case sensitive. If you specify multiple values for a filter,
	// the values are joined with an OR , and the request returns all results that
	// match any of the specified values.
	Filters map[string][]string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	StartToken *string

	noSmithyDocumentSerde
}

type ListOrdersOutput struct {

	// The token for the next page of results.
	NextToken *string

	// Information about the orders.
	Orders []types.Order

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOrdersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOrders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOrders{}, middleware.After)
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
	if err = addOpListOrdersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOrders(options.Region), middleware.Before); err != nil {
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

// ListOrdersAPIClient is a client that implements the ListOrders operation.
type ListOrdersAPIClient interface {
	ListOrders(context.Context, *ListOrdersInput, ...func(*Options)) (*ListOrdersOutput, error)
}

var _ ListOrdersAPIClient = (*Client)(nil)

// ListOrdersPaginatorOptions is the paginator options for ListOrders
type ListOrdersPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOrdersPaginator is a paginator for ListOrders
type ListOrdersPaginator struct {
	options   ListOrdersPaginatorOptions
	client    ListOrdersAPIClient
	params    *ListOrdersInput
	nextToken *string
	firstPage bool
}

// NewListOrdersPaginator returns a new ListOrdersPaginator
func NewListOrdersPaginator(client ListOrdersAPIClient, params *ListOrdersInput, optFns ...func(*ListOrdersPaginatorOptions)) *ListOrdersPaginator {
	if params == nil {
		params = &ListOrdersInput{}
	}

	options := ListOrdersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOrdersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.StartToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOrdersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOrders page.
func (p *ListOrdersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOrdersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.StartToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListOrders(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOrders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "private-networks",
		OperationName: "ListOrders",
	}
}
