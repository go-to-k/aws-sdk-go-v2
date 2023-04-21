// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagevod

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a collection of MediaPackage VOD Asset resources.
func (c *Client) ListAssets(ctx context.Context, params *ListAssetsInput, optFns ...func(*Options)) (*ListAssetsOutput, error) {
	if params == nil {
		params = &ListAssetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssets", params, optFns, c.addOperationListAssetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssetsInput struct {

	// Upper bound on number of records to return.
	MaxResults int32

	// A token used to resume pagination from the end of a previous request.
	NextToken *string

	// Returns Assets associated with the specified PackagingGroup.
	PackagingGroupId *string

	noSmithyDocumentSerde
}

type ListAssetsOutput struct {

	// A list of MediaPackage VOD Asset resources.
	Assets []types.AssetShallow

	// A token that can be used to resume pagination from the end of the collection.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssets(options.Region), middleware.Before); err != nil {
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

// ListAssetsAPIClient is a client that implements the ListAssets operation.
type ListAssetsAPIClient interface {
	ListAssets(context.Context, *ListAssetsInput, ...func(*Options)) (*ListAssetsOutput, error)
}

var _ ListAssetsAPIClient = (*Client)(nil)

// ListAssetsPaginatorOptions is the paginator options for ListAssets
type ListAssetsPaginatorOptions struct {
	// Upper bound on number of records to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssetsPaginator is a paginator for ListAssets
type ListAssetsPaginator struct {
	options   ListAssetsPaginatorOptions
	client    ListAssetsAPIClient
	params    *ListAssetsInput
	nextToken *string
	firstPage bool
}

// NewListAssetsPaginator returns a new ListAssetsPaginator
func NewListAssetsPaginator(client ListAssetsAPIClient, params *ListAssetsInput, optFns ...func(*ListAssetsPaginatorOptions)) *ListAssetsPaginator {
	if params == nil {
		params = &ListAssetsInput{}
	}

	options := ListAssetsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssets page.
func (p *ListAssetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssetsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListAssets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage-vod",
		OperationName: "ListAssets",
	}
}
