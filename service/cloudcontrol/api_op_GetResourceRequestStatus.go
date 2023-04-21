// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudcontrol

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns the current status of a resource operation request. For more
// information, see Tracking the progress of resource operation requests (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-manage-requests.html#resource-operations-manage-requests-track)
// in the Amazon Web Services Cloud Control API User Guide.
func (c *Client) GetResourceRequestStatus(ctx context.Context, params *GetResourceRequestStatusInput, optFns ...func(*Options)) (*GetResourceRequestStatusOutput, error) {
	if params == nil {
		params = &GetResourceRequestStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourceRequestStatus", params, optFns, c.addOperationGetResourceRequestStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourceRequestStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourceRequestStatusInput struct {

	// A unique token used to track the progress of the resource operation request.
	// Request tokens are included in the ProgressEvent type returned by a resource
	// operation request.
	//
	// This member is required.
	RequestToken *string

	noSmithyDocumentSerde
}

type GetResourceRequestStatusOutput struct {

	// Represents the current status of the resource operation request.
	ProgressEvent *types.ProgressEvent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourceRequestStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetResourceRequestStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetResourceRequestStatus{}, middleware.After)
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
	if err = addOpGetResourceRequestStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourceRequestStatus(options.Region), middleware.Before); err != nil {
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

// GetResourceRequestStatusAPIClient is a client that implements the
// GetResourceRequestStatus operation.
type GetResourceRequestStatusAPIClient interface {
	GetResourceRequestStatus(context.Context, *GetResourceRequestStatusInput, ...func(*Options)) (*GetResourceRequestStatusOutput, error)
}

var _ GetResourceRequestStatusAPIClient = (*Client)(nil)

// ResourceRequestSuccessWaiterOptions are waiter options for
// ResourceRequestSuccessWaiter
type ResourceRequestSuccessWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ResourceRequestSuccessWaiter will use default minimum delay of 5 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ResourceRequestSuccessWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetResourceRequestStatusInput, *GetResourceRequestStatusOutput, error) (bool, error)
}

// ResourceRequestSuccessWaiter defines the waiters for ResourceRequestSuccess
type ResourceRequestSuccessWaiter struct {
	client GetResourceRequestStatusAPIClient

	options ResourceRequestSuccessWaiterOptions
}

// NewResourceRequestSuccessWaiter constructs a ResourceRequestSuccessWaiter.
func NewResourceRequestSuccessWaiter(client GetResourceRequestStatusAPIClient, optFns ...func(*ResourceRequestSuccessWaiterOptions)) *ResourceRequestSuccessWaiter {
	options := ResourceRequestSuccessWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = resourceRequestSuccessStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ResourceRequestSuccessWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ResourceRequestSuccess waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ResourceRequestSuccessWaiter) Wait(ctx context.Context, params *GetResourceRequestStatusInput, maxWaitDur time.Duration, optFns ...func(*ResourceRequestSuccessWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ResourceRequestSuccess waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *ResourceRequestSuccessWaiter) WaitForOutput(ctx context.Context, params *GetResourceRequestStatusInput, maxWaitDur time.Duration, optFns ...func(*ResourceRequestSuccessWaiterOptions)) (*GetResourceRequestStatusOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetResourceRequestStatus(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for ResourceRequestSuccess waiter")
}

func resourceRequestSuccessStateRetryable(ctx context.Context, input *GetResourceRequestStatusInput, output *GetResourceRequestStatusOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("ProgressEvent.OperationStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "SUCCESS"
		value, ok := pathValue.(types.OperationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.OperationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ProgressEvent.OperationStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.OperationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.OperationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ProgressEvent.OperationStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CANCEL_COMPLETE"
		value, ok := pathValue.(types.OperationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.OperationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetResourceRequestStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudcontrolapi",
		OperationName: "GetResourceRequestStatus",
	}
}
