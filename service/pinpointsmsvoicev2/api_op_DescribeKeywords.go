// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified keywords or all keywords on your origination phone
// number or pool. A keyword is a word that you can search for on a particular
// phone number or pool. It is also a specific word or phrase that an end user can
// send to your number to elicit a response, such as an informational message or a
// special offer. When your number receives a message that begins with a keyword,
// Amazon Pinpoint responds with a customizable message. If you specify a keyword
// that isn't valid, an Error is returned.
func (c *Client) DescribeKeywords(ctx context.Context, params *DescribeKeywordsInput, optFns ...func(*Options)) (*DescribeKeywordsOutput, error) {
	if params == nil {
		params = &DescribeKeywordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeKeywords", params, optFns, c.addOperationDescribeKeywordsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeKeywordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeKeywordsInput struct {

	// The origination identity to use such as a PhoneNumberId, PhoneNumberArn,
	// SenderId or SenderIdArn. You can use DescribePhoneNumbers to find the values
	// for PhoneNumberId and PhoneNumberArn while DescribeSenderIds can be used to get
	// the values for SenderId and SenderIdArn.
	//
	// This member is required.
	OriginationIdentity *string

	// An array of keyword filters to filter the results.
	Filters []types.KeywordFilter

	// An array of keywords to search for.
	Keywords []string

	// The maximum number of results to return per each request.
	MaxResults *int32

	// The token to be used for the next set of paginated results. You don't need to
	// supply a value for this field in the initial request.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeKeywordsOutput struct {

	// An array of KeywordInformation objects that contain the results.
	Keywords []types.KeywordInformation

	// The token to be used for the next set of paginated results. If this field is
	// empty then there are no more results.
	NextToken *string

	// The PhoneNumberId or PoolId that is associated with the OriginationIdentity.
	OriginationIdentity *string

	// The PhoneNumberArn or PoolArn that is associated with the OriginationIdentity.
	OriginationIdentityArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeKeywordsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeKeywords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeKeywords{}, middleware.After)
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
	if err = addOpDescribeKeywordsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeKeywords(options.Region), middleware.Before); err != nil {
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

// DescribeKeywordsAPIClient is a client that implements the DescribeKeywords
// operation.
type DescribeKeywordsAPIClient interface {
	DescribeKeywords(context.Context, *DescribeKeywordsInput, ...func(*Options)) (*DescribeKeywordsOutput, error)
}

var _ DescribeKeywordsAPIClient = (*Client)(nil)

// DescribeKeywordsPaginatorOptions is the paginator options for DescribeKeywords
type DescribeKeywordsPaginatorOptions struct {
	// The maximum number of results to return per each request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeKeywordsPaginator is a paginator for DescribeKeywords
type DescribeKeywordsPaginator struct {
	options   DescribeKeywordsPaginatorOptions
	client    DescribeKeywordsAPIClient
	params    *DescribeKeywordsInput
	nextToken *string
	firstPage bool
}

// NewDescribeKeywordsPaginator returns a new DescribeKeywordsPaginator
func NewDescribeKeywordsPaginator(client DescribeKeywordsAPIClient, params *DescribeKeywordsInput, optFns ...func(*DescribeKeywordsPaginatorOptions)) *DescribeKeywordsPaginator {
	if params == nil {
		params = &DescribeKeywordsInput{}
	}

	options := DescribeKeywordsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeKeywordsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeKeywordsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeKeywords page.
func (p *DescribeKeywordsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeKeywordsOutput, error) {
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

	result, err := p.client.DescribeKeywords(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeKeywords(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "DescribeKeywords",
	}
}
