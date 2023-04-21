// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a continuous deployment policy, including metadata (the policy's
// identifier and the date and time when the policy was last modified).
func (c *Client) GetContinuousDeploymentPolicy(ctx context.Context, params *GetContinuousDeploymentPolicyInput, optFns ...func(*Options)) (*GetContinuousDeploymentPolicyOutput, error) {
	if params == nil {
		params = &GetContinuousDeploymentPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetContinuousDeploymentPolicy", params, optFns, c.addOperationGetContinuousDeploymentPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetContinuousDeploymentPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContinuousDeploymentPolicyInput struct {

	// The identifier of the continuous deployment policy that you are getting.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetContinuousDeploymentPolicyOutput struct {

	// A continuous deployment policy.
	ContinuousDeploymentPolicy *types.ContinuousDeploymentPolicy

	// The version identifier for the current version of the continuous deployment
	// policy.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetContinuousDeploymentPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetContinuousDeploymentPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetContinuousDeploymentPolicy{}, middleware.After)
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
	if err = addOpGetContinuousDeploymentPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetContinuousDeploymentPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetContinuousDeploymentPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetContinuousDeploymentPolicy",
	}
}
