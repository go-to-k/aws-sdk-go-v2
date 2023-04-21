// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/synthetics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of runs for a specified canary.
func (c *Client) GetCanaryRuns(ctx context.Context, params *GetCanaryRunsInput, optFns ...func(*Options)) (*GetCanaryRunsOutput, error) {
	if params == nil {
		params = &GetCanaryRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCanaryRuns", params, optFns, c.addOperationGetCanaryRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCanaryRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCanaryRunsInput struct {

	// The name of the canary that you want to see runs for.
	//
	// This member is required.
	Name *string

	// Specify this parameter to limit how many runs are returned each time you use
	// the GetCanaryRuns operation. If you omit this parameter, the default of 100 is
	// used.
	MaxResults *int32

	// A token that indicates that there is more data available. You can use this
	// token in a subsequent GetCanaryRuns operation to retrieve the next set of
	// results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetCanaryRunsOutput struct {

	// An array of structures. Each structure contains the details of one of the
	// retrieved canary runs.
	CanaryRuns []types.CanaryRun

	// A token that indicates that there is more data available. You can use this
	// token in a subsequent GetCanaryRuns operation to retrieve the next set of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCanaryRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCanaryRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCanaryRuns{}, middleware.After)
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
	if err = addOpGetCanaryRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCanaryRuns(options.Region), middleware.Before); err != nil {
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

// GetCanaryRunsAPIClient is a client that implements the GetCanaryRuns operation.
type GetCanaryRunsAPIClient interface {
	GetCanaryRuns(context.Context, *GetCanaryRunsInput, ...func(*Options)) (*GetCanaryRunsOutput, error)
}

var _ GetCanaryRunsAPIClient = (*Client)(nil)

// GetCanaryRunsPaginatorOptions is the paginator options for GetCanaryRuns
type GetCanaryRunsPaginatorOptions struct {
	// Specify this parameter to limit how many runs are returned each time you use
	// the GetCanaryRuns operation. If you omit this parameter, the default of 100 is
	// used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCanaryRunsPaginator is a paginator for GetCanaryRuns
type GetCanaryRunsPaginator struct {
	options   GetCanaryRunsPaginatorOptions
	client    GetCanaryRunsAPIClient
	params    *GetCanaryRunsInput
	nextToken *string
	firstPage bool
}

// NewGetCanaryRunsPaginator returns a new GetCanaryRunsPaginator
func NewGetCanaryRunsPaginator(client GetCanaryRunsAPIClient, params *GetCanaryRunsInput, optFns ...func(*GetCanaryRunsPaginatorOptions)) *GetCanaryRunsPaginator {
	if params == nil {
		params = &GetCanaryRunsInput{}
	}

	options := GetCanaryRunsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetCanaryRunsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCanaryRunsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetCanaryRuns page.
func (p *GetCanaryRunsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCanaryRunsOutput, error) {
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

	result, err := p.client.GetCanaryRuns(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetCanaryRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "synthetics",
		OperationName: "GetCanaryRuns",
	}
}
