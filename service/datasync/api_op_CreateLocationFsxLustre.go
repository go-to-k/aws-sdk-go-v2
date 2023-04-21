// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an endpoint for an Amazon FSx for Lustre file system.
func (c *Client) CreateLocationFsxLustre(ctx context.Context, params *CreateLocationFsxLustreInput, optFns ...func(*Options)) (*CreateLocationFsxLustreOutput, error) {
	if params == nil {
		params = &CreateLocationFsxLustreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationFsxLustre", params, optFns, c.addOperationCreateLocationFsxLustreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationFsxLustreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLocationFsxLustreInput struct {

	// The Amazon Resource Name (ARN) for the FSx for Lustre file system.
	//
	// This member is required.
	FsxFilesystemArn *string

	// The Amazon Resource Names (ARNs) of the security groups that are used to
	// configure the FSx for Lustre file system.
	//
	// This member is required.
	SecurityGroupArns []string

	// A subdirectory in the location's path. This subdirectory in the FSx for Lustre
	// file system is used to read data from the FSx for Lustre source location or
	// write data to the FSx for Lustre destination.
	Subdirectory *string

	// The key-value pair that represents a tag that you want to add to the resource.
	// The value can be an empty string. This value helps you manage, filter, and
	// search for your resources. We recommend that you create a name tag for your
	// location.
	Tags []types.TagListEntry

	noSmithyDocumentSerde
}

type CreateLocationFsxLustreOutput struct {

	// The Amazon Resource Name (ARN) of the FSx for Lustre file system location
	// that's created.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLocationFsxLustreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationFsxLustre{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationFsxLustre{}, middleware.After)
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
	if err = addOpCreateLocationFsxLustreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationFsxLustre(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocationFsxLustre(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CreateLocationFsxLustre",
	}
}
