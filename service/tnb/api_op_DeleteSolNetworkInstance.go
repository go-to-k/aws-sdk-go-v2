// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a network instance. A network instance is a single network created in
// Amazon Web Services TNB that can be deployed and on which life-cycle operations
// (like terminate, update, and delete) can be performed. To delete a network
// instance, the instance must be in a stopped or terminated state. To terminate a
// network instance, see TerminateSolNetworkInstance (https://docs.aws.amazon.com/tnb/latest/APIReference/API_TerminateSolNetworkInstance.html)
// .
func (c *Client) DeleteSolNetworkInstance(ctx context.Context, params *DeleteSolNetworkInstanceInput, optFns ...func(*Options)) (*DeleteSolNetworkInstanceOutput, error) {
	if params == nil {
		params = &DeleteSolNetworkInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSolNetworkInstance", params, optFns, c.addOperationDeleteSolNetworkInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSolNetworkInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSolNetworkInstanceInput struct {

	// Network instance ID.
	//
	// This member is required.
	NsInstanceId *string

	noSmithyDocumentSerde
}

type DeleteSolNetworkInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSolNetworkInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteSolNetworkInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteSolNetworkInstance{}, middleware.After)
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
	if err = addOpDeleteSolNetworkInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSolNetworkInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSolNetworkInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tnb",
		OperationName: "DeleteSolNetworkInstance",
	}
}
