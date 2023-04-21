// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a paginated list of properties associated with an asset. If you
// update properties associated with the model before you finish listing all the
// properties, you need to start all over again.
func (c *Client) ListAssetProperties(ctx context.Context, params *ListAssetPropertiesInput, optFns ...func(*Options)) (*ListAssetPropertiesOutput, error) {
	if params == nil {
		params = &ListAssetPropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssetProperties", params, optFns, c.addOperationListAssetPropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssetPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssetPropertiesInput struct {

	// The ID of the asset.
	//
	// This member is required.
	AssetId *string

	// Filters the requested list of asset properties. You can choose one of the
	// following options:
	//   - ALL – The list includes all asset properties for a given asset model ID.
	//   - BASE – The list includes only base asset properties for a given asset model
	//   ID.
	// Default: BASE
	Filter types.ListAssetPropertiesFilter

	// The maximum number of results to return for each paginated request. If not
	// specified, the default value is 50.
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssetPropertiesOutput struct {

	// A list that summarizes the properties associated with the specified asset.
	//
	// This member is required.
	AssetPropertySummaries []types.AssetPropertySummary

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssetPropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssetProperties{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssetProperties{}, middleware.After)
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
	if err = addEndpointPrefix_opListAssetPropertiesMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListAssetPropertiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssetProperties(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListAssetPropertiesMiddleware struct {
}

func (*endpointPrefix_opListAssetPropertiesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListAssetPropertiesMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListAssetPropertiesMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListAssetPropertiesMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListAssetPropertiesAPIClient is a client that implements the
// ListAssetProperties operation.
type ListAssetPropertiesAPIClient interface {
	ListAssetProperties(context.Context, *ListAssetPropertiesInput, ...func(*Options)) (*ListAssetPropertiesOutput, error)
}

var _ ListAssetPropertiesAPIClient = (*Client)(nil)

// ListAssetPropertiesPaginatorOptions is the paginator options for
// ListAssetProperties
type ListAssetPropertiesPaginatorOptions struct {
	// The maximum number of results to return for each paginated request. If not
	// specified, the default value is 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssetPropertiesPaginator is a paginator for ListAssetProperties
type ListAssetPropertiesPaginator struct {
	options   ListAssetPropertiesPaginatorOptions
	client    ListAssetPropertiesAPIClient
	params    *ListAssetPropertiesInput
	nextToken *string
	firstPage bool
}

// NewListAssetPropertiesPaginator returns a new ListAssetPropertiesPaginator
func NewListAssetPropertiesPaginator(client ListAssetPropertiesAPIClient, params *ListAssetPropertiesInput, optFns ...func(*ListAssetPropertiesPaginatorOptions)) *ListAssetPropertiesPaginator {
	if params == nil {
		params = &ListAssetPropertiesInput{}
	}

	options := ListAssetPropertiesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssetPropertiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssetPropertiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssetProperties page.
func (p *ListAssetPropertiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssetPropertiesOutput, error) {
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

	result, err := p.client.ListAssetProperties(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssetProperties(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "ListAssetProperties",
	}
}
