// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a virtual router within a service mesh. Specify a listener for any
// inbound traffic that your virtual router receives. Create a virtual router for
// each protocol and port that you need to route. Virtual routers handle traffic
// for one or more virtual services within your mesh. After you create your virtual
// router, create and associate routes for your virtual router that direct incoming
// requests to different virtual nodes. For more information about virtual routers,
// see Virtual routers (https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_routers.html)
// .
func (c *Client) CreateVirtualRouter(ctx context.Context, params *CreateVirtualRouterInput, optFns ...func(*Options)) (*CreateVirtualRouterOutput, error) {
	if params == nil {
		params = &CreateVirtualRouterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVirtualRouter", params, optFns, c.addOperationCreateVirtualRouterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVirtualRouterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVirtualRouterInput struct {

	// The name of the service mesh to create the virtual router in.
	//
	// This member is required.
	MeshName *string

	// The virtual router specification to apply.
	//
	// This member is required.
	Spec *types.VirtualRouterSpec

	// The name to use for the virtual router.
	//
	// This member is required.
	VirtualRouterName *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. Up to 36 letters, numbers, hyphens, and underscores are allowed.
	ClientToken *string

	// The Amazon Web Services IAM account ID of the service mesh owner. If the
	// account ID is not your own, then the account that you specify must share the
	// mesh with your account before you can create the resource in the service mesh.
	// For more information about mesh sharing, see Working with shared meshes (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html)
	// .
	MeshOwner *string

	// Optional metadata that you can apply to the virtual router to assist with
	// categorization and organization. Each tag consists of a key and an optional
	// value, both of which you define. Tag keys can have a maximum character length of
	// 128 characters, and tag values can have a maximum length of 256 characters.
	Tags []types.TagRef

	noSmithyDocumentSerde
}

type CreateVirtualRouterOutput struct {

	// The full description of your virtual router following the create call.
	//
	// This member is required.
	VirtualRouter *types.VirtualRouterData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVirtualRouterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVirtualRouter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVirtualRouter{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateVirtualRouterMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateVirtualRouterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVirtualRouter(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateVirtualRouter struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateVirtualRouter) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateVirtualRouter) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateVirtualRouterInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateVirtualRouterInput ")
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
func addIdempotencyToken_opCreateVirtualRouterMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateVirtualRouter{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateVirtualRouter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appmesh",
		OperationName: "CreateVirtualRouter",
	}
}
