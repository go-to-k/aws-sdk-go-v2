// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows the developer to delete the user pool client.
func (c *Client) DeleteUserPoolClient(ctx context.Context, params *DeleteUserPoolClientInput, optFns ...func(*Options)) (*DeleteUserPoolClientOutput, error) {
	if params == nil {
		params = &DeleteUserPoolClientInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteUserPoolClient", params, optFns, c.addOperationDeleteUserPoolClientMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteUserPoolClientOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to delete a user pool client.
type DeleteUserPoolClientInput struct {

	// The app client ID of the app associated with the user pool.
	//
	// This member is required.
	ClientId *string

	// The user pool ID for the user pool where you want to delete the client.
	//
	// This member is required.
	UserPoolId *string

	noSmithyDocumentSerde
}

type DeleteUserPoolClientOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteUserPoolClientMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteUserPoolClient{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteUserPoolClient{}, middleware.After)
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
	if err = addOpDeleteUserPoolClientValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUserPoolClient(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteUserPoolClient(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "DeleteUserPoolClient",
	}
}
