// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides cluster configuration information about the specified Elasticsearch
// domain, such as the state, creation date, update version, and update date for
// cluster options.
func (c *Client) DescribeElasticsearchDomainConfig(ctx context.Context, params *DescribeElasticsearchDomainConfigInput, optFns ...func(*Options)) (*DescribeElasticsearchDomainConfigOutput, error) {
	if params == nil {
		params = &DescribeElasticsearchDomainConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeElasticsearchDomainConfig", params, optFns, c.addOperationDescribeElasticsearchDomainConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeElasticsearchDomainConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeElasticsearchDomainConfig
// operation. Specifies the domain name for which you want configuration
// information.
type DescribeElasticsearchDomainConfigInput struct {

	// The Elasticsearch domain that you want to get information about.
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

// The result of a DescribeElasticsearchDomainConfig request. Contains the
// configuration information of the requested domain.
type DescribeElasticsearchDomainConfigOutput struct {

	// The configuration information of the domain requested in the
	// DescribeElasticsearchDomainConfig request.
	//
	// This member is required.
	DomainConfig *types.ElasticsearchDomainConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeElasticsearchDomainConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeElasticsearchDomainConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeElasticsearchDomainConfig{}, middleware.After)
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
	if err = addOpDescribeElasticsearchDomainConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeElasticsearchDomainConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeElasticsearchDomainConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeElasticsearchDomainConfig",
	}
}
