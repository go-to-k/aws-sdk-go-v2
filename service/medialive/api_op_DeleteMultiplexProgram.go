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

// Delete a program from a multiplex.
func (c *Client) DeleteMultiplexProgram(ctx context.Context, params *DeleteMultiplexProgramInput, optFns ...func(*Options)) (*DeleteMultiplexProgramOutput, error) {
	if params == nil {
		params = &DeleteMultiplexProgramInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMultiplexProgram", params, optFns, c.addOperationDeleteMultiplexProgramMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMultiplexProgramOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DeleteMultiplexProgramRequest
type DeleteMultiplexProgramInput struct {

	// The ID of the multiplex that the program belongs to.
	//
	// This member is required.
	MultiplexId *string

	// The multiplex program name.
	//
	// This member is required.
	ProgramName *string

	noSmithyDocumentSerde
}

// Placeholder documentation for DeleteMultiplexProgramResponse
type DeleteMultiplexProgramOutput struct {

	// The MediaLive channel associated with the program.
	ChannelId *string

	// The settings for this multiplex program.
	MultiplexProgramSettings *types.MultiplexProgramSettings

	// The packet identifier map for this multiplex program.
	PacketIdentifiersMap *types.MultiplexProgramPacketIdentifiersMap

	// Contains information about the current sources for the specified program in the
	// specified multiplex. Keep in mind that each multiplex pipeline connects to both
	// pipelines in a given source channel (the channel identified by the program). But
	// only one of those channel pipelines is ever active at one time.
	PipelineDetails []types.MultiplexProgramPipelineDetail

	// The name of the multiplex program.
	ProgramName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteMultiplexProgramMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteMultiplexProgram{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteMultiplexProgram{}, middleware.After)
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
	if err = addOpDeleteMultiplexProgramValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMultiplexProgram(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteMultiplexProgram(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DeleteMultiplexProgram",
	}
}
