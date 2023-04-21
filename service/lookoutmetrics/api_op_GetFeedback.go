// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutmetrics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutmetrics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get feedback for an anomaly group.
func (c *Client) GetFeedback(ctx context.Context, params *GetFeedbackInput, optFns ...func(*Options)) (*GetFeedbackOutput, error) {
	if params == nil {
		params = &GetFeedbackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFeedback", params, optFns, c.addOperationGetFeedbackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFeedbackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFeedbackInput struct {

	// The Amazon Resource Name (ARN) of the anomaly detector.
	//
	// This member is required.
	AnomalyDetectorArn *string

	// The anomalous metric and group ID.
	//
	// This member is required.
	AnomalyGroupTimeSeriesFeedback *types.AnomalyGroupTimeSeries

	// The maximum number of results to return.
	MaxResults *int32

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetFeedbackOutput struct {

	// Feedback for an anomalous metric.
	AnomalyGroupTimeSeriesFeedback []types.TimeSeriesFeedback

	// The pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFeedbackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFeedback{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFeedback{}, middleware.After)
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
	if err = addOpGetFeedbackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFeedback(options.Region), middleware.Before); err != nil {
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

// GetFeedbackAPIClient is a client that implements the GetFeedback operation.
type GetFeedbackAPIClient interface {
	GetFeedback(context.Context, *GetFeedbackInput, ...func(*Options)) (*GetFeedbackOutput, error)
}

var _ GetFeedbackAPIClient = (*Client)(nil)

// GetFeedbackPaginatorOptions is the paginator options for GetFeedback
type GetFeedbackPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetFeedbackPaginator is a paginator for GetFeedback
type GetFeedbackPaginator struct {
	options   GetFeedbackPaginatorOptions
	client    GetFeedbackAPIClient
	params    *GetFeedbackInput
	nextToken *string
	firstPage bool
}

// NewGetFeedbackPaginator returns a new GetFeedbackPaginator
func NewGetFeedbackPaginator(client GetFeedbackAPIClient, params *GetFeedbackInput, optFns ...func(*GetFeedbackPaginatorOptions)) *GetFeedbackPaginator {
	if params == nil {
		params = &GetFeedbackInput{}
	}

	options := GetFeedbackPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetFeedbackPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetFeedbackPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetFeedback page.
func (p *GetFeedbackPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetFeedbackOutput, error) {
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

	result, err := p.client.GetFeedback(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetFeedback(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutmetrics",
		OperationName: "GetFeedback",
	}
}
