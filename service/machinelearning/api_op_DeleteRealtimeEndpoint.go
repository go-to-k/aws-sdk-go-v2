// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a real time endpoint of an MLModel .
func (c *Client) DeleteRealtimeEndpoint(ctx context.Context, params *DeleteRealtimeEndpointInput, optFns ...func(*Options)) (*DeleteRealtimeEndpointOutput, error) {
	if params == nil {
		params = &DeleteRealtimeEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRealtimeEndpoint", params, optFns, c.addOperationDeleteRealtimeEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRealtimeEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRealtimeEndpointInput struct {

	// The ID assigned to the MLModel during creation.
	//
	// This member is required.
	MLModelId *string

	noSmithyDocumentSerde
}

// Represents the output of an DeleteRealtimeEndpoint operation. The result
// contains the MLModelId and the endpoint information for the MLModel .
type DeleteRealtimeEndpointOutput struct {

	// A user-supplied ID that uniquely identifies the MLModel . This value should be
	// identical to the value of the MLModelId in the request.
	MLModelId *string

	// The endpoint information of the MLModel
	RealtimeEndpointInfo *types.RealtimeEndpointInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteRealtimeEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteRealtimeEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteRealtimeEndpoint{}, middleware.After)
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
	if err = addOpDeleteRealtimeEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRealtimeEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRealtimeEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "DeleteRealtimeEndpoint",
	}
}
