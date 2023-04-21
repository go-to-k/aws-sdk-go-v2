// Code generated by smithy-go-codegen DO NOT EDIT.

package grafana

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the current configuration string for the given workspace.
func (c *Client) DescribeWorkspaceConfiguration(ctx context.Context, params *DescribeWorkspaceConfigurationInput, optFns ...func(*Options)) (*DescribeWorkspaceConfigurationOutput, error) {
	if params == nil {
		params = &DescribeWorkspaceConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorkspaceConfiguration", params, optFns, c.addOperationDescribeWorkspaceConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorkspaceConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkspaceConfigurationInput struct {

	// The ID of the workspace to get configuration information for.
	//
	// This member is required.
	WorkspaceId *string

	noSmithyDocumentSerde
}

type DescribeWorkspaceConfigurationOutput struct {

	// The configuration string for the workspace that you requested. For more
	// information about the format and configuration options available, see Working
	// in your Grafana workspace (https://docs.aws.amazon.com/grafana/latest/userguide/AMG-configure-workspace.html)
	// .
	//
	// This value conforms to the media type: application/json
	//
	// This member is required.
	Configuration *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeWorkspaceConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeWorkspaceConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeWorkspaceConfiguration{}, middleware.After)
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
	if err = addOpDescribeWorkspaceConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkspaceConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeWorkspaceConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "grafana",
		OperationName: "DescribeWorkspaceConfiguration",
	}
}
