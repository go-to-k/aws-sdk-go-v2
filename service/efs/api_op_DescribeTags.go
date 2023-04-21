// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// DEPRECATED - The DescribeTags action is deprecated and not maintained. To view
// tags associated with EFS resources, use the ListTagsForResource API action.
// Returns the tags associated with a file system. The order of tags returned in
// the response of one DescribeTags call and the order of tags returned across the
// responses of a multiple-call iteration (when using pagination) is unspecified.
// This operation requires permissions for the elasticfilesystem:DescribeTags
// action.
//
// Deprecated: Use ListTagsForResource.
func (c *Client) DescribeTags(ctx context.Context, params *DescribeTagsInput, optFns ...func(*Options)) (*DescribeTagsOutput, error) {
	if params == nil {
		params = &DescribeTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTags", params, optFns, c.addOperationDescribeTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTagsInput struct {

	// The ID of the file system whose tag set you want to retrieve.
	//
	// This member is required.
	FileSystemId *string

	// (Optional) An opaque pagination token returned from a previous DescribeTags
	// operation (String). If present, it specifies to continue the list from where the
	// previous call left off.
	Marker *string

	// (Optional) The maximum number of file system tags to return in the response.
	// Currently, this number is automatically set to 100, and other values are
	// ignored. The response is paginated at 100 per page if you have more than 100
	// tags.
	MaxItems *int32

	noSmithyDocumentSerde
}

type DescribeTagsOutput struct {

	// Returns tags associated with the file system as an array of Tag objects.
	//
	// This member is required.
	Tags []types.Tag

	// If the request included a Marker , the response returns that value in this field.
	Marker *string

	// If a value is present, there are more tags to return. In a subsequent request,
	// you can provide the value of NextMarker as the value of the Marker parameter in
	// your next request to retrieve the next set of tags.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeTags{}, middleware.After)
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
	if err = addOpDescribeTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTags(options.Region), middleware.Before); err != nil {
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

// DescribeTagsAPIClient is a client that implements the DescribeTags operation.
type DescribeTagsAPIClient interface {
	DescribeTags(context.Context, *DescribeTagsInput, ...func(*Options)) (*DescribeTagsOutput, error)
}

var _ DescribeTagsAPIClient = (*Client)(nil)

// DescribeTagsPaginatorOptions is the paginator options for DescribeTags
type DescribeTagsPaginatorOptions struct {
	// (Optional) The maximum number of file system tags to return in the response.
	// Currently, this number is automatically set to 100, and other values are
	// ignored. The response is paginated at 100 per page if you have more than 100
	// tags.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTagsPaginator is a paginator for DescribeTags
type DescribeTagsPaginator struct {
	options   DescribeTagsPaginatorOptions
	client    DescribeTagsAPIClient
	params    *DescribeTagsInput
	nextToken *string
	firstPage bool
}

// NewDescribeTagsPaginator returns a new DescribeTagsPaginator
func NewDescribeTagsPaginator(client DescribeTagsAPIClient, params *DescribeTagsInput, optFns ...func(*DescribeTagsPaginatorOptions)) *DescribeTagsPaginator {
	if params == nil {
		params = &DescribeTagsInput{}
	}

	options := DescribeTagsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTagsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTagsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTags page.
func (p *DescribeTagsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTagsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.DescribeTags(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "DescribeTags",
	}
}
