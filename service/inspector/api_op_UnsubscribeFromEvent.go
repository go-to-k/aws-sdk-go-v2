// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables the process of sending Amazon Simple Notification Service (SNS)
// notifications about a specified event to a specified SNS topic.
func (c *Client) UnsubscribeFromEvent(ctx context.Context, params *UnsubscribeFromEventInput, optFns ...func(*Options)) (*UnsubscribeFromEventOutput, error) {
	if params == nil {
		params = &UnsubscribeFromEventInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UnsubscribeFromEvent", params, optFns, c.addOperationUnsubscribeFromEventMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UnsubscribeFromEventOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UnsubscribeFromEventInput struct {

	// The event for which you want to stop receiving SNS notifications.
	//
	// This member is required.
	Event types.InspectorEvent

	// The ARN of the assessment template that is used during the event for which you
	// want to stop receiving SNS notifications.
	//
	// This member is required.
	ResourceArn *string

	// The ARN of the SNS topic to which SNS notifications are sent.
	//
	// This member is required.
	TopicArn *string

	noSmithyDocumentSerde
}

type UnsubscribeFromEventOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUnsubscribeFromEventMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUnsubscribeFromEvent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUnsubscribeFromEvent{}, middleware.After)
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
	if err = addOpUnsubscribeFromEventValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUnsubscribeFromEvent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUnsubscribeFromEvent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "UnsubscribeFromEvent",
	}
}
