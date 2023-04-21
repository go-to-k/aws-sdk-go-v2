// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the multicast groups registered to your AWS account.
func (c *Client) ListMulticastGroups(ctx context.Context, params *ListMulticastGroupsInput, optFns ...func(*Options)) (*ListMulticastGroupsOutput, error) {
	if params == nil {
		params = &ListMulticastGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMulticastGroups", params, optFns, c.addOperationListMulticastGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMulticastGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMulticastGroupsInput struct {

	// The maximum number of results to return in this operation.
	MaxResults int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMulticastGroupsOutput struct {

	// List of multicast groups.
	MulticastGroupList []types.MulticastGroup

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMulticastGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMulticastGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMulticastGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMulticastGroups(options.Region), middleware.Before); err != nil {
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

// ListMulticastGroupsAPIClient is a client that implements the
// ListMulticastGroups operation.
type ListMulticastGroupsAPIClient interface {
	ListMulticastGroups(context.Context, *ListMulticastGroupsInput, ...func(*Options)) (*ListMulticastGroupsOutput, error)
}

var _ ListMulticastGroupsAPIClient = (*Client)(nil)

// ListMulticastGroupsPaginatorOptions is the paginator options for
// ListMulticastGroups
type ListMulticastGroupsPaginatorOptions struct {
	// The maximum number of results to return in this operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMulticastGroupsPaginator is a paginator for ListMulticastGroups
type ListMulticastGroupsPaginator struct {
	options   ListMulticastGroupsPaginatorOptions
	client    ListMulticastGroupsAPIClient
	params    *ListMulticastGroupsInput
	nextToken *string
	firstPage bool
}

// NewListMulticastGroupsPaginator returns a new ListMulticastGroupsPaginator
func NewListMulticastGroupsPaginator(client ListMulticastGroupsAPIClient, params *ListMulticastGroupsInput, optFns ...func(*ListMulticastGroupsPaginatorOptions)) *ListMulticastGroupsPaginator {
	if params == nil {
		params = &ListMulticastGroupsInput{}
	}

	options := ListMulticastGroupsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMulticastGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMulticastGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMulticastGroups page.
func (p *ListMulticastGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMulticastGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListMulticastGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMulticastGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "ListMulticastGroups",
	}
}
