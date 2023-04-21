// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables your application to create a new flow using Amazon AppFlow. You must
// create a connector profile before calling this API. Please note that the Request
// Syntax below shows syntax for multiple destinations, however, you can only
// transfer data to one item in this list at a time. Amazon AppFlow does not
// currently support flows to multiple destinations at once.
func (c *Client) CreateFlow(ctx context.Context, params *CreateFlowInput, optFns ...func(*Options)) (*CreateFlowOutput, error) {
	if params == nil {
		params = &CreateFlowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFlow", params, optFns, c.addOperationCreateFlowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFlowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFlowInput struct {

	// The configuration that controls how Amazon AppFlow places data in the
	// destination connector.
	//
	// This member is required.
	DestinationFlowConfigList []types.DestinationFlowConfig

	// The specified name of the flow. Spaces are not allowed. Use underscores (_) or
	// hyphens (-) only.
	//
	// This member is required.
	FlowName *string

	// The configuration that controls how Amazon AppFlow retrieves data from the
	// source connector.
	//
	// This member is required.
	SourceFlowConfig *types.SourceFlowConfig

	// A list of tasks that Amazon AppFlow performs while transferring the data in the
	// flow run.
	//
	// This member is required.
	Tasks []types.Task

	// The trigger settings that determine how and when the flow runs.
	//
	// This member is required.
	TriggerConfig *types.TriggerConfig

	// A description of the flow you want to create.
	Description *string

	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you
	// provide for encryption. This is required if you do not want to use the Amazon
	// AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses
	// the Amazon AppFlow-managed KMS key.
	KmsArn *string

	// Specifies the configuration that Amazon AppFlow uses when it catalogs the data
	// that's transferred by the associated flow. When Amazon AppFlow catalogs the data
	// from a flow, it stores metadata in a data catalog.
	MetadataCatalogConfig *types.MetadataCatalogConfig

	// The tags used to organize, track, or control access for your flow.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateFlowOutput struct {

	// The flow's Amazon Resource Name (ARN).
	FlowArn *string

	// Indicates the current status of the flow.
	FlowStatus types.FlowStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFlowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFlow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFlow{}, middleware.After)
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
	if err = addOpCreateFlowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFlow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFlow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appflow",
		OperationName: "CreateFlow",
	}
}
