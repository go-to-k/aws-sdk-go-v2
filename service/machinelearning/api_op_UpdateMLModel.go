// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the MLModelName and the ScoreThreshold of an MLModel . You can use the
// GetMLModel operation to view the contents of the updated data element.
func (c *Client) UpdateMLModel(ctx context.Context, params *UpdateMLModelInput, optFns ...func(*Options)) (*UpdateMLModelOutput, error) {
	if params == nil {
		params = &UpdateMLModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMLModel", params, optFns, c.addOperationUpdateMLModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMLModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMLModelInput struct {

	// The ID assigned to the MLModel during creation.
	//
	// This member is required.
	MLModelId *string

	// A user-supplied name or description of the MLModel .
	MLModelName *string

	// The ScoreThreshold used in binary classification MLModel that marks the
	// boundary between a positive prediction and a negative prediction. Output values
	// greater than or equal to the ScoreThreshold receive a positive result from the
	// MLModel , such as true . Output values less than the ScoreThreshold receive a
	// negative response from the MLModel , such as false .
	ScoreThreshold *float32

	noSmithyDocumentSerde
}

// Represents the output of an UpdateMLModel operation. You can see the updated
// content by using the GetMLModel operation.
type UpdateMLModelOutput struct {

	// The ID assigned to the MLModel during creation. This value should be identical
	// to the value of the MLModelID in the request.
	MLModelId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMLModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMLModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMLModel{}, middleware.After)
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
	if err = addOpUpdateMLModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMLModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMLModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "UpdateMLModel",
	}
}
