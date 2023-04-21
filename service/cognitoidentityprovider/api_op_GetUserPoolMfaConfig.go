// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the user pool multi-factor authentication (MFA) configuration.
func (c *Client) GetUserPoolMfaConfig(ctx context.Context, params *GetUserPoolMfaConfigInput, optFns ...func(*Options)) (*GetUserPoolMfaConfigOutput, error) {
	if params == nil {
		params = &GetUserPoolMfaConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUserPoolMfaConfig", params, optFns, c.addOperationGetUserPoolMfaConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUserPoolMfaConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUserPoolMfaConfigInput struct {

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	noSmithyDocumentSerde
}

type GetUserPoolMfaConfigOutput struct {

	// The multi-factor authentication (MFA) configuration. Valid values include:
	//   - OFF MFA won't be used for any users.
	//   - ON MFA is required for all users to sign in.
	//   - OPTIONAL MFA will be required only for individual users who have an MFA
	//   factor activated.
	MfaConfiguration types.UserPoolMfaType

	// The SMS text message multi-factor authentication (MFA) configuration.
	SmsMfaConfiguration *types.SmsMfaConfigType

	// The software token multi-factor authentication (MFA) configuration.
	SoftwareTokenMfaConfiguration *types.SoftwareTokenMfaConfigType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUserPoolMfaConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetUserPoolMfaConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUserPoolMfaConfig{}, middleware.After)
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
	if err = addOpGetUserPoolMfaConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUserPoolMfaConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUserPoolMfaConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "GetUserPoolMfaConfig",
	}
}
