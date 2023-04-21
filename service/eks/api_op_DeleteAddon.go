// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete an Amazon EKS add-on. When you remove the add-on, it will also be
// deleted from the cluster. You can always manually start an add-on on the cluster
// using the Kubernetes API.
func (c *Client) DeleteAddon(ctx context.Context, params *DeleteAddonInput, optFns ...func(*Options)) (*DeleteAddonOutput, error) {
	if params == nil {
		params = &DeleteAddonInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAddon", params, optFns, c.addOperationDeleteAddonMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAddonOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAddonInput struct {

	// The name of the add-on. The name must match one of the names returned by
	// ListAddons (https://docs.aws.amazon.com/eks/latest/APIReference/API_ListAddons.html)
	// .
	//
	// This member is required.
	AddonName *string

	// The name of the cluster to delete the add-on from.
	//
	// This member is required.
	ClusterName *string

	// Specifying this option preserves the add-on software on your cluster but Amazon
	// EKS stops managing any settings for the add-on. If an IAM account is associated
	// with the add-on, it isn't removed.
	Preserve bool

	noSmithyDocumentSerde
}

type DeleteAddonOutput struct {

	// An Amazon EKS add-on. For more information, see Amazon EKS add-ons (https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html)
	// in the Amazon EKS User Guide.
	Addon *types.Addon

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAddonMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteAddon{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteAddon{}, middleware.After)
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
	if err = addOpDeleteAddonValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAddon(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAddon(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "DeleteAddon",
	}
}
