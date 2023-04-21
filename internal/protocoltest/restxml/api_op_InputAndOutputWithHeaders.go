// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// The example tests how requests and responses are serialized when there is no
// input or output payload but there are HTTP header bindings.
func (c *Client) InputAndOutputWithHeaders(ctx context.Context, params *InputAndOutputWithHeadersInput, optFns ...func(*Options)) (*InputAndOutputWithHeadersOutput, error) {
	if params == nil {
		params = &InputAndOutputWithHeadersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InputAndOutputWithHeaders", params, optFns, c.addOperationInputAndOutputWithHeadersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InputAndOutputWithHeadersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InputAndOutputWithHeadersInput struct {
	HeaderBooleanList []bool

	HeaderByte *int8

	HeaderDouble *float64

	HeaderEnum types.FooEnum

	HeaderEnumList []types.FooEnum

	HeaderFalseBool *bool

	HeaderFloat *float32

	HeaderInteger *int32

	HeaderIntegerList []int32

	HeaderLong *int64

	HeaderShort *int16

	HeaderString *string

	HeaderStringList []string

	HeaderStringSet []string

	HeaderTimestampList []time.Time

	HeaderTrueBool *bool

	noSmithyDocumentSerde
}

type InputAndOutputWithHeadersOutput struct {
	HeaderBooleanList []bool

	HeaderByte *int8

	HeaderDouble *float64

	HeaderEnum types.FooEnum

	HeaderEnumList []types.FooEnum

	HeaderFalseBool *bool

	HeaderFloat *float32

	HeaderInteger *int32

	HeaderIntegerList []int32

	HeaderLong *int64

	HeaderShort *int16

	HeaderString *string

	HeaderStringList []string

	HeaderStringSet []string

	HeaderTimestampList []time.Time

	HeaderTrueBool *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInputAndOutputWithHeadersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpInputAndOutputWithHeaders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpInputAndOutputWithHeaders{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInputAndOutputWithHeaders(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInputAndOutputWithHeaders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InputAndOutputWithHeaders",
	}
}
