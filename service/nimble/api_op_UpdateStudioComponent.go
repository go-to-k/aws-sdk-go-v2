// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a studio component resource.
func (c *Client) UpdateStudioComponent(ctx context.Context, params *UpdateStudioComponentInput, optFns ...func(*Options)) (*UpdateStudioComponentOutput, error) {
	if params == nil {
		params = &UpdateStudioComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStudioComponent", params, optFns, c.addOperationUpdateStudioComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStudioComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStudioComponentInput struct {

	// The studio component ID.
	//
	// This member is required.
	StudioComponentId *string

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If you don’t specify a client token, the Amazon Web Services SDK
	// automatically generates a client token and uses it for the request to ensure
	// idempotency.
	ClientToken *string

	// The configuration of the studio component, based on component type.
	Configuration *types.StudioComponentConfiguration

	// The description.
	Description *string

	// The EC2 security groups that control access to the studio component.
	Ec2SecurityGroupIds []string

	// Initialization scripts for studio components.
	InitializationScripts []types.StudioComponentInitializationScript

	// The name for the studio component.
	Name *string

	// An IAM role attached to a Studio Component that gives the studio component
	// access to Amazon Web Services resources at anytime while the instance is
	// running.
	RuntimeRoleArn *string

	// Parameters for the studio component scripts.
	ScriptParameters []types.ScriptParameterKeyValue

	// An IAM role attached to Studio Component when the system initialization script
	// runs which give the studio component access to Amazon Web Services resources
	// when the system initialization script runs.
	SecureInitializationRoleArn *string

	// The specific subtype of a studio component.
	Subtype types.StudioComponentSubtype

	// The type of the studio component.
	Type types.StudioComponentType

	noSmithyDocumentSerde
}

type UpdateStudioComponentOutput struct {

	// Information about the studio component.
	StudioComponent *types.StudioComponent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateStudioComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateStudioComponent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateStudioComponent{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateStudioComponentMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateStudioComponentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStudioComponent(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateStudioComponent struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateStudioComponent) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateStudioComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateStudioComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateStudioComponentInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateStudioComponentMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateStudioComponent{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateStudioComponent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "UpdateStudioComponent",
	}
}
