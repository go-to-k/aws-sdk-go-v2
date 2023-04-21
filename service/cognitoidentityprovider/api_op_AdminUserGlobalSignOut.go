// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Signs out a user from all devices. You must sign AdminUserGlobalSignOut
// requests with Amazon Web Services credentials. It also invalidates all refresh
// tokens that Amazon Cognito has issued to a user. The user's current access and
// ID tokens remain valid until they expire. By default, access and ID tokens
// expire one hour after they're issued. A user can still use a hosted UI cookie to
// retrieve new tokens for the duration of the cookie validity period of 1 hour.
// Calling this action requires developer credentials.
func (c *Client) AdminUserGlobalSignOut(ctx context.Context, params *AdminUserGlobalSignOutInput, optFns ...func(*Options)) (*AdminUserGlobalSignOutOutput, error) {
	if params == nil {
		params = &AdminUserGlobalSignOutInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminUserGlobalSignOut", params, optFns, c.addOperationAdminUserGlobalSignOutMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminUserGlobalSignOutOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to sign out of all devices, as an administrator.
type AdminUserGlobalSignOutInput struct {

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The user name.
	//
	// This member is required.
	Username *string

	noSmithyDocumentSerde
}

// The global sign-out response, as an administrator.
type AdminUserGlobalSignOutOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminUserGlobalSignOutMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminUserGlobalSignOut{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminUserGlobalSignOut{}, middleware.After)
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
	if err = addOpAdminUserGlobalSignOutValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminUserGlobalSignOut(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminUserGlobalSignOut(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminUserGlobalSignOut",
	}
}
