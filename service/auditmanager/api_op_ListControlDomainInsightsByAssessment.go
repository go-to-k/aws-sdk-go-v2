// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists analytics data for control domains within a specified active assessment.
// A control domain is listed only if at least one of the controls within that
// domain collected evidence on the lastUpdated date of controlDomainInsights . If
// this condition isn’t met, no data is listed for that domain.
func (c *Client) ListControlDomainInsightsByAssessment(ctx context.Context, params *ListControlDomainInsightsByAssessmentInput, optFns ...func(*Options)) (*ListControlDomainInsightsByAssessmentOutput, error) {
	if params == nil {
		params = &ListControlDomainInsightsByAssessmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListControlDomainInsightsByAssessment", params, optFns, c.addOperationListControlDomainInsightsByAssessmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListControlDomainInsightsByAssessmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListControlDomainInsightsByAssessmentInput struct {

	// The unique identifier for the active assessment.
	//
	// This member is required.
	AssessmentId *string

	// Represents the maximum number of results on a page or for an API request call.
	MaxResults *int32

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListControlDomainInsightsByAssessmentOutput struct {

	// The control domain analytics data that the ListControlDomainInsightsByAssessment
	// API returned.
	ControlDomainInsights []types.ControlDomainInsights

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListControlDomainInsightsByAssessmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListControlDomainInsightsByAssessment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListControlDomainInsightsByAssessment{}, middleware.After)
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
	if err = addOpListControlDomainInsightsByAssessmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListControlDomainInsightsByAssessment(options.Region), middleware.Before); err != nil {
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

// ListControlDomainInsightsByAssessmentAPIClient is a client that implements the
// ListControlDomainInsightsByAssessment operation.
type ListControlDomainInsightsByAssessmentAPIClient interface {
	ListControlDomainInsightsByAssessment(context.Context, *ListControlDomainInsightsByAssessmentInput, ...func(*Options)) (*ListControlDomainInsightsByAssessmentOutput, error)
}

var _ ListControlDomainInsightsByAssessmentAPIClient = (*Client)(nil)

// ListControlDomainInsightsByAssessmentPaginatorOptions is the paginator options
// for ListControlDomainInsightsByAssessment
type ListControlDomainInsightsByAssessmentPaginatorOptions struct {
	// Represents the maximum number of results on a page or for an API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListControlDomainInsightsByAssessmentPaginator is a paginator for
// ListControlDomainInsightsByAssessment
type ListControlDomainInsightsByAssessmentPaginator struct {
	options   ListControlDomainInsightsByAssessmentPaginatorOptions
	client    ListControlDomainInsightsByAssessmentAPIClient
	params    *ListControlDomainInsightsByAssessmentInput
	nextToken *string
	firstPage bool
}

// NewListControlDomainInsightsByAssessmentPaginator returns a new
// ListControlDomainInsightsByAssessmentPaginator
func NewListControlDomainInsightsByAssessmentPaginator(client ListControlDomainInsightsByAssessmentAPIClient, params *ListControlDomainInsightsByAssessmentInput, optFns ...func(*ListControlDomainInsightsByAssessmentPaginatorOptions)) *ListControlDomainInsightsByAssessmentPaginator {
	if params == nil {
		params = &ListControlDomainInsightsByAssessmentInput{}
	}

	options := ListControlDomainInsightsByAssessmentPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListControlDomainInsightsByAssessmentPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListControlDomainInsightsByAssessmentPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListControlDomainInsightsByAssessment page.
func (p *ListControlDomainInsightsByAssessmentPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListControlDomainInsightsByAssessmentOutput, error) {
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

	result, err := p.client.ListControlDomainInsightsByAssessment(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListControlDomainInsightsByAssessment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "ListControlDomainInsightsByAssessment",
	}
}
