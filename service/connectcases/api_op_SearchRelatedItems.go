// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcases

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connectcases/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches for related items that are associated with a case. If no filters are
// provided, this returns all related items associated with a case.
func (c *Client) SearchRelatedItems(ctx context.Context, params *SearchRelatedItemsInput, optFns ...func(*Options)) (*SearchRelatedItemsOutput, error) {
	if params == nil {
		params = &SearchRelatedItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchRelatedItems", params, optFns, c.addOperationSearchRelatedItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchRelatedItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchRelatedItemsInput struct {

	// A unique identifier of the case.
	//
	// This member is required.
	CaseId *string

	// The unique identifier of the Cases domain.
	//
	// This member is required.
	DomainId *string

	// The list of types of related items and their parameters to use for filtering.
	Filters []types.RelatedItemTypeFilter

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchRelatedItemsOutput struct {

	// A list of items related to a case.
	//
	// This member is required.
	RelatedItems []*types.SearchRelatedItemsResponseItem

	// The token for the next set of results. This is null if there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchRelatedItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchRelatedItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchRelatedItems{}, middleware.After)
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
	if err = addOpSearchRelatedItemsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchRelatedItems(options.Region), middleware.Before); err != nil {
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

// SearchRelatedItemsAPIClient is a client that implements the SearchRelatedItems
// operation.
type SearchRelatedItemsAPIClient interface {
	SearchRelatedItems(context.Context, *SearchRelatedItemsInput, ...func(*Options)) (*SearchRelatedItemsOutput, error)
}

var _ SearchRelatedItemsAPIClient = (*Client)(nil)

// SearchRelatedItemsPaginatorOptions is the paginator options for
// SearchRelatedItems
type SearchRelatedItemsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchRelatedItemsPaginator is a paginator for SearchRelatedItems
type SearchRelatedItemsPaginator struct {
	options   SearchRelatedItemsPaginatorOptions
	client    SearchRelatedItemsAPIClient
	params    *SearchRelatedItemsInput
	nextToken *string
	firstPage bool
}

// NewSearchRelatedItemsPaginator returns a new SearchRelatedItemsPaginator
func NewSearchRelatedItemsPaginator(client SearchRelatedItemsAPIClient, params *SearchRelatedItemsInput, optFns ...func(*SearchRelatedItemsPaginatorOptions)) *SearchRelatedItemsPaginator {
	if params == nil {
		params = &SearchRelatedItemsInput{}
	}

	options := SearchRelatedItemsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchRelatedItemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchRelatedItemsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchRelatedItems page.
func (p *SearchRelatedItemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchRelatedItemsOutput, error) {
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

	result, err := p.client.SearchRelatedItems(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchRelatedItems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cases",
		OperationName: "SearchRelatedItems",
	}
}
