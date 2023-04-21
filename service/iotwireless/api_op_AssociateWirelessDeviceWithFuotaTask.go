// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associate a wireless device with a FUOTA task.
func (c *Client) AssociateWirelessDeviceWithFuotaTask(ctx context.Context, params *AssociateWirelessDeviceWithFuotaTaskInput, optFns ...func(*Options)) (*AssociateWirelessDeviceWithFuotaTaskOutput, error) {
	if params == nil {
		params = &AssociateWirelessDeviceWithFuotaTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateWirelessDeviceWithFuotaTask", params, optFns, c.addOperationAssociateWirelessDeviceWithFuotaTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateWirelessDeviceWithFuotaTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateWirelessDeviceWithFuotaTaskInput struct {

	// The ID of a FUOTA task.
	//
	// This member is required.
	Id *string

	// The ID of the wireless device.
	//
	// This member is required.
	WirelessDeviceId *string

	noSmithyDocumentSerde
}

type AssociateWirelessDeviceWithFuotaTaskOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateWirelessDeviceWithFuotaTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateWirelessDeviceWithFuotaTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateWirelessDeviceWithFuotaTask{}, middleware.After)
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
	if err = addOpAssociateWirelessDeviceWithFuotaTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateWirelessDeviceWithFuotaTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateWirelessDeviceWithFuotaTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "AssociateWirelessDeviceWithFuotaTask",
	}
}
