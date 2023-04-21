// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a description of the specified Amazon Resource Name (ARN) of virtual
// tapes. If a TapeARN is not specified, returns a description of all virtual
// tapes associated with the specified gateway. This operation is only supported in
// the tape gateway type.
func (c *Client) DescribeTapes(ctx context.Context, params *DescribeTapesInput, optFns ...func(*Options)) (*DescribeTapesOutput, error) {
	if params == nil {
		params = &DescribeTapesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTapes", params, optFns, c.addOperationDescribeTapesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTapesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeTapesInput
type DescribeTapesInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	// Specifies that the number of virtual tapes described be limited to the
	// specified number. Amazon Web Services may impose its own limit, if this field is
	// not set.
	Limit *int32

	// A marker value, obtained in a previous call to DescribeTapes . This marker
	// indicates which page of results to retrieve. If not specified, the first page of
	// results is retrieved.
	Marker *string

	// Specifies one or more unique Amazon Resource Names (ARNs) that represent the
	// virtual tapes you want to describe. If this parameter is not specified, Tape
	// gateway returns a description of all virtual tapes associated with the specified
	// gateway.
	TapeARNs []string

	noSmithyDocumentSerde
}

// DescribeTapesOutput
type DescribeTapesOutput struct {

	// An opaque string that can be used as part of a subsequent DescribeTapes call to
	// retrieve the next page of results. If a response does not contain a marker, then
	// there are no more results to be retrieved.
	Marker *string

	// An array of virtual tape descriptions.
	Tapes []types.Tape

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTapesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTapes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTapes{}, middleware.After)
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
	if err = addOpDescribeTapesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTapes(options.Region), middleware.Before); err != nil {
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

// DescribeTapesAPIClient is a client that implements the DescribeTapes operation.
type DescribeTapesAPIClient interface {
	DescribeTapes(context.Context, *DescribeTapesInput, ...func(*Options)) (*DescribeTapesOutput, error)
}

var _ DescribeTapesAPIClient = (*Client)(nil)

// DescribeTapesPaginatorOptions is the paginator options for DescribeTapes
type DescribeTapesPaginatorOptions struct {
	// Specifies that the number of virtual tapes described be limited to the
	// specified number. Amazon Web Services may impose its own limit, if this field is
	// not set.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTapesPaginator is a paginator for DescribeTapes
type DescribeTapesPaginator struct {
	options   DescribeTapesPaginatorOptions
	client    DescribeTapesAPIClient
	params    *DescribeTapesInput
	nextToken *string
	firstPage bool
}

// NewDescribeTapesPaginator returns a new DescribeTapesPaginator
func NewDescribeTapesPaginator(client DescribeTapesAPIClient, params *DescribeTapesInput, optFns ...func(*DescribeTapesPaginatorOptions)) *DescribeTapesPaginator {
	if params == nil {
		params = &DescribeTapesInput{}
	}

	options := DescribeTapesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTapesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTapesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTapes page.
func (p *DescribeTapesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTapesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeTapes(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeTapes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeTapes",
	}
}
