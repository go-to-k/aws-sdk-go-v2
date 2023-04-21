// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update the Proton Ops config file.
func (c *Client) UpdateServiceSyncConfig(ctx context.Context, params *UpdateServiceSyncConfigInput, optFns ...func(*Options)) (*UpdateServiceSyncConfigOutput, error) {
	if params == nil {
		params = &UpdateServiceSyncConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateServiceSyncConfig", params, optFns, c.addOperationUpdateServiceSyncConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServiceSyncConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServiceSyncConfigInput struct {

	// The name of the code repository branch where the Proton Ops file is found.
	//
	// This member is required.
	Branch *string

	// The path to the Proton Ops file.
	//
	// This member is required.
	FilePath *string

	// The name of the repository where the Proton Ops file is found.
	//
	// This member is required.
	RepositoryName *string

	// The name of the repository provider where the Proton Ops file is found.
	//
	// This member is required.
	RepositoryProvider types.RepositoryProvider

	// The name of the service the Proton Ops file is for.
	//
	// This member is required.
	ServiceName *string

	noSmithyDocumentSerde
}

type UpdateServiceSyncConfigOutput struct {

	// The detailed data of the Proton Ops file.
	ServiceSyncConfig *types.ServiceSyncConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateServiceSyncConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateServiceSyncConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateServiceSyncConfig{}, middleware.After)
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
	if err = addOpUpdateServiceSyncConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServiceSyncConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateServiceSyncConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "UpdateServiceSyncConfig",
	}
}
