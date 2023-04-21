// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// The example tests how requests are serialized when there's no input payload but
// there are HTTP labels.
func (c *Client) HttpRequestWithLabels(ctx context.Context, params *HttpRequestWithLabelsInput, optFns ...func(*Options)) (*HttpRequestWithLabelsOutput, error) {
	if params == nil {
		params = &HttpRequestWithLabelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HttpRequestWithLabels", params, optFns, c.addOperationHttpRequestWithLabelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HttpRequestWithLabelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpRequestWithLabelsInput struct {

	// Serialized in the path as true or false.
	//
	// This member is required.
	Boolean *bool

	// This member is required.
	Double *float64

	// This member is required.
	Float *float32

	// This member is required.
	Integer *int32

	// This member is required.
	Long *int64

	// This member is required.
	Short *int16

	// This member is required.
	String_ *string

	// Note that this member has no format, so it's serialized as an RFC 3399
	// date-time.
	//
	// This member is required.
	Timestamp *time.Time

	noSmithyDocumentSerde
}

type HttpRequestWithLabelsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationHttpRequestWithLabelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpHttpRequestWithLabels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpHttpRequestWithLabels{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addOpHttpRequestWithLabelsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opHttpRequestWithLabels(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opHttpRequestWithLabels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpRequestWithLabels",
	}
}
