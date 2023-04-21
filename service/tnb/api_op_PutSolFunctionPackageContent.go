// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/tnb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Uploads the contents of a function package. A function package is a .zip file
// in CSAR (Cloud Service Archive) format that contains a network function (an ETSI
// standard telecommunication application) and function package descriptor that
// uses the TOSCA standard to describe how the network functions should run on your
// network.
func (c *Client) PutSolFunctionPackageContent(ctx context.Context, params *PutSolFunctionPackageContentInput, optFns ...func(*Options)) (*PutSolFunctionPackageContentOutput, error) {
	if params == nil {
		params = &PutSolFunctionPackageContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutSolFunctionPackageContent", params, optFns, c.addOperationPutSolFunctionPackageContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutSolFunctionPackageContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSolFunctionPackageContentInput struct {

	// Function package file.
	//
	// This member is required.
	File []byte

	// Function package ID.
	//
	// This member is required.
	VnfPkgId *string

	// Function package content type.
	ContentType types.PackageContentType

	noSmithyDocumentSerde
}

type PutSolFunctionPackageContentOutput struct {

	// Function package ID.
	//
	// This member is required.
	Id *string

	// Function package metadata.
	//
	// This member is required.
	Metadata *types.PutSolFunctionPackageContentMetadata

	// Function product name.
	//
	// This member is required.
	VnfProductName *string

	// Function provider.
	//
	// This member is required.
	VnfProvider *string

	// Function package descriptor ID.
	//
	// This member is required.
	VnfdId *string

	// Function package descriptor version.
	//
	// This member is required.
	VnfdVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutSolFunctionPackageContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutSolFunctionPackageContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutSolFunctionPackageContent{}, middleware.After)
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
	if err = addOpPutSolFunctionPackageContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutSolFunctionPackageContent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutSolFunctionPackageContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tnb",
		OperationName: "PutSolFunctionPackageContent",
	}
}
