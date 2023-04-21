// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists API key resources in your Amazon Web Services account. The API keys
// feature is in preview. We may add, change, or remove features before announcing
// general availability. For more information, see Using API keys (https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html)
// .
func (c *Client) ListKeys(ctx context.Context, params *ListKeysInput, optFns ...func(*Options)) (*ListKeysOutput, error) {
	if params == nil {
		params = &ListKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListKeys", params, optFns, c.addOperationListKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListKeysInput struct {

	// Optionally filter the list to only Active or Expired API keys.
	Filter *types.ApiKeyFilter

	// An optional limit for the number of resources returned in a single call.
	// Default value: 100
	MaxResults *int32

	// The pagination token specifying which page of results to return in the
	// response. If no token is provided, the default page is the first page. Default
	// value: null
	NextToken *string

	noSmithyDocumentSerde
}

type ListKeysOutput struct {

	// Contains API key resources in your Amazon Web Services account. Details include
	// API key name, allowed referers and timestamp for when the API key will expire.
	//
	// This member is required.
	Entries []types.ListKeysResponseEntry

	// A pagination token indicating there are additional pages available. You can use
	// the token in a following request to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListKeys{}, middleware.After)
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
	if err = addEndpointPrefix_opListKeysMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListKeys(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListKeysMiddleware struct {
}

func (*endpointPrefix_opListKeysMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListKeysMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "metadata." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListKeysMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListKeysMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListKeysAPIClient is a client that implements the ListKeys operation.
type ListKeysAPIClient interface {
	ListKeys(context.Context, *ListKeysInput, ...func(*Options)) (*ListKeysOutput, error)
}

var _ ListKeysAPIClient = (*Client)(nil)

// ListKeysPaginatorOptions is the paginator options for ListKeys
type ListKeysPaginatorOptions struct {
	// An optional limit for the number of resources returned in a single call.
	// Default value: 100
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListKeysPaginator is a paginator for ListKeys
type ListKeysPaginator struct {
	options   ListKeysPaginatorOptions
	client    ListKeysAPIClient
	params    *ListKeysInput
	nextToken *string
	firstPage bool
}

// NewListKeysPaginator returns a new ListKeysPaginator
func NewListKeysPaginator(client ListKeysAPIClient, params *ListKeysInput, optFns ...func(*ListKeysPaginatorOptions)) *ListKeysPaginator {
	if params == nil {
		params = &ListKeysInput{}
	}

	options := ListKeysPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListKeysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListKeysPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListKeys page.
func (p *ListKeysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListKeysOutput, error) {
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

	result, err := p.client.ListKeys(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "ListKeys",
	}
}
