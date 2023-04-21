// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the Resilience Hub application version.
func (c *Client) DescribeAppVersion(ctx context.Context, params *DescribeAppVersionInput, optFns ...func(*Options)) (*DescribeAppVersionOutput, error) {
	if params == nil {
		params = &DescribeAppVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAppVersion", params, optFns, c.addOperationDescribeAppVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAppVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAppVersionInput struct {

	// The Amazon Resource Name (ARN) of the Resilience Hub application. The format
	// for this ARN is: arn: partition :resiliencehub: region : account :app/ app-id .
	// For more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference guide.
	//
	// This member is required.
	AppArn *string

	// The Resilience Hub application version.
	//
	// This member is required.
	AppVersion *string

	noSmithyDocumentSerde
}

type DescribeAppVersionOutput struct {

	// The Amazon Resource Name (ARN) of the Resilience Hub application. The format
	// for this ARN is: arn: partition :resiliencehub: region : account :app/ app-id .
	// For more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference guide.
	//
	// This member is required.
	AppArn *string

	// The Resilience Hub application version.
	//
	// This member is required.
	AppVersion *string

	// Additional configuration parameters for an Resilience Hub application. If you
	// want to implement additionalInfo through the Resilience Hub console rather than
	// using an API call, see Configure the application configuration parameters (https://docs.aws.amazon.com/resilience-hub/latest/userguide/app-config-param.html)
	// . Currently, this parameter supports only failover region and account.
	AdditionalInfo map[string][]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAppVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAppVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAppVersion{}, middleware.After)
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
	if err = addOpDescribeAppVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAppVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAppVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "DescribeAppVersion",
	}
}
