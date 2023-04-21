// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a new task from the specified task definition on the specified container
// instance or instances. Starting April 15, 2023, Amazon Web Services will not
// onboard new customers to Amazon Elastic Inference (EI), and will help current
// customers migrate their workloads to options that offer better price and
// performance. After April 15, 2023, new customers will not be able to launch
// instances with Amazon EI accelerators in Amazon SageMaker, Amazon ECS, or Amazon
// EC2. However, customers who have used Amazon EI at least once during the past
// 30-day period are considered current customers and will be able to continue
// using the service. Alternatively, you can use RunTask to place tasks for you.
// For more information, see Scheduling Tasks (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/scheduling_tasks.html)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) StartTask(ctx context.Context, params *StartTaskInput, optFns ...func(*Options)) (*StartTaskOutput, error) {
	if params == nil {
		params = &StartTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartTask", params, optFns, c.addOperationStartTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartTaskInput struct {

	// The container instance IDs or full ARN entries for the container instances
	// where you would like to place your task. You can specify up to 10 container
	// instances.
	//
	// This member is required.
	ContainerInstances []string

	// The family and revision ( family:revision ) or full ARN of the task definition
	// to start. If a revision isn't specified, the latest ACTIVE revision is used.
	//
	// This member is required.
	TaskDefinition *string

	// The short name or full Amazon Resource Name (ARN) of the cluster where to start
	// your task. If you do not specify a cluster, the default cluster is assumed.
	Cluster *string

	// Specifies whether to use Amazon ECS managed tags for the task. For more
	// information, see Tagging Your Amazon ECS Resources (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// in the Amazon Elastic Container Service Developer Guide.
	EnableECSManagedTags bool

	// Whether or not the execute command functionality is turned on for the task. If
	// true , this enables execute command functionality on all containers in the task.
	EnableExecuteCommand bool

	// The name of the task group to associate with the task. The default value is the
	// family name of the task definition (for example, family:my-family-name).
	Group *string

	// The VPC subnet and security group configuration for tasks that receive their
	// own elastic network interface by using the awsvpc networking mode.
	NetworkConfiguration *types.NetworkConfiguration

	// A list of container overrides in JSON format that specify the name of a
	// container in the specified task definition and the overrides it receives. You
	// can override the default command for a container (that's specified in the task
	// definition or Docker image) with a command override. You can also override
	// existing environment variables (that are specified in the task definition or
	// Docker image) on a container or add new environment variables to it with an
	// environment override. A total of 8192 characters are allowed for overrides. This
	// limit includes the JSON formatting characters of the override structure.
	Overrides *types.TaskOverride

	// Specifies whether to propagate the tags from the task definition or the service
	// to the task. If no value is specified, the tags aren't propagated.
	PropagateTags types.PropagateTags

	// The reference ID to use for the task.
	ReferenceId *string

	// An optional tag specified when a task is started. For example, if you
	// automatically trigger a task to run a batch process job, you could apply a
	// unique identifier for that job to your task with the startedBy parameter. You
	// can then identify which tasks belong to that job by filtering the results of a
	// ListTasks call with the startedBy value. Up to 36 letters (uppercase and
	// lowercase), numbers, hyphens (-), and underscores (_) are allowed. If a task is
	// started by an Amazon ECS service, the startedBy parameter contains the
	// deployment ID of the service that starts it.
	StartedBy *string

	// The metadata that you apply to the task to help you categorize and organize
	// them. Each tag consists of a key and an optional value, both of which you
	// define. The following basic restrictions apply to tags:
	//   - Maximum number of tags per resource - 50
	//   - For each resource, each tag key must be unique, and each tag key can have
	//   only one value.
	//   - Maximum key length - 128 Unicode characters in UTF-8
	//   - Maximum value length - 256 Unicode characters in UTF-8
	//   - If your tagging schema is used across multiple services and resources,
	//   remember that other services may have restrictions on allowed characters.
	//   Generally allowed characters are: letters, numbers, and spaces representable in
	//   UTF-8, and the following characters: + - = . _ : / @.
	//   - Tag keys and values are case-sensitive.
	//   - Do not use aws: , AWS: , or any upper or lowercase combination of such as a
	//   prefix for either keys or values as it is reserved for Amazon Web Services use.
	//   You cannot edit or delete tag keys or values with this prefix. Tags with this
	//   prefix do not count against your tags per resource limit.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type StartTaskOutput struct {

	// Any failures associated with the call.
	Failures []types.Failure

	// A full description of the tasks that were started. Each task that was
	// successfully placed on your container instances is described.
	Tasks []types.Task

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartTask{}, middleware.After)
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
	if err = addOpStartTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "StartTask",
	}
}
