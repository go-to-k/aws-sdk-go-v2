// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified listener for the specified service.
func (c *Client) UpdateListener(ctx context.Context, params *UpdateListenerInput, optFns ...func(*Options)) (*UpdateListenerOutput, error) {
	if params == nil {
		params = &UpdateListenerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateListener", params, optFns, c.addOperationUpdateListenerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateListenerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateListenerInput struct {

	// The action for the default rule.
	//
	// This member is required.
	DefaultAction types.RuleAction

	// The ID or Amazon Resource Name (ARN) of the listener.
	//
	// This member is required.
	ListenerIdentifier *string

	// The ID or Amazon Resource Name (ARN) of the service.
	//
	// This member is required.
	ServiceIdentifier *string

	noSmithyDocumentSerde
}

type UpdateListenerOutput struct {

	// The Amazon Resource Name (ARN) of the listener.
	Arn *string

	// The action for the default rule.
	DefaultAction types.RuleAction

	// The ID of the listener.
	Id *string

	// The name of the listener.
	Name *string

	// The listener port.
	Port *int32

	// The protocol of the listener.
	Protocol types.ListenerProtocol

	// The Amazon Resource Name (ARN) of the service.
	ServiceArn *string

	// The ID of the service.
	ServiceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateListenerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateListener{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateListener{}, middleware.After)
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
	if err = addOpUpdateListenerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateListener(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateListener(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "vpc-lattice",
		OperationName: "UpdateListener",
	}
}
