// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more of your compute environments. If you're using an
// unmanaged compute environment, you can use the DescribeComputeEnvironment
// operation to determine the ecsClusterArn that you launch your Amazon ECS
// container instances into.
func (c *Client) DescribeComputeEnvironments(ctx context.Context, params *DescribeComputeEnvironmentsInput, optFns ...func(*Options)) (*DescribeComputeEnvironmentsOutput, error) {
	if params == nil {
		params = &DescribeComputeEnvironmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeComputeEnvironments", params, optFns, c.addOperationDescribeComputeEnvironmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeComputeEnvironmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeComputeEnvironments .
type DescribeComputeEnvironmentsInput struct {

	// A list of up to 100 compute environment names or full Amazon Resource Name
	// (ARN) entries.
	ComputeEnvironments []string

	// The maximum number of cluster results returned by DescribeComputeEnvironments
	// in paginated output. When this parameter is used, DescribeComputeEnvironments
	// only returns maxResults results in a single page along with a nextToken
	// response element. The remaining results of the initial request can be seen by
	// sending another DescribeComputeEnvironments request with the returned nextToken
	// value. This value can be between 1 and 100. If this parameter isn't used, then
	// DescribeComputeEnvironments returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated
	// DescribeComputeEnvironments request where maxResults was used and the results
	// exceeded the value of that parameter. Pagination continues from the end of the
	// previous results that returned the nextToken value. This value is null when
	// there are no more results to return. Treat this token as an opaque identifier
	// that's only used to retrieve the next items in a list and not for other
	// programmatic purposes.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeComputeEnvironmentsOutput struct {

	// The list of compute environments.
	ComputeEnvironments []types.ComputeEnvironmentDetail

	// The nextToken value to include in a future DescribeComputeEnvironments request.
	// When the results of a DescribeComputeEnvironments request exceed maxResults ,
	// this value can be used to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeComputeEnvironmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeComputeEnvironments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeComputeEnvironments{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeComputeEnvironments(options.Region), middleware.Before); err != nil {
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

// DescribeComputeEnvironmentsAPIClient is a client that implements the
// DescribeComputeEnvironments operation.
type DescribeComputeEnvironmentsAPIClient interface {
	DescribeComputeEnvironments(context.Context, *DescribeComputeEnvironmentsInput, ...func(*Options)) (*DescribeComputeEnvironmentsOutput, error)
}

var _ DescribeComputeEnvironmentsAPIClient = (*Client)(nil)

// DescribeComputeEnvironmentsPaginatorOptions is the paginator options for
// DescribeComputeEnvironments
type DescribeComputeEnvironmentsPaginatorOptions struct {
	// The maximum number of cluster results returned by DescribeComputeEnvironments
	// in paginated output. When this parameter is used, DescribeComputeEnvironments
	// only returns maxResults results in a single page along with a nextToken
	// response element. The remaining results of the initial request can be seen by
	// sending another DescribeComputeEnvironments request with the returned nextToken
	// value. This value can be between 1 and 100. If this parameter isn't used, then
	// DescribeComputeEnvironments returns up to 100 results and a nextToken value if
	// applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeComputeEnvironmentsPaginator is a paginator for
// DescribeComputeEnvironments
type DescribeComputeEnvironmentsPaginator struct {
	options   DescribeComputeEnvironmentsPaginatorOptions
	client    DescribeComputeEnvironmentsAPIClient
	params    *DescribeComputeEnvironmentsInput
	nextToken *string
	firstPage bool
}

// NewDescribeComputeEnvironmentsPaginator returns a new
// DescribeComputeEnvironmentsPaginator
func NewDescribeComputeEnvironmentsPaginator(client DescribeComputeEnvironmentsAPIClient, params *DescribeComputeEnvironmentsInput, optFns ...func(*DescribeComputeEnvironmentsPaginatorOptions)) *DescribeComputeEnvironmentsPaginator {
	if params == nil {
		params = &DescribeComputeEnvironmentsInput{}
	}

	options := DescribeComputeEnvironmentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeComputeEnvironmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeComputeEnvironmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeComputeEnvironments page.
func (p *DescribeComputeEnvironmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeComputeEnvironmentsOutput, error) {
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

	result, err := p.client.DescribeComputeEnvironments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeComputeEnvironments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "DescribeComputeEnvironments",
	}
}
