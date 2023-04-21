// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkidentity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the full details of an AppInstanceAdmin .
func (c *Client) DescribeAppInstanceAdmin(ctx context.Context, params *DescribeAppInstanceAdminInput, optFns ...func(*Options)) (*DescribeAppInstanceAdminOutput, error) {
	if params == nil {
		params = &DescribeAppInstanceAdminInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAppInstanceAdmin", params, optFns, c.addOperationDescribeAppInstanceAdminMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAppInstanceAdminOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAppInstanceAdminInput struct {

	// The ARN of the AppInstanceAdmin .
	//
	// This member is required.
	AppInstanceAdminArn *string

	// The ARN of the AppInstance .
	//
	// This member is required.
	AppInstanceArn *string

	noSmithyDocumentSerde
}

type DescribeAppInstanceAdminOutput struct {

	// The ARN and name of the AppInstanceUser , the ARN of the AppInstance , and the
	// created and last-updated timestamps. All timestamps use epoch milliseconds.
	AppInstanceAdmin *types.AppInstanceAdmin

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAppInstanceAdminMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAppInstanceAdmin{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAppInstanceAdmin{}, middleware.After)
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
	if err = addOpDescribeAppInstanceAdminValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAppInstanceAdmin(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAppInstanceAdmin(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DescribeAppInstanceAdmin",
	}
}
