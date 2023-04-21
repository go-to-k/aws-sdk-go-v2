// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the source locations for a channel. A source location defines the host
// server URL, and contains a list of sources.
func (c *Client) ListSourceLocations(ctx context.Context, params *ListSourceLocationsInput, optFns ...func(*Options)) (*ListSourceLocationsOutput, error) {
	if params == nil {
		params = &ListSourceLocationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSourceLocations", params, optFns, c.addOperationListSourceLocationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSourceLocationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSourceLocationsInput struct {

	// The maximum number of source locations that you want MediaTailor to return in
	// response to the current request. If there are more than MaxResults source
	// locations, use the value of NextToken in the response to get the next page of
	// results.
	MaxResults int32

	// Pagination token returned by the list request when results exceed the maximum
	// allowed. Use the token to fetch the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSourceLocationsOutput struct {

	// A list of source locations.
	Items []types.SourceLocation

	// Pagination token returned by the list request when results exceed the maximum
	// allowed. Use the token to fetch the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSourceLocationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSourceLocations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSourceLocations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSourceLocations(options.Region), middleware.Before); err != nil {
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

// ListSourceLocationsAPIClient is a client that implements the
// ListSourceLocations operation.
type ListSourceLocationsAPIClient interface {
	ListSourceLocations(context.Context, *ListSourceLocationsInput, ...func(*Options)) (*ListSourceLocationsOutput, error)
}

var _ ListSourceLocationsAPIClient = (*Client)(nil)

// ListSourceLocationsPaginatorOptions is the paginator options for
// ListSourceLocations
type ListSourceLocationsPaginatorOptions struct {
	// The maximum number of source locations that you want MediaTailor to return in
	// response to the current request. If there are more than MaxResults source
	// locations, use the value of NextToken in the response to get the next page of
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSourceLocationsPaginator is a paginator for ListSourceLocations
type ListSourceLocationsPaginator struct {
	options   ListSourceLocationsPaginatorOptions
	client    ListSourceLocationsAPIClient
	params    *ListSourceLocationsInput
	nextToken *string
	firstPage bool
}

// NewListSourceLocationsPaginator returns a new ListSourceLocationsPaginator
func NewListSourceLocationsPaginator(client ListSourceLocationsAPIClient, params *ListSourceLocationsInput, optFns ...func(*ListSourceLocationsPaginatorOptions)) *ListSourceLocationsPaginator {
	if params == nil {
		params = &ListSourceLocationsInput{}
	}

	options := ListSourceLocationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSourceLocationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSourceLocationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSourceLocations page.
func (p *ListSourceLocationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSourceLocationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListSourceLocations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSourceLocations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "ListSourceLocations",
	}
}
