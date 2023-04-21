// Code generated by smithy-go-codegen DO NOT EDIT.

package healthlake

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/healthlake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists all FHIR import jobs associated with an account and their statuses.
func (c *Client) ListFHIRImportJobs(ctx context.Context, params *ListFHIRImportJobsInput, optFns ...func(*Options)) (*ListFHIRImportJobsOutput, error) {
	if params == nil {
		params = &ListFHIRImportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFHIRImportJobs", params, optFns, c.addOperationListFHIRImportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFHIRImportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFHIRImportJobsInput struct {

	// This parameter limits the response to the import job with the specified Data
	// Store ID.
	//
	// This member is required.
	DatastoreId *string

	// This parameter limits the response to the import job with the specified job
	// name.
	JobName *string

	// This parameter limits the response to the import job with the specified job
	// status.
	JobStatus types.JobStatus

	// This parameter limits the number of results returned for a ListFHIRImportJobs
	// to a maximum quantity specified by the user.
	MaxResults *int32

	// A pagination token used to identify the next page of results to return for a
	// ListFHIRImportJobs query.
	NextToken *string

	// This parameter limits the response to FHIR import jobs submitted after a user
	// specified date.
	SubmittedAfter *time.Time

	// This parameter limits the response to FHIR import jobs submitted before a user
	// specified date.
	SubmittedBefore *time.Time

	noSmithyDocumentSerde
}

type ListFHIRImportJobsOutput struct {

	// The properties of a listed FHIR import jobs, including the ID, ARN, name, and
	// the status of the job.
	//
	// This member is required.
	ImportJobPropertiesList []types.ImportJobProperties

	// A pagination token used to identify the next page of results to return for a
	// ListFHIRImportJobs query.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFHIRImportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListFHIRImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListFHIRImportJobs{}, middleware.After)
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
	if err = addOpListFHIRImportJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFHIRImportJobs(options.Region), middleware.Before); err != nil {
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

// ListFHIRImportJobsAPIClient is a client that implements the ListFHIRImportJobs
// operation.
type ListFHIRImportJobsAPIClient interface {
	ListFHIRImportJobs(context.Context, *ListFHIRImportJobsInput, ...func(*Options)) (*ListFHIRImportJobsOutput, error)
}

var _ ListFHIRImportJobsAPIClient = (*Client)(nil)

// ListFHIRImportJobsPaginatorOptions is the paginator options for
// ListFHIRImportJobs
type ListFHIRImportJobsPaginatorOptions struct {
	// This parameter limits the number of results returned for a ListFHIRImportJobs
	// to a maximum quantity specified by the user.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFHIRImportJobsPaginator is a paginator for ListFHIRImportJobs
type ListFHIRImportJobsPaginator struct {
	options   ListFHIRImportJobsPaginatorOptions
	client    ListFHIRImportJobsAPIClient
	params    *ListFHIRImportJobsInput
	nextToken *string
	firstPage bool
}

// NewListFHIRImportJobsPaginator returns a new ListFHIRImportJobsPaginator
func NewListFHIRImportJobsPaginator(client ListFHIRImportJobsAPIClient, params *ListFHIRImportJobsInput, optFns ...func(*ListFHIRImportJobsPaginatorOptions)) *ListFHIRImportJobsPaginator {
	if params == nil {
		params = &ListFHIRImportJobsInput{}
	}

	options := ListFHIRImportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFHIRImportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFHIRImportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFHIRImportJobs page.
func (p *ListFHIRImportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFHIRImportJobsOutput, error) {
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

	result, err := p.client.ListFHIRImportJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFHIRImportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "healthlake",
		OperationName: "ListFHIRImportJobs",
	}
}
