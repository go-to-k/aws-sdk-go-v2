// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// To deregister a consumer, provide its ARN. Alternatively, you can provide the
// ARN of the data stream and the name you gave the consumer when you registered
// it. You may also provide all three parameters, as long as they don't conflict
// with each other. If you don't know the name or ARN of the consumer that you want
// to deregister, you can use the ListStreamConsumers operation to get a list of
// the descriptions of all the consumers that are currently registered with a given
// data stream. The description of a consumer contains its name and ARN. This
// operation has a limit of five transactions per second per stream.
func (c *Client) DeregisterStreamConsumer(ctx context.Context, params *DeregisterStreamConsumerInput, optFns ...func(*Options)) (*DeregisterStreamConsumerOutput, error) {
	if params == nil {
		params = &DeregisterStreamConsumerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterStreamConsumer", params, optFns, c.addOperationDeregisterStreamConsumerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterStreamConsumerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterStreamConsumerInput struct {

	// The ARN returned by Kinesis Data Streams when you registered the consumer. If
	// you don't know the ARN of the consumer that you want to deregister, you can use
	// the ListStreamConsumers operation to get a list of the descriptions of all the
	// consumers that are currently registered with a given data stream. The
	// description of a consumer contains its ARN.
	ConsumerARN *string

	// The name that you gave to the consumer.
	ConsumerName *string

	// The ARN of the Kinesis data stream that the consumer is registered with. For
	// more information, see Amazon Resource Names (ARNs) and Amazon Web Services
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kinesis-streams)
	// .
	StreamARN *string

	noSmithyDocumentSerde
}

type DeregisterStreamConsumerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterStreamConsumerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterStreamConsumer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterStreamConsumer{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterStreamConsumer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterStreamConsumer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "DeregisterStreamConsumer",
	}
}
