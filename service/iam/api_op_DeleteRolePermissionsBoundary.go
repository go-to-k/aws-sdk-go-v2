// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the permissions boundary for the specified IAM role. You cannot set the
// boundary for a service-linked role. Deleting the permissions boundary for a role
// might increase its permissions. For example, it might allow anyone who assumes
// the role to perform all the actions granted in its permissions policies.
func (c *Client) DeleteRolePermissionsBoundary(ctx context.Context, params *DeleteRolePermissionsBoundaryInput, optFns ...func(*Options)) (*DeleteRolePermissionsBoundaryOutput, error) {
	if params == nil {
		params = &DeleteRolePermissionsBoundaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRolePermissionsBoundary", params, optFns, c.addOperationDeleteRolePermissionsBoundaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRolePermissionsBoundaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRolePermissionsBoundaryInput struct {

	// The name (friendly name, not ARN) of the IAM role from which you want to remove
	// the permissions boundary.
	//
	// This member is required.
	RoleName *string

	noSmithyDocumentSerde
}

type DeleteRolePermissionsBoundaryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteRolePermissionsBoundaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteRolePermissionsBoundary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteRolePermissionsBoundary{}, middleware.After)
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
	if err = addOpDeleteRolePermissionsBoundaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRolePermissionsBoundary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRolePermissionsBoundary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DeleteRolePermissionsBoundary",
	}
}
