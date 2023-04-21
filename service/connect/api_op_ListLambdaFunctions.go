// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Returns a paginated list of all Lambda functions that display in the dropdown
// options in the relevant flow blocks.
func (c *Client) ListLambdaFunctions(ctx context.Context, params *ListLambdaFunctionsInput, optFns ...func(*Options)) (*ListLambdaFunctionsOutput, error) {
	if params == nil {
		params = &ListLambdaFunctionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLambdaFunctions", params, optFns, c.addOperationListLambdaFunctionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLambdaFunctionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLambdaFunctionsInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLambdaFunctionsOutput struct {

	// The Lambdafunction ARNs associated with the specified instance.
	LambdaFunctions []string

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLambdaFunctionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLambdaFunctions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLambdaFunctions{}, middleware.After)
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
	if err = addOpListLambdaFunctionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLambdaFunctions(options.Region), middleware.Before); err != nil {
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

// ListLambdaFunctionsAPIClient is a client that implements the
// ListLambdaFunctions operation.
type ListLambdaFunctionsAPIClient interface {
	ListLambdaFunctions(context.Context, *ListLambdaFunctionsInput, ...func(*Options)) (*ListLambdaFunctionsOutput, error)
}

var _ ListLambdaFunctionsAPIClient = (*Client)(nil)

// ListLambdaFunctionsPaginatorOptions is the paginator options for
// ListLambdaFunctions
type ListLambdaFunctionsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLambdaFunctionsPaginator is a paginator for ListLambdaFunctions
type ListLambdaFunctionsPaginator struct {
	options   ListLambdaFunctionsPaginatorOptions
	client    ListLambdaFunctionsAPIClient
	params    *ListLambdaFunctionsInput
	nextToken *string
	firstPage bool
}

// NewListLambdaFunctionsPaginator returns a new ListLambdaFunctionsPaginator
func NewListLambdaFunctionsPaginator(client ListLambdaFunctionsAPIClient, params *ListLambdaFunctionsInput, optFns ...func(*ListLambdaFunctionsPaginatorOptions)) *ListLambdaFunctionsPaginator {
	if params == nil {
		params = &ListLambdaFunctionsInput{}
	}

	options := ListLambdaFunctionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLambdaFunctionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLambdaFunctionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLambdaFunctions page.
func (p *ListLambdaFunctionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLambdaFunctionsOutput, error) {
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

	result, err := p.client.ListLambdaFunctions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLambdaFunctions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListLambdaFunctions",
	}
}
