// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Amazon EKS managed node groups associated with the specified cluster
// in your Amazon Web Services account in the specified Region. Self-managed node
// groups are not listed.
func (c *Client) ListNodegroups(ctx context.Context, params *ListNodegroupsInput, optFns ...func(*Options)) (*ListNodegroupsOutput, error) {
	if params == nil {
		params = &ListNodegroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNodegroups", params, optFns, c.addOperationListNodegroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNodegroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNodegroupsInput struct {

	// The name of the Amazon EKS cluster that you would like to list node groups in.
	//
	// This member is required.
	ClusterName *string

	// The maximum number of node group results returned by ListNodegroups in
	// paginated output. When you use this parameter, ListNodegroups returns only
	// maxResults results in a single page along with a nextToken response element.
	// You can see the remaining results of the initial request by sending another
	// ListNodegroups request with the returned nextToken value. This value can be
	// between 1 and 100. If you don't use this parameter, ListNodegroups returns up
	// to 100 results and a nextToken value if applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated ListNodegroups request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListNodegroupsOutput struct {

	// The nextToken value to include in a future ListNodegroups request. When the
	// results of a ListNodegroups request exceed maxResults , you can use this value
	// to retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// A list of all of the node groups associated with the specified cluster.
	Nodegroups []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNodegroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListNodegroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListNodegroups{}, middleware.After)
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
	if err = addOpListNodegroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNodegroups(options.Region), middleware.Before); err != nil {
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

// ListNodegroupsAPIClient is a client that implements the ListNodegroups
// operation.
type ListNodegroupsAPIClient interface {
	ListNodegroups(context.Context, *ListNodegroupsInput, ...func(*Options)) (*ListNodegroupsOutput, error)
}

var _ ListNodegroupsAPIClient = (*Client)(nil)

// ListNodegroupsPaginatorOptions is the paginator options for ListNodegroups
type ListNodegroupsPaginatorOptions struct {
	// The maximum number of node group results returned by ListNodegroups in
	// paginated output. When you use this parameter, ListNodegroups returns only
	// maxResults results in a single page along with a nextToken response element.
	// You can see the remaining results of the initial request by sending another
	// ListNodegroups request with the returned nextToken value. This value can be
	// between 1 and 100. If you don't use this parameter, ListNodegroups returns up
	// to 100 results and a nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNodegroupsPaginator is a paginator for ListNodegroups
type ListNodegroupsPaginator struct {
	options   ListNodegroupsPaginatorOptions
	client    ListNodegroupsAPIClient
	params    *ListNodegroupsInput
	nextToken *string
	firstPage bool
}

// NewListNodegroupsPaginator returns a new ListNodegroupsPaginator
func NewListNodegroupsPaginator(client ListNodegroupsAPIClient, params *ListNodegroupsInput, optFns ...func(*ListNodegroupsPaginatorOptions)) *ListNodegroupsPaginator {
	if params == nil {
		params = &ListNodegroupsInput{}
	}

	options := ListNodegroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNodegroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNodegroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNodegroups page.
func (p *ListNodegroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNodegroupsOutput, error) {
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

	result, err := p.client.ListNodegroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListNodegroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "ListNodegroups",
	}
}
