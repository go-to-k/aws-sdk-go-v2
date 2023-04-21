// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the snapshot schedule for the specified gateway volume. The snapshot
// schedule information includes intervals at which snapshots are automatically
// initiated on the volume. This operation is only supported in the cached volume
// and stored volume types.
func (c *Client) DescribeSnapshotSchedule(ctx context.Context, params *DescribeSnapshotScheduleInput, optFns ...func(*Options)) (*DescribeSnapshotScheduleOutput, error) {
	if params == nil {
		params = &DescribeSnapshotScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSnapshotSchedule", params, optFns, c.addOperationDescribeSnapshotScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSnapshotScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the DescribeSnapshotScheduleInput$VolumeARN of the
// volume.
type DescribeSnapshotScheduleInput struct {

	// The Amazon Resource Name (ARN) of the volume. Use the ListVolumes operation to
	// return a list of gateway volumes.
	//
	// This member is required.
	VolumeARN *string

	noSmithyDocumentSerde
}

type DescribeSnapshotScheduleOutput struct {

	// The snapshot description.
	Description *string

	// The number of hours between snapshots.
	RecurrenceInHours *int32

	// The hour of the day at which the snapshot schedule begins represented as hh,
	// where hh is the hour (0 to 23). The hour of the day is in the time zone of the
	// gateway.
	StartAt *int32

	// A list of up to 50 tags assigned to the snapshot schedule, sorted
	// alphabetically by key name. Each tag is a key-value pair. For a gateway with
	// more than 10 tags assigned, you can view all tags using the ListTagsForResource
	// API operation.
	Tags []types.Tag

	// A value that indicates the time zone of the gateway.
	Timezone *string

	// The Amazon Resource Name (ARN) of the volume that was specified in the request.
	VolumeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSnapshotScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSnapshotSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSnapshotSchedule{}, middleware.After)
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
	if err = addOpDescribeSnapshotScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSnapshotSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSnapshotSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeSnapshotSchedule",
	}
}
