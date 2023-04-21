// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists coverage details for your GuardDuty account. If you're a GuardDuty
// administrator, you can retrieve all resources associated with the active member
// accounts in your organization. Make sure the accounts have EKS Runtime
// Monitoring enabled and GuardDuty agent running on their EKS nodes.
func (c *Client) ListCoverage(ctx context.Context, params *ListCoverageInput, optFns ...func(*Options)) (*ListCoverageOutput, error) {
	if params == nil {
		params = &ListCoverageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCoverage", params, optFns, c.addOperationListCoverageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCoverageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCoverageInput struct {

	// The unique ID of the detector whose coverage details you want to retrieve.
	//
	// This member is required.
	DetectorId *string

	// Represents the criteria used to filter the coverage details.
	FilterCriteria *types.CoverageFilterCriteria

	// The maximum number of results to return in the response.
	MaxResults int32

	// A token to use for paginating results that are returned in the response. Set
	// the value of this parameter to null for the first request to a list action. For
	// subsequent calls, use the NextToken value returned from the previous request to
	// continue listing results after the first page.
	NextToken *string

	// Represents the criteria used to sort the coverage details.
	SortCriteria *types.CoverageSortCriteria

	noSmithyDocumentSerde
}

type ListCoverageOutput struct {

	// A list of resources and their attributes providing cluster details.
	//
	// This member is required.
	Resources []types.CoverageResource

	// The pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCoverageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCoverage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCoverage{}, middleware.After)
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
	if err = addOpListCoverageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCoverage(options.Region), middleware.Before); err != nil {
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

// ListCoverageAPIClient is a client that implements the ListCoverage operation.
type ListCoverageAPIClient interface {
	ListCoverage(context.Context, *ListCoverageInput, ...func(*Options)) (*ListCoverageOutput, error)
}

var _ ListCoverageAPIClient = (*Client)(nil)

// ListCoveragePaginatorOptions is the paginator options for ListCoverage
type ListCoveragePaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCoveragePaginator is a paginator for ListCoverage
type ListCoveragePaginator struct {
	options   ListCoveragePaginatorOptions
	client    ListCoverageAPIClient
	params    *ListCoverageInput
	nextToken *string
	firstPage bool
}

// NewListCoveragePaginator returns a new ListCoveragePaginator
func NewListCoveragePaginator(client ListCoverageAPIClient, params *ListCoverageInput, optFns ...func(*ListCoveragePaginatorOptions)) *ListCoveragePaginator {
	if params == nil {
		params = &ListCoverageInput{}
	}

	options := ListCoveragePaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCoveragePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCoveragePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCoverage page.
func (p *ListCoveragePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCoverageOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListCoverage(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCoverage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "ListCoverage",
	}
}
