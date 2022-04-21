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

// Updates a source location on a specific channel.
func (c *Client) UpdateSourceLocation(ctx context.Context, params *UpdateSourceLocationInput, optFns ...func(*Options)) (*UpdateSourceLocationOutput, error) {
	if params == nil {
		params = &UpdateSourceLocationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSourceLocation", params, optFns, c.addOperationUpdateSourceLocationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSourceLocationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSourceLocationInput struct {

	// The HTTP configuration for the source location.
	//
	// This member is required.
	HttpConfiguration *types.HttpConfiguration

	// The identifier for the source location you are working on.
	//
	// This member is required.
	SourceLocationName *string

	// Access configuration parameters. Configures the type of authentication used to
	// access content from your source location.
	AccessConfiguration *types.AccessConfiguration

	// The optional configuration for the host server that serves segments.
	DefaultSegmentDeliveryConfiguration *types.DefaultSegmentDeliveryConfiguration

	// A list of the segment delivery configurations associated with this resource.
	SegmentDeliveryConfigurations []types.SegmentDeliveryConfiguration

	noSmithyDocumentSerde
}

type UpdateSourceLocationOutput struct {

	// The access configuration for the source location.
	AccessConfiguration *types.AccessConfiguration

	// The ARN of the source location.
	Arn *string

	// The timestamp that indicates when the source location was created.
	CreationTime *time.Time

	// The default segment delivery configuration settings.
	DefaultSegmentDeliveryConfiguration *types.DefaultSegmentDeliveryConfiguration

	// The HTTP package configuration settings for the source location.
	HttpConfiguration *types.HttpConfiguration

	// The timestamp that indicates when the source location was last modified.
	LastModifiedTime *time.Time

	// A list of the segment delivery configurations associated with this resource.
	SegmentDeliveryConfigurations []types.SegmentDeliveryConfiguration

	// The name of the source location.
	SourceLocationName *string

	// The tags assigned to the source location.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSourceLocationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSourceLocation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSourceLocation{}, middleware.After)
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
	if err = addOpUpdateSourceLocationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSourceLocation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSourceLocation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "UpdateSourceLocation",
	}
}
