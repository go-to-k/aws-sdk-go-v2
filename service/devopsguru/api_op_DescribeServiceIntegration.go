// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devopsguru/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the integration status of services that are integrated with DevOps
// Guru. The one service that can be integrated with DevOps Guru is Amazon Web
// Services Systems Manager, which can be used to create an OpsItem for each
// generated insight.
func (c *Client) DescribeServiceIntegration(ctx context.Context, params *DescribeServiceIntegrationInput, optFns ...func(*Options)) (*DescribeServiceIntegrationOutput, error) {
	if params == nil {
		params = &DescribeServiceIntegrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeServiceIntegration", params, optFns, c.addOperationDescribeServiceIntegrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeServiceIntegrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeServiceIntegrationInput struct {
	noSmithyDocumentSerde
}

type DescribeServiceIntegrationOutput struct {

	// Information about the integration of DevOps Guru with another Amazon Web
	// Services service, such as Amazon Web Services Systems Manager.
	ServiceIntegration *types.ServiceIntegrationConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeServiceIntegrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeServiceIntegration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeServiceIntegration{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeServiceIntegration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeServiceIntegration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "DescribeServiceIntegration",
	}
}
