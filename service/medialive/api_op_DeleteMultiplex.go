// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete a multiplex. The multiplex must be idle.
func (c *Client) DeleteMultiplex(ctx context.Context, params *DeleteMultiplexInput, optFns ...func(*Options)) (*DeleteMultiplexOutput, error) {
	if params == nil {
		params = &DeleteMultiplexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMultiplex", params, optFns, c.addOperationDeleteMultiplexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMultiplexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DeleteMultiplexRequest
type DeleteMultiplexInput struct {

	// The ID of the multiplex.
	//
	// This member is required.
	MultiplexId *string

	noSmithyDocumentSerde
}

// Placeholder documentation for DeleteMultiplexResponse
type DeleteMultiplexOutput struct {

	// The unique arn of the multiplex.
	Arn *string

	// A list of availability zones for the multiplex.
	AvailabilityZones []string

	// A list of the multiplex output destinations.
	Destinations []types.MultiplexOutputDestination

	// The unique id of the multiplex.
	Id *string

	// Configuration for a multiplex event.
	MultiplexSettings *types.MultiplexSettings

	// The name of the multiplex.
	Name *string

	// The number of currently healthy pipelines.
	PipelinesRunningCount int32

	// The number of programs in the multiplex.
	ProgramCount int32

	// The current state of the multiplex.
	State types.MultiplexState

	// A collection of key-value pairs.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteMultiplexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteMultiplex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteMultiplex{}, middleware.After)
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
	if err = addOpDeleteMultiplexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMultiplex(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteMultiplex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DeleteMultiplex",
	}
}
