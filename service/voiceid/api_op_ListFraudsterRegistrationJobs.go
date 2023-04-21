// Code generated by smithy-go-codegen DO NOT EDIT.

package voiceid

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/voiceid/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the fraudster registration jobs in the domain with the given JobStatus
// . If JobStatus is not provided, this lists all fraudster registration jobs in
// the given domain.
func (c *Client) ListFraudsterRegistrationJobs(ctx context.Context, params *ListFraudsterRegistrationJobsInput, optFns ...func(*Options)) (*ListFraudsterRegistrationJobsOutput, error) {
	if params == nil {
		params = &ListFraudsterRegistrationJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFraudsterRegistrationJobs", params, optFns, c.addOperationListFraudsterRegistrationJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFraudsterRegistrationJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFraudsterRegistrationJobsInput struct {

	// The identifier of the domain that contains the fraudster registration Jobs.
	//
	// This member is required.
	DomainId *string

	// Provides the status of your fraudster registration job.
	JobStatus types.FraudsterRegistrationJobStatus

	// The maximum number of results that are returned per call. You can use NextToken
	// to obtain more pages of results. The default is 100; the maximum allowed page
	// size is also 100.
	MaxResults *int32

	// If NextToken is returned, there are more results available. The value of
	// NextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFraudsterRegistrationJobsOutput struct {

	// A list containing details about each specified fraudster registration job.
	JobSummaries []types.FraudsterRegistrationJobSummary

	// If NextToken is returned, there are more results available. The value of
	// NextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFraudsterRegistrationJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListFraudsterRegistrationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListFraudsterRegistrationJobs{}, middleware.After)
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
	if err = addOpListFraudsterRegistrationJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFraudsterRegistrationJobs(options.Region), middleware.Before); err != nil {
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

// ListFraudsterRegistrationJobsAPIClient is a client that implements the
// ListFraudsterRegistrationJobs operation.
type ListFraudsterRegistrationJobsAPIClient interface {
	ListFraudsterRegistrationJobs(context.Context, *ListFraudsterRegistrationJobsInput, ...func(*Options)) (*ListFraudsterRegistrationJobsOutput, error)
}

var _ ListFraudsterRegistrationJobsAPIClient = (*Client)(nil)

// ListFraudsterRegistrationJobsPaginatorOptions is the paginator options for
// ListFraudsterRegistrationJobs
type ListFraudsterRegistrationJobsPaginatorOptions struct {
	// The maximum number of results that are returned per call. You can use NextToken
	// to obtain more pages of results. The default is 100; the maximum allowed page
	// size is also 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFraudsterRegistrationJobsPaginator is a paginator for
// ListFraudsterRegistrationJobs
type ListFraudsterRegistrationJobsPaginator struct {
	options   ListFraudsterRegistrationJobsPaginatorOptions
	client    ListFraudsterRegistrationJobsAPIClient
	params    *ListFraudsterRegistrationJobsInput
	nextToken *string
	firstPage bool
}

// NewListFraudsterRegistrationJobsPaginator returns a new
// ListFraudsterRegistrationJobsPaginator
func NewListFraudsterRegistrationJobsPaginator(client ListFraudsterRegistrationJobsAPIClient, params *ListFraudsterRegistrationJobsInput, optFns ...func(*ListFraudsterRegistrationJobsPaginatorOptions)) *ListFraudsterRegistrationJobsPaginator {
	if params == nil {
		params = &ListFraudsterRegistrationJobsInput{}
	}

	options := ListFraudsterRegistrationJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFraudsterRegistrationJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFraudsterRegistrationJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFraudsterRegistrationJobs page.
func (p *ListFraudsterRegistrationJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFraudsterRegistrationJobsOutput, error) {
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

	result, err := p.client.ListFraudsterRegistrationJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFraudsterRegistrationJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "voiceid",
		OperationName: "ListFraudsterRegistrationJobs",
	}
}
