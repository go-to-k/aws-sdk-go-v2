// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associate resources to a Firewall Manager resource set.
func (c *Client) BatchAssociateResource(ctx context.Context, params *BatchAssociateResourceInput, optFns ...func(*Options)) (*BatchAssociateResourceOutput, error) {
	if params == nil {
		params = &BatchAssociateResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchAssociateResource", params, optFns, c.addOperationBatchAssociateResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchAssociateResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchAssociateResourceInput struct {

	// The uniform resource identifiers (URIs) of resources that should be associated
	// to the resource set. The URIs must be Amazon Resource Names (ARNs).
	//
	// This member is required.
	Items []string

	// A unique identifier for the resource set, used in a TODO to refer to the
	// resource set.
	//
	// This member is required.
	ResourceSetIdentifier *string

	noSmithyDocumentSerde
}

type BatchAssociateResourceOutput struct {

	// The resources that failed to associate to the resource set.
	//
	// This member is required.
	FailedItems []types.FailedItem

	// A unique identifier for the resource set, used in a TODO to refer to the
	// resource set.
	//
	// This member is required.
	ResourceSetIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchAssociateResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchAssociateResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchAssociateResource{}, middleware.After)
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
	if err = addOpBatchAssociateResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchAssociateResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchAssociateResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "BatchAssociateResource",
	}
}
