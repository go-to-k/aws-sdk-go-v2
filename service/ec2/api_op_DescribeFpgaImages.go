// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the Amazon FPGA Images (AFIs) available to you. These include public
// AFIs, private AFIs that you own, and AFIs owned by other Amazon Web Services
// accounts for which you have load permissions.
func (c *Client) DescribeFpgaImages(ctx context.Context, params *DescribeFpgaImagesInput, optFns ...func(*Options)) (*DescribeFpgaImagesOutput, error) {
	if params == nil {
		params = &DescribeFpgaImagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFpgaImages", params, optFns, c.addOperationDescribeFpgaImagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFpgaImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFpgaImagesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The filters.
	//   - create-time - The creation time of the AFI.
	//   - fpga-image-id - The FPGA image identifier (AFI ID).
	//   - fpga-image-global-id - The global FPGA image identifier (AGFI ID).
	//   - name - The name of the AFI.
	//   - owner-id - The Amazon Web Services account ID of the AFI owner.
	//   - product-code - The product code.
	//   - shell-version - The version of the Amazon Web Services Shell that was used
	//   to create the bitstream.
	//   - state - The state of the AFI ( pending | failed | available | unavailable ).
	//   - tag : - The key/value combination of a tag assigned to the resource. Use the
	//   tag key in the filter name and the tag value as the filter value. For example,
	//   to find all resources that have a tag with the key Owner and the value TeamA ,
	//   specify tag:Owner for the filter name and TeamA for the filter value.
	//   - tag-key - The key of a tag assigned to the resource. Use this filter to find
	//   all resources assigned a tag with a specific key, regardless of the tag value.
	//   - update-time - The time of the most recent update.
	Filters []types.Filter

	// The AFI IDs.
	FpgaImageIds []string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token to retrieve the next page of results.
	NextToken *string

	// Filters the AFI by owner. Specify an Amazon Web Services account ID, self
	// (owner is the sender of the request), or an Amazon Web Services owner alias
	// (valid values are amazon | aws-marketplace ).
	Owners []string

	noSmithyDocumentSerde
}

type DescribeFpgaImagesOutput struct {

	// Information about the FPGA images.
	FpgaImages []types.FpgaImage

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFpgaImagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeFpgaImages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeFpgaImages{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFpgaImages(options.Region), middleware.Before); err != nil {
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

// DescribeFpgaImagesAPIClient is a client that implements the DescribeFpgaImages
// operation.
type DescribeFpgaImagesAPIClient interface {
	DescribeFpgaImages(context.Context, *DescribeFpgaImagesInput, ...func(*Options)) (*DescribeFpgaImagesOutput, error)
}

var _ DescribeFpgaImagesAPIClient = (*Client)(nil)

// DescribeFpgaImagesPaginatorOptions is the paginator options for
// DescribeFpgaImages
type DescribeFpgaImagesPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFpgaImagesPaginator is a paginator for DescribeFpgaImages
type DescribeFpgaImagesPaginator struct {
	options   DescribeFpgaImagesPaginatorOptions
	client    DescribeFpgaImagesAPIClient
	params    *DescribeFpgaImagesInput
	nextToken *string
	firstPage bool
}

// NewDescribeFpgaImagesPaginator returns a new DescribeFpgaImagesPaginator
func NewDescribeFpgaImagesPaginator(client DescribeFpgaImagesAPIClient, params *DescribeFpgaImagesInput, optFns ...func(*DescribeFpgaImagesPaginatorOptions)) *DescribeFpgaImagesPaginator {
	if params == nil {
		params = &DescribeFpgaImagesInput{}
	}

	options := DescribeFpgaImagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFpgaImagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFpgaImagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFpgaImages page.
func (p *DescribeFpgaImagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFpgaImagesOutput, error) {
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

	result, err := p.client.DescribeFpgaImages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeFpgaImages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeFpgaImages",
	}
}
