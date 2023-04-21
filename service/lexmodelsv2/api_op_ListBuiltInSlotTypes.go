// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of built-in slot types that meet the specified criteria.
func (c *Client) ListBuiltInSlotTypes(ctx context.Context, params *ListBuiltInSlotTypesInput, optFns ...func(*Options)) (*ListBuiltInSlotTypesOutput, error) {
	if params == nil {
		params = &ListBuiltInSlotTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBuiltInSlotTypes", params, optFns, c.addOperationListBuiltInSlotTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBuiltInSlotTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBuiltInSlotTypesInput struct {

	// The identifier of the language and locale of the slot types to list. The string
	// must match one of the supported locales. For more information, see Supported
	// languages (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html) .
	//
	// This member is required.
	LocaleId *string

	// The maximum number of built-in slot types to return in each page of results. If
	// there are fewer results than the max page size, only the actual number of
	// results are returned.
	MaxResults *int32

	// If the response from the ListBuiltInSlotTypes operation contains more results
	// than specified in the maxResults parameter, a token is returned in the
	// response. Use that token in the nextToken parameter to return the next page of
	// results.
	NextToken *string

	// Determines the sort order for the response from the ListBuiltInSlotTypes
	// operation. You can choose to sort by the slot type signature in either ascending
	// or descending order.
	SortBy *types.BuiltInSlotTypeSortBy

	noSmithyDocumentSerde
}

type ListBuiltInSlotTypesOutput struct {

	// Summary information for the built-in slot types that meet the filter criteria
	// specified in the request. The length of the list is specified in the maxResults
	// parameter of the request. If there are more slot types available, the nextToken
	// field contains a token to get the next page of results.
	BuiltInSlotTypeSummaries []types.BuiltInSlotTypeSummary

	// The language and locale of the slot types in the list.
	LocaleId *string

	// A token that indicates whether there are more results to return in a response
	// to the ListBuiltInSlotTypes operation. If the nextToken field is present, you
	// send the contents as the nextToken parameter of a LIstBuiltInSlotTypes
	// operation request to get the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBuiltInSlotTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBuiltInSlotTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBuiltInSlotTypes{}, middleware.After)
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
	if err = addOpListBuiltInSlotTypesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBuiltInSlotTypes(options.Region), middleware.Before); err != nil {
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

// ListBuiltInSlotTypesAPIClient is a client that implements the
// ListBuiltInSlotTypes operation.
type ListBuiltInSlotTypesAPIClient interface {
	ListBuiltInSlotTypes(context.Context, *ListBuiltInSlotTypesInput, ...func(*Options)) (*ListBuiltInSlotTypesOutput, error)
}

var _ ListBuiltInSlotTypesAPIClient = (*Client)(nil)

// ListBuiltInSlotTypesPaginatorOptions is the paginator options for
// ListBuiltInSlotTypes
type ListBuiltInSlotTypesPaginatorOptions struct {
	// The maximum number of built-in slot types to return in each page of results. If
	// there are fewer results than the max page size, only the actual number of
	// results are returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBuiltInSlotTypesPaginator is a paginator for ListBuiltInSlotTypes
type ListBuiltInSlotTypesPaginator struct {
	options   ListBuiltInSlotTypesPaginatorOptions
	client    ListBuiltInSlotTypesAPIClient
	params    *ListBuiltInSlotTypesInput
	nextToken *string
	firstPage bool
}

// NewListBuiltInSlotTypesPaginator returns a new ListBuiltInSlotTypesPaginator
func NewListBuiltInSlotTypesPaginator(client ListBuiltInSlotTypesAPIClient, params *ListBuiltInSlotTypesInput, optFns ...func(*ListBuiltInSlotTypesPaginatorOptions)) *ListBuiltInSlotTypesPaginator {
	if params == nil {
		params = &ListBuiltInSlotTypesInput{}
	}

	options := ListBuiltInSlotTypesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBuiltInSlotTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBuiltInSlotTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBuiltInSlotTypes page.
func (p *ListBuiltInSlotTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBuiltInSlotTypesOutput, error) {
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

	result, err := p.client.ListBuiltInSlotTypes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBuiltInSlotTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "ListBuiltInSlotTypes",
	}
}
