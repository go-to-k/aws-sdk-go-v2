// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Device Defender security profiles you've created. You can filter
// security profiles by dimension or custom metric. Requires permission to access
// the ListSecurityProfiles (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action. dimensionName and metricName cannot be used in the same request.
func (c *Client) ListSecurityProfiles(ctx context.Context, params *ListSecurityProfilesInput, optFns ...func(*Options)) (*ListSecurityProfilesOutput, error) {
	if params == nil {
		params = &ListSecurityProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSecurityProfiles", params, optFns, c.addOperationListSecurityProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSecurityProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSecurityProfilesInput struct {

	// A filter to limit results to the security profiles that use the defined
	// dimension. Cannot be used with metricName
	DimensionName *string

	// The maximum number of results to return at one time.
	MaxResults *int32

	// The name of the custom metric. Cannot be used with dimensionName .
	MetricName *string

	// The token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSecurityProfilesOutput struct {

	// A token that can be used to retrieve the next set of results, or null if there
	// are no additional results.
	NextToken *string

	// A list of security profile identifiers (names and ARNs).
	SecurityProfileIdentifiers []types.SecurityProfileIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSecurityProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSecurityProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSecurityProfiles{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSecurityProfiles(options.Region), middleware.Before); err != nil {
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

// ListSecurityProfilesAPIClient is a client that implements the
// ListSecurityProfiles operation.
type ListSecurityProfilesAPIClient interface {
	ListSecurityProfiles(context.Context, *ListSecurityProfilesInput, ...func(*Options)) (*ListSecurityProfilesOutput, error)
}

var _ ListSecurityProfilesAPIClient = (*Client)(nil)

// ListSecurityProfilesPaginatorOptions is the paginator options for
// ListSecurityProfiles
type ListSecurityProfilesPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSecurityProfilesPaginator is a paginator for ListSecurityProfiles
type ListSecurityProfilesPaginator struct {
	options   ListSecurityProfilesPaginatorOptions
	client    ListSecurityProfilesAPIClient
	params    *ListSecurityProfilesInput
	nextToken *string
	firstPage bool
}

// NewListSecurityProfilesPaginator returns a new ListSecurityProfilesPaginator
func NewListSecurityProfilesPaginator(client ListSecurityProfilesAPIClient, params *ListSecurityProfilesInput, optFns ...func(*ListSecurityProfilesPaginatorOptions)) *ListSecurityProfilesPaginator {
	if params == nil {
		params = &ListSecurityProfilesInput{}
	}

	options := ListSecurityProfilesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSecurityProfilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSecurityProfilesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSecurityProfiles page.
func (p *ListSecurityProfilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSecurityProfilesOutput, error) {
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

	result, err := p.client.ListSecurityProfiles(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSecurityProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListSecurityProfiles",
	}
}
