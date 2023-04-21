// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update a component. There are a few modes for updating a component. The
// deploymentType field defines the mode. You can't update a component while its
// deployment status, or the deployment status of a service instance attached to
// it, is IN_PROGRESS . For more information about components, see Proton
// components (https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html)
// in the Proton User Guide.
func (c *Client) UpdateComponent(ctx context.Context, params *UpdateComponentInput, optFns ...func(*Options)) (*UpdateComponentOutput, error) {
	if params == nil {
		params = &UpdateComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateComponent", params, optFns, c.addOperationUpdateComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateComponentInput struct {

	// The deployment type. It defines the mode for updating a component, as follows:
	// NONE In this mode, a deployment doesn't occur. Only the requested metadata
	// parameters are updated. You can only specify description in this mode.
	// CURRENT_VERSION In this mode, the component is deployed and updated with the new
	// serviceSpec , templateSource , and/or type that you provide. Only requested
	// parameters are updated.
	//
	// This member is required.
	DeploymentType types.ComponentDeploymentUpdateType

	// The name of the component to update.
	//
	// This member is required.
	Name *string

	// The client token for the updated component.
	ClientToken *string

	// An optional customer-provided description of the component.
	Description *string

	// The name of the service instance that you want to attach this component to.
	// Don't specify to keep the component's current service instance attachment.
	// Specify an empty string to detach the component from the service instance it's
	// attached to. Specify non-empty values for both serviceInstanceName and
	// serviceName or for neither of them.
	ServiceInstanceName *string

	// The name of the service that serviceInstanceName is associated with. Don't
	// specify to keep the component's current service instance attachment. Specify an
	// empty string to detach the component from the service instance it's attached to.
	// Specify non-empty values for both serviceInstanceName and serviceName or for
	// neither of them.
	ServiceName *string

	// The service spec that you want the component to use to access service inputs.
	// Set this only when the component is attached to a service instance.
	//
	// This value conforms to the media type: application/yaml
	ServiceSpec *string

	// A path to the Infrastructure as Code (IaC) file describing infrastructure that
	// a custom component provisions. Components support a single IaC file, even if you
	// use Terraform as your template language.
	//
	// This value conforms to the media type: application/yaml
	TemplateFile *string

	noSmithyDocumentSerde
}

type UpdateComponentOutput struct {

	// The detailed data of the updated component.
	//
	// This member is required.
	Component *types.Component

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateComponent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateComponent{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateComponentMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateComponentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateComponent(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateComponent struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateComponent) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateComponent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateComponentInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateComponentInput ")
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
func addIdempotencyToken_opUpdateComponentMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateComponent{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateComponent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "UpdateComponent",
	}
}
