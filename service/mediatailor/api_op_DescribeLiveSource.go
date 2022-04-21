// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides details about a specific live source in a specific source location.
func (c *Client) DescribeLiveSource(ctx context.Context, params *DescribeLiveSourceInput, optFns ...func(*Options)) (*DescribeLiveSourceOutput, error) {
	if params == nil {
		params = &DescribeLiveSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLiveSource", params, optFns, c.addOperationDescribeLiveSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLiveSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLiveSourceInput struct {

	// The identifier for the live source you are working on.
	//
	// This member is required.
	LiveSourceName *string

	// The identifier for the source location you are working on.
	//
	// This member is required.
	SourceLocationName *string

	noSmithyDocumentSerde
}

type DescribeLiveSourceOutput struct {

	// The ARN of the live source.
	Arn *string

	// The timestamp that indicates when the live source was created.
	CreationTime *time.Time

	// The HTTP package configurations.
	HttpPackageConfigurations []types.HttpPackageConfiguration

	// The timestamp that indicates when the live source was modified.
	LastModifiedTime *time.Time

	// The name of the live source.
	LiveSourceName *string

	// The name of the source location associated with the VOD source.
	SourceLocationName *string

	// The tags assigned to the live source.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLiveSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeLiveSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeLiveSource{}, middleware.After)
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
	if err = addOpDescribeLiveSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLiveSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLiveSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "DescribeLiveSource",
	}
}
