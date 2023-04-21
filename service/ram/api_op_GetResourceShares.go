// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves details about the resource shares that you own or that are shared
// with you.
func (c *Client) GetResourceShares(ctx context.Context, params *GetResourceSharesInput, optFns ...func(*Options)) (*GetResourceSharesOutput, error) {
	if params == nil {
		params = &GetResourceSharesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourceShares", params, optFns, c.addOperationGetResourceSharesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourceSharesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourceSharesInput struct {

	// Specifies that you want to retrieve details of only those resource shares that
	// match the following:
	//   - SELF – resource shares that your account shares with other accounts
	//   - OTHER-ACCOUNTS – resource shares that other accounts share with your account
	//
	// This member is required.
	ResourceOwner types.ResourceOwner

	// Specifies the total number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value that
	// is specific to the operation. If additional items exist beyond the number you
	// specify, the NextToken response element is returned with a value (not null).
	// Include the specified value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32

	// Specifies the name of an individual resource share that you want to retrieve
	// details about.
	Name *string

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a NextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's NextToken response to request the next page of results.
	NextToken *string

	// Specifies that you want to retrieve details of only those resource shares that
	// use the RAM permission with this Amazon Resoure Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// .
	PermissionArn *string

	// Specifies the Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of individual resource shares that you want information about.
	ResourceShareArns []string

	// Specifies that you want to retrieve details of only those resource shares that
	// have this status.
	ResourceShareStatus types.ResourceShareStatus

	// Specifies that you want to retrieve details of only those resource shares that
	// match the specified tag keys and values.
	TagFilters []types.TagFilter

	noSmithyDocumentSerde
}

type GetResourceSharesOutput struct {

	// If present, this value indicates that more output is available than is included
	// in the current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null . This
	// indicates that this is the last page of results.
	NextToken *string

	// An array of objects that contain the information about the resource shares.
	ResourceShares []types.ResourceShare

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourceSharesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetResourceShares{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetResourceShares{}, middleware.After)
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
	if err = addOpGetResourceSharesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourceShares(options.Region), middleware.Before); err != nil {
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

// GetResourceSharesAPIClient is a client that implements the GetResourceShares
// operation.
type GetResourceSharesAPIClient interface {
	GetResourceShares(context.Context, *GetResourceSharesInput, ...func(*Options)) (*GetResourceSharesOutput, error)
}

var _ GetResourceSharesAPIClient = (*Client)(nil)

// GetResourceSharesPaginatorOptions is the paginator options for GetResourceShares
type GetResourceSharesPaginatorOptions struct {
	// Specifies the total number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value that
	// is specific to the operation. If additional items exist beyond the number you
	// specify, the NextToken response element is returned with a value (not null).
	// Include the specified value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that the service
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetResourceSharesPaginator is a paginator for GetResourceShares
type GetResourceSharesPaginator struct {
	options   GetResourceSharesPaginatorOptions
	client    GetResourceSharesAPIClient
	params    *GetResourceSharesInput
	nextToken *string
	firstPage bool
}

// NewGetResourceSharesPaginator returns a new GetResourceSharesPaginator
func NewGetResourceSharesPaginator(client GetResourceSharesAPIClient, params *GetResourceSharesInput, optFns ...func(*GetResourceSharesPaginatorOptions)) *GetResourceSharesPaginator {
	if params == nil {
		params = &GetResourceSharesInput{}
	}

	options := GetResourceSharesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetResourceSharesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetResourceSharesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetResourceShares page.
func (p *GetResourceSharesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetResourceSharesOutput, error) {
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

	result, err := p.client.GetResourceShares(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetResourceShares(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "GetResourceShares",
	}
}
