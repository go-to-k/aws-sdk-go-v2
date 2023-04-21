// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Put position configuration for a given resource. This action is no longer
// supported. Calls to update the position configuration should use the
// UpdateResourcePosition (https://docs.aws.amazon.com/iot-wireless/2020-11-22/apireference/API_UpdateResourcePosition.html)
// API operation instead.
//
// Deprecated: This operation is no longer supported.
func (c *Client) PutPositionConfiguration(ctx context.Context, params *PutPositionConfigurationInput, optFns ...func(*Options)) (*PutPositionConfigurationOutput, error) {
	if params == nil {
		params = &PutPositionConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPositionConfiguration", params, optFns, c.addOperationPutPositionConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPositionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPositionConfigurationInput struct {

	// Resource identifier used to update the position configuration.
	//
	// This member is required.
	ResourceIdentifier *string

	// Resource type of the resource for which you want to update the position
	// configuration.
	//
	// This member is required.
	ResourceType types.PositionResourceType

	// The position data destination that describes the AWS IoT rule that processes
	// the device's position data for use by AWS IoT Core for LoRaWAN.
	Destination *string

	// The positioning solvers used to update the position configuration of the
	// resource.
	Solvers *types.PositionSolverConfigurations

	noSmithyDocumentSerde
}

type PutPositionConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutPositionConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutPositionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutPositionConfiguration{}, middleware.After)
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
	if err = addOpPutPositionConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutPositionConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutPositionConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "PutPositionConfiguration",
	}
}
