// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an endpoint configuration that SageMaker hosting services uses to
// deploy models. In the configuration, you identify one or more models, created
// using the CreateModel API, to deploy and the resources that you want SageMaker
// to provision. Then you call the CreateEndpoint API. Use this API if you want to
// use SageMaker hosting services to deploy models into production. In the request,
// you define a ProductionVariant , for each model that you want to deploy. Each
// ProductionVariant parameter also describes the resources that you want SageMaker
// to provision. This includes the number and type of ML compute instances to
// deploy. If you are hosting multiple models, you also assign a VariantWeight to
// specify how much traffic you want to allocate to each model. For example,
// suppose that you want to host two models, A and B, and you assign traffic weight
// 2 for model A and 1 for model B. SageMaker distributes two-thirds of the traffic
// to Model A, and one-third to model B. When you call CreateEndpoint , a load call
// is made to DynamoDB to verify that your endpoint configuration exists. When you
// read data from a DynamoDB table supporting Eventually Consistent Reads (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadConsistency.html)
// , the response might not reflect the results of a recently completed write
// operation. The response might include some stale data. If the dependent entities
// are not yet in DynamoDB, this causes a validation error. If you repeat your read
// request after a short time, the response should return the latest data. So retry
// logic is recommended to handle these possible issues. We also recommend that
// customers call DescribeEndpointConfig before calling CreateEndpoint to minimize
// the potential impact of a DynamoDB eventually consistent read.
func (c *Client) CreateEndpointConfig(ctx context.Context, params *CreateEndpointConfigInput, optFns ...func(*Options)) (*CreateEndpointConfigOutput, error) {
	if params == nil {
		params = &CreateEndpointConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEndpointConfig", params, optFns, c.addOperationCreateEndpointConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEndpointConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEndpointConfigInput struct {

	// The name of the endpoint configuration. You specify this name in a
	// CreateEndpoint request.
	//
	// This member is required.
	EndpointConfigName *string

	// An array of ProductionVariant objects, one for each model that you want to host
	// at this endpoint.
	//
	// This member is required.
	ProductionVariants []types.ProductionVariant

	// Specifies configuration for how an endpoint performs asynchronous inference.
	// This is a required field in order for your Endpoint to be invoked using
	// InvokeEndpointAsync (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpointAsync.html)
	// .
	AsyncInferenceConfig *types.AsyncInferenceConfig

	// Configuration to control how SageMaker captures inference data.
	DataCaptureConfig *types.DataCaptureConfig

	// A member of CreateEndpointConfig that enables explainers.
	ExplainerConfig *types.ExplainerConfig

	// The Amazon Resource Name (ARN) of a Amazon Web Services Key Management Service
	// key that SageMaker uses to encrypt data on the storage volume attached to the ML
	// compute instance that hosts the endpoint. The KmsKeyId can be any of the
	// following formats:
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Key ARN:
	//   arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Alias name: alias/ExampleAlias
	//   - Alias name ARN: arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias
	// The KMS key policy must grant permission to the IAM role that you specify in
	// your CreateEndpoint , UpdateEndpoint requests. For more information, refer to
	// the Amazon Web Services Key Management Service section Using Key Policies in
	// Amazon Web Services KMS  (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html)
	// Certain Nitro-based instances include local storage, dependent on the instance
	// type. Local storage volumes are encrypted using a hardware module on the
	// instance. You can't request a KmsKeyId when using an instance type with local
	// storage. If any of the models that you specify in the ProductionVariants
	// parameter use nitro-based instances with local storage, do not specify a value
	// for the KmsKeyId parameter. If you specify a value for KmsKeyId when using any
	// nitro-based instances with local storage, the call to CreateEndpointConfig
	// fails. For a list of instance types that support local instance storage, see
	// Instance Store Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#instance-store-volumes)
	// . For more information about local instance storage encryption, see SSD
	// Instance Store Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ssd-instance-store.html)
	// .
	KmsKeyId *string

	// An array of ProductionVariant objects, one for each model that you want to host
	// at this endpoint in shadow mode with production traffic replicated from the
	// model specified on ProductionVariants . If you use this field, you can only
	// specify one variant for ProductionVariants and one variant for
	// ShadowProductionVariants .
	ShadowProductionVariants []types.ProductionVariant

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// .
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateEndpointConfigOutput struct {

	// The Amazon Resource Name (ARN) of the endpoint configuration.
	//
	// This member is required.
	EndpointConfigArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEndpointConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEndpointConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEndpointConfig{}, middleware.After)
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
	if err = addOpCreateEndpointConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEndpointConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEndpointConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateEndpointConfig",
	}
}
