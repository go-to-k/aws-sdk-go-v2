// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation returns a list of the ARNs of the canaries that are associated
// with the specified group.
func (c *Client) ListGroupResources(ctx context.Context, params *ListGroupResourcesInput, optFns ...func(*Options)) (*ListGroupResourcesOutput, error) {
	if params == nil {
		params = &ListGroupResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGroupResources", params, optFns, c.addOperationListGroupResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGroupResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroupResourcesInput struct {

	// Specifies the group to return information for. You can specify the group name,
	// the ARN, or the group ID as the GroupIdentifier.
	//
	// This member is required.
	GroupIdentifier *string

	// Specify this parameter to limit how many canary ARNs are returned each time you
	// use the ListGroupResources operation. If you omit this parameter, the default of
	// 20 is used.
	MaxResults *int32

	// A token that indicates that there is more data available. You can use this token
	// in a subsequent operation to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGroupResourcesOutput struct {

	// A token that indicates that there is more data available. You can use this token
	// in a subsequent ListGroupResources operation to retrieve the next set of
	// results.
	NextToken *string

	// An array of ARNs. These ARNs are for the canaries that are associated with the
	// group.
	Resources []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGroupResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListGroupResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListGroupResources{}, middleware.After)
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
	if err = addOpListGroupResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGroupResources(options.Region), middleware.Before); err != nil {
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

// ListGroupResourcesAPIClient is a client that implements the ListGroupResources
// operation.
type ListGroupResourcesAPIClient interface {
	ListGroupResources(context.Context, *ListGroupResourcesInput, ...func(*Options)) (*ListGroupResourcesOutput, error)
}

var _ ListGroupResourcesAPIClient = (*Client)(nil)

// ListGroupResourcesPaginatorOptions is the paginator options for
// ListGroupResources
type ListGroupResourcesPaginatorOptions struct {
	// Specify this parameter to limit how many canary ARNs are returned each time you
	// use the ListGroupResources operation. If you omit this parameter, the default of
	// 20 is used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGroupResourcesPaginator is a paginator for ListGroupResources
type ListGroupResourcesPaginator struct {
	options   ListGroupResourcesPaginatorOptions
	client    ListGroupResourcesAPIClient
	params    *ListGroupResourcesInput
	nextToken *string
	firstPage bool
}

// NewListGroupResourcesPaginator returns a new ListGroupResourcesPaginator
func NewListGroupResourcesPaginator(client ListGroupResourcesAPIClient, params *ListGroupResourcesInput, optFns ...func(*ListGroupResourcesPaginatorOptions)) *ListGroupResourcesPaginator {
	if params == nil {
		params = &ListGroupResourcesInput{}
	}

	options := ListGroupResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGroupResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGroupResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGroupResources page.
func (p *ListGroupResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGroupResourcesOutput, error) {
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

	result, err := p.client.ListGroupResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGroupResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "synthetics",
		OperationName: "ListGroupResources",
	}
}
