// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the status of the AWS account assignment creation requests for a
// specified IAM Identity Center instance.
func (c *Client) ListAccountAssignmentCreationStatus(ctx context.Context, params *ListAccountAssignmentCreationStatusInput, optFns ...func(*Options)) (*ListAccountAssignmentCreationStatusOutput, error) {
	if params == nil {
		params = &ListAccountAssignmentCreationStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccountAssignmentCreationStatus", params, optFns, c.addOperationListAccountAssignmentCreationStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccountAssignmentCreationStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccountAssignmentCreationStatusInput struct {

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed. For more information about ARNs, see Amazon Resource Names (ARNs) and
	// AWS Service Namespaces in the AWS General Reference.
	//
	// This member is required.
	InstanceArn *string

	// Filters results based on the passed attribute value.
	Filter *types.OperationStatusFilter

	// The maximum number of results to display for the assignment.
	MaxResults *int32

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAccountAssignmentCreationStatusOutput struct {

	// The status object for the account assignment creation operation.
	AccountAssignmentsCreationStatus []types.AccountAssignmentOperationStatusMetadata

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccountAssignmentCreationStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAccountAssignmentCreationStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAccountAssignmentCreationStatus{}, middleware.After)
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
	if err = addOpListAccountAssignmentCreationStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccountAssignmentCreationStatus(options.Region), middleware.Before); err != nil {
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

// ListAccountAssignmentCreationStatusAPIClient is a client that implements the
// ListAccountAssignmentCreationStatus operation.
type ListAccountAssignmentCreationStatusAPIClient interface {
	ListAccountAssignmentCreationStatus(context.Context, *ListAccountAssignmentCreationStatusInput, ...func(*Options)) (*ListAccountAssignmentCreationStatusOutput, error)
}

var _ ListAccountAssignmentCreationStatusAPIClient = (*Client)(nil)

// ListAccountAssignmentCreationStatusPaginatorOptions is the paginator options
// for ListAccountAssignmentCreationStatus
type ListAccountAssignmentCreationStatusPaginatorOptions struct {
	// The maximum number of results to display for the assignment.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccountAssignmentCreationStatusPaginator is a paginator for
// ListAccountAssignmentCreationStatus
type ListAccountAssignmentCreationStatusPaginator struct {
	options   ListAccountAssignmentCreationStatusPaginatorOptions
	client    ListAccountAssignmentCreationStatusAPIClient
	params    *ListAccountAssignmentCreationStatusInput
	nextToken *string
	firstPage bool
}

// NewListAccountAssignmentCreationStatusPaginator returns a new
// ListAccountAssignmentCreationStatusPaginator
func NewListAccountAssignmentCreationStatusPaginator(client ListAccountAssignmentCreationStatusAPIClient, params *ListAccountAssignmentCreationStatusInput, optFns ...func(*ListAccountAssignmentCreationStatusPaginatorOptions)) *ListAccountAssignmentCreationStatusPaginator {
	if params == nil {
		params = &ListAccountAssignmentCreationStatusInput{}
	}

	options := ListAccountAssignmentCreationStatusPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccountAssignmentCreationStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccountAssignmentCreationStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccountAssignmentCreationStatus page.
func (p *ListAccountAssignmentCreationStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccountAssignmentCreationStatusOutput, error) {
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

	result, err := p.client.ListAccountAssignmentCreationStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAccountAssignmentCreationStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "ListAccountAssignmentCreationStatus",
	}
}
