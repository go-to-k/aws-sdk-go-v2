// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists your delivery streams in alphabetical order of their names. The number of
// delivery streams might be too large to return using a single call to
// ListDeliveryStreams . You can limit the number of delivery streams returned,
// using the Limit parameter. To determine whether there are more delivery streams
// to list, check the value of HasMoreDeliveryStreams in the output. If there are
// more delivery streams to list, you can request them by calling this operation
// again and setting the ExclusiveStartDeliveryStreamName parameter to the name of
// the last delivery stream returned in the last call.
func (c *Client) ListDeliveryStreams(ctx context.Context, params *ListDeliveryStreamsInput, optFns ...func(*Options)) (*ListDeliveryStreamsOutput, error) {
	if params == nil {
		params = &ListDeliveryStreamsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeliveryStreams", params, optFns, c.addOperationListDeliveryStreamsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeliveryStreamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeliveryStreamsInput struct {

	// The delivery stream type. This can be one of the following values:
	//   - DirectPut : Provider applications access the delivery stream directly.
	//   - KinesisStreamAsSource : The delivery stream uses a Kinesis data stream as a
	//   source.
	// This parameter is optional. If this parameter is omitted, delivery streams of
	// all types are returned.
	DeliveryStreamType types.DeliveryStreamType

	// The list of delivery streams returned by this call to ListDeliveryStreams will
	// start with the delivery stream whose name comes alphabetically immediately after
	// the name you specify in ExclusiveStartDeliveryStreamName .
	ExclusiveStartDeliveryStreamName *string

	// The maximum number of delivery streams to list. The default value is 10.
	Limit *int32

	noSmithyDocumentSerde
}

type ListDeliveryStreamsOutput struct {

	// The names of the delivery streams.
	//
	// This member is required.
	DeliveryStreamNames []string

	// Indicates whether there are more delivery streams available to list.
	//
	// This member is required.
	HasMoreDeliveryStreams *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDeliveryStreamsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDeliveryStreams{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDeliveryStreams{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeliveryStreams(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDeliveryStreams(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "firehose",
		OperationName: "ListDeliveryStreams",
	}
}
