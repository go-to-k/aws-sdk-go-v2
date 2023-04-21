// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a Device Defender ML Detect mitigation actions task. Requires permission
// to access the StartDetectMitigationActionsTask (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) StartDetectMitigationActionsTask(ctx context.Context, params *StartDetectMitigationActionsTaskInput, optFns ...func(*Options)) (*StartDetectMitigationActionsTaskOutput, error) {
	if params == nil {
		params = &StartDetectMitigationActionsTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDetectMitigationActionsTask", params, optFns, c.addOperationStartDetectMitigationActionsTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDetectMitigationActionsTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDetectMitigationActionsTaskInput struct {

	// The actions to be performed when a device has unexpected behavior.
	//
	// This member is required.
	Actions []string

	// Each mitigation action task must have a unique client request token. If you try
	// to create a new task with the same token as a task that already exists, an
	// exception occurs. If you omit this value, Amazon Web Services SDKs will
	// automatically generate a unique client request.
	//
	// This member is required.
	ClientRequestToken *string

	// Specifies the ML Detect findings to which the mitigation actions are applied.
	//
	// This member is required.
	Target *types.DetectMitigationActionsTaskTarget

	// The unique identifier of the task.
	//
	// This member is required.
	TaskId *string

	// Specifies to list only active violations.
	IncludeOnlyActiveViolations *bool

	// Specifies to include suppressed alerts.
	IncludeSuppressedAlerts *bool

	// Specifies the time period of which violation events occurred between.
	ViolationEventOccurrenceRange *types.ViolationEventOccurrenceRange

	noSmithyDocumentSerde
}

type StartDetectMitigationActionsTaskOutput struct {

	// The unique identifier of the task.
	TaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartDetectMitigationActionsTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartDetectMitigationActionsTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartDetectMitigationActionsTask{}, middleware.After)
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
	if err = addIdempotencyToken_opStartDetectMitigationActionsTaskMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartDetectMitigationActionsTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDetectMitigationActionsTask(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartDetectMitigationActionsTask struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartDetectMitigationActionsTask) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartDetectMitigationActionsTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartDetectMitigationActionsTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartDetectMitigationActionsTaskInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartDetectMitigationActionsTaskMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartDetectMitigationActionsTask{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartDetectMitigationActionsTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "StartDetectMitigationActionsTask",
	}
}
