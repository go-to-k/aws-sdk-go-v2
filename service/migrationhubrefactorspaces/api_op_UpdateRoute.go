// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubrefactorspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an Amazon Web Services Migration Hub Refactor Spaces route.
func (c *Client) UpdateRoute(ctx context.Context, params *UpdateRouteInput, optFns ...func(*Options)) (*UpdateRouteOutput, error) {
	if params == nil {
		params = &UpdateRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRoute", params, optFns, c.addOperationUpdateRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRouteInput struct {

	// If set to ACTIVE , traffic is forwarded to this route’s service after the route
	// is updated.
	//
	// This member is required.
	ActivationState types.RouteActivationState

	// The ID of the application within which the route is being updated.
	//
	// This member is required.
	ApplicationIdentifier *string

	// The ID of the environment in which the route is being updated.
	//
	// This member is required.
	EnvironmentIdentifier *string

	// The unique identifier of the route to update.
	//
	// This member is required.
	RouteIdentifier *string

	noSmithyDocumentSerde
}

type UpdateRouteOutput struct {

	// The ID of the application in which the route is being updated.
	ApplicationId *string

	// The Amazon Resource Name (ARN) of the route. The format for this ARN is
	// arn:aws:refactor-spaces:region:account-id:resource-type/resource-id . For more
	// information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	Arn *string

	// A timestamp that indicates when the route was last updated.
	LastUpdatedTime *time.Time

	// The unique identifier of the route.
	RouteId *string

	// The ID of service in which the route was created. Traffic that matches this
	// route is forwarded to this service.
	ServiceId *string

	// The current state of the route.
	State types.RouteState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRoute{}, middleware.After)
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
	if err = addOpUpdateRouteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRoute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRoute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "refactor-spaces",
		OperationName: "UpdateRoute",
	}
}
