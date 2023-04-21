// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarconnections

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codestarconnections/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the hosts associated with your account.
func (c *Client) ListHosts(ctx context.Context, params *ListHostsInput, optFns ...func(*Options)) (*ListHostsOutput, error) {
	if params == nil {
		params = &ListHostsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHosts", params, optFns, c.addOperationListHostsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHostsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHostsInput struct {

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults int32

	// The token that was returned from the previous ListHosts call, which can be used
	// to return the next set of hosts in the list.
	NextToken *string

	noSmithyDocumentSerde
}

type ListHostsOutput struct {

	// A list of hosts and the details for each host, such as status, endpoint, and
	// provider type.
	Hosts []types.Host

	// A token that can be used in the next ListHosts call. To view all items in the
	// list, continue to call this operation with each subsequent token until no more
	// nextToken values are returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHostsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListHosts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListHosts{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHosts(options.Region), middleware.Before); err != nil {
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

// ListHostsAPIClient is a client that implements the ListHosts operation.
type ListHostsAPIClient interface {
	ListHosts(context.Context, *ListHostsInput, ...func(*Options)) (*ListHostsOutput, error)
}

var _ ListHostsAPIClient = (*Client)(nil)

// ListHostsPaginatorOptions is the paginator options for ListHosts
type ListHostsPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListHostsPaginator is a paginator for ListHosts
type ListHostsPaginator struct {
	options   ListHostsPaginatorOptions
	client    ListHostsAPIClient
	params    *ListHostsInput
	nextToken *string
	firstPage bool
}

// NewListHostsPaginator returns a new ListHostsPaginator
func NewListHostsPaginator(client ListHostsAPIClient, params *ListHostsInput, optFns ...func(*ListHostsPaginatorOptions)) *ListHostsPaginator {
	if params == nil {
		params = &ListHostsInput{}
	}

	options := ListHostsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListHostsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListHostsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListHosts page.
func (p *ListHostsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListHostsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListHosts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListHosts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-connections",
		OperationName: "ListHosts",
	}
}
