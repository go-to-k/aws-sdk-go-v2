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

// Returns a list of sent or received share requests for custom frameworks in
// Audit Manager.
func (c *Client) ListAssessmentFrameworkShareRequests(ctx context.Context, params *ListAssessmentFrameworkShareRequestsInput, optFns ...func(*Options)) (*ListAssessmentFrameworkShareRequestsOutput, error) {
	if params == nil {
		params = &ListAssessmentFrameworkShareRequestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssessmentFrameworkShareRequests", params, optFns, c.addOperationListAssessmentFrameworkShareRequestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssessmentFrameworkShareRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssessmentFrameworkShareRequestsInput struct {

	// Specifies whether the share request is a sent request or a received request.
	//
	// This member is required.
	RequestType types.ShareRequestType

	// Represents the maximum number of results on a page or for an API request call.
	MaxResults *int32

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssessmentFrameworkShareRequestsOutput struct {

	// The list of share requests that the ListAssessmentFrameworkShareRequests API
	// returned.
	AssessmentFrameworkShareRequests []types.AssessmentFrameworkShareRequest

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssessmentFrameworkShareRequestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssessmentFrameworkShareRequests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssessmentFrameworkShareRequests{}, middleware.After)
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
	if err = addOpListAssessmentFrameworkShareRequestsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssessmentFrameworkShareRequests(options.Region), middleware.Before); err != nil {
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

// ListAssessmentFrameworkShareRequestsAPIClient is a client that implements the
// ListAssessmentFrameworkShareRequests operation.
type ListAssessmentFrameworkShareRequestsAPIClient interface {
	ListAssessmentFrameworkShareRequests(context.Context, *ListAssessmentFrameworkShareRequestsInput, ...func(*Options)) (*ListAssessmentFrameworkShareRequestsOutput, error)
}

var _ ListAssessmentFrameworkShareRequestsAPIClient = (*Client)(nil)

// ListAssessmentFrameworkShareRequestsPaginatorOptions is the paginator options
// for ListAssessmentFrameworkShareRequests
type ListAssessmentFrameworkShareRequestsPaginatorOptions struct {
	// Represents the maximum number of results on a page or for an API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssessmentFrameworkShareRequestsPaginator is a paginator for
// ListAssessmentFrameworkShareRequests
type ListAssessmentFrameworkShareRequestsPaginator struct {
	options   ListAssessmentFrameworkShareRequestsPaginatorOptions
	client    ListAssessmentFrameworkShareRequestsAPIClient
	params    *ListAssessmentFrameworkShareRequestsInput
	nextToken *string
	firstPage bool
}

// NewListAssessmentFrameworkShareRequestsPaginator returns a new
// ListAssessmentFrameworkShareRequestsPaginator
func NewListAssessmentFrameworkShareRequestsPaginator(client ListAssessmentFrameworkShareRequestsAPIClient, params *ListAssessmentFrameworkShareRequestsInput, optFns ...func(*ListAssessmentFrameworkShareRequestsPaginatorOptions)) *ListAssessmentFrameworkShareRequestsPaginator {
	if params == nil {
		params = &ListAssessmentFrameworkShareRequestsInput{}
	}

	options := ListAssessmentFrameworkShareRequestsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssessmentFrameworkShareRequestsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssessmentFrameworkShareRequestsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssessmentFrameworkShareRequests page.
func (p *ListAssessmentFrameworkShareRequestsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssessmentFrameworkShareRequestsOutput, error) {
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

	result, err := p.client.ListAssessmentFrameworkShareRequests(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssessmentFrameworkShareRequests(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "ListAssessmentFrameworkShareRequests",
	}
}
