// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all readiness rules, or lists the readiness rules for a specific resource
// type.
func (c *Client) ListRules(ctx context.Context, params *ListRulesInput, optFns ...func(*Options)) (*ListRulesOutput, error) {
	if params == nil {
		params = &ListRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRules", params, optFns, c.addOperationListRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRulesInput struct {

	// The number of objects that you want to return with this call.
	MaxResults int32

	// The token that identifies which batch of results you want to see.
	NextToken *string

	// The resource type that a readiness rule applies to.
	ResourceType *string

	noSmithyDocumentSerde
}

type ListRulesOutput struct {

	// The token that identifies which batch of results you want to see.
	NextToken *string

	// A list of readiness rules for a specific resource type.
	Rules []types.ListRulesOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRules{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRules(options.Region), middleware.Before); err != nil {
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

// ListRulesAPIClient is a client that implements the ListRules operation.
type ListRulesAPIClient interface {
	ListRules(context.Context, *ListRulesInput, ...func(*Options)) (*ListRulesOutput, error)
}

var _ ListRulesAPIClient = (*Client)(nil)

// ListRulesPaginatorOptions is the paginator options for ListRules
type ListRulesPaginatorOptions struct {
	// The number of objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRulesPaginator is a paginator for ListRules
type ListRulesPaginator struct {
	options   ListRulesPaginatorOptions
	client    ListRulesAPIClient
	params    *ListRulesInput
	nextToken *string
	firstPage bool
}

// NewListRulesPaginator returns a new ListRulesPaginator
func NewListRulesPaginator(client ListRulesAPIClient, params *ListRulesInput, optFns ...func(*ListRulesPaginatorOptions)) *ListRulesPaginator {
	if params == nil {
		params = &ListRulesInput{}
	}

	options := ListRulesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRules page.
func (p *ListRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRulesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListRules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-readiness",
		OperationName: "ListRules",
	}
}
