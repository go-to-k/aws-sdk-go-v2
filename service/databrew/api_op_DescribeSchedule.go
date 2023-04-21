// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the definition of a specific DataBrew schedule.
func (c *Client) DescribeSchedule(ctx context.Context, params *DescribeScheduleInput, optFns ...func(*Options)) (*DescribeScheduleOutput, error) {
	if params == nil {
		params = &DescribeScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSchedule", params, optFns, c.addOperationDescribeScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScheduleInput struct {

	// The name of the schedule to be described.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DescribeScheduleOutput struct {

	// The name of the schedule.
	//
	// This member is required.
	Name *string

	// The date and time that the schedule was created.
	CreateDate *time.Time

	// The identifier (user name) of the user who created the schedule.
	CreatedBy *string

	// The date or dates and time or times when the jobs are to be run for the
	// schedule. For more information, see Cron expressions (https://docs.aws.amazon.com/databrew/latest/dg/jobs.cron.html)
	// in the Glue DataBrew Developer Guide.
	CronExpression *string

	// The name or names of one or more jobs to be run by using the schedule.
	JobNames []string

	// The identifier (user name) of the user who last modified the schedule.
	LastModifiedBy *string

	// The date and time that the schedule was last modified.
	LastModifiedDate *time.Time

	// The Amazon Resource Name (ARN) of the schedule.
	ResourceArn *string

	// Metadata tags associated with this schedule.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSchedule{}, middleware.After)
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
	if err = addOpDescribeScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "databrew",
		OperationName: "DescribeSchedule",
	}
}
