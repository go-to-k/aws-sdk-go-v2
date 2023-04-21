// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information for each created fleet in an Amazon Web Services account.
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func (c *Client) ListFleets(ctx context.Context, params *ListFleetsInput, optFns ...func(*Options)) (*ListFleetsOutput, error) {
	if params == nil {
		params = &ListFleetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFleets", params, optFns, c.addOperationListFleetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFleetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFleetsInput struct {

	// The maximum number of items to return, between 1 and 100, inclusive.
	MaxResults *int32

	// A pagination token for the next set of results. If the results of a search are
	// large, only a portion of the results are returned, and a nextToken pagination
	// token is returned in the response. To retrieve the next set of results, reissue
	// the search request and include the returned token. When all results have been
	// returned, the response does not contain a pagination token value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFleetsOutput struct {

	// A list of information for each fleet.
	FleetSummaries []types.FleetSummary

	// The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFleetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListFleets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListFleets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFleets(options.Region), middleware.Before); err != nil {
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

// ListFleetsAPIClient is a client that implements the ListFleets operation.
type ListFleetsAPIClient interface {
	ListFleets(context.Context, *ListFleetsInput, ...func(*Options)) (*ListFleetsOutput, error)
}

var _ ListFleetsAPIClient = (*Client)(nil)

// ListFleetsPaginatorOptions is the paginator options for ListFleets
type ListFleetsPaginatorOptions struct {
	// The maximum number of items to return, between 1 and 100, inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFleetsPaginator is a paginator for ListFleets
type ListFleetsPaginator struct {
	options   ListFleetsPaginatorOptions
	client    ListFleetsAPIClient
	params    *ListFleetsInput
	nextToken *string
	firstPage bool
}

// NewListFleetsPaginator returns a new ListFleetsPaginator
func NewListFleetsPaginator(client ListFleetsAPIClient, params *ListFleetsInput, optFns ...func(*ListFleetsPaginatorOptions)) *ListFleetsPaginator {
	if params == nil {
		params = &ListFleetsInput{}
	}

	options := ListFleetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFleetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFleetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFleets page.
func (p *ListFleetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFleetsOutput, error) {
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

	result, err := p.client.ListFleets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFleets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "ListFleets",
	}
}
