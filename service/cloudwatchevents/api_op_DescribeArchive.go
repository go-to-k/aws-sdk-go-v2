// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves details about an archive.
func (c *Client) DescribeArchive(ctx context.Context, params *DescribeArchiveInput, optFns ...func(*Options)) (*DescribeArchiveOutput, error) {
	if params == nil {
		params = &DescribeArchiveInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeArchive", params, optFns, c.addOperationDescribeArchiveMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeArchiveOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeArchiveInput struct {

	// The name of the archive to retrieve.
	//
	// This member is required.
	ArchiveName *string

	noSmithyDocumentSerde
}

type DescribeArchiveOutput struct {

	// The ARN of the archive.
	ArchiveArn *string

	// The name of the archive.
	ArchiveName *string

	// The time at which the archive was created.
	CreationTime *time.Time

	// The description of the archive.
	Description *string

	// The number of events in the archive.
	EventCount int64

	// The event pattern used to filter events sent to the archive.
	EventPattern *string

	// The ARN of the event source associated with the archive.
	EventSourceArn *string

	// The number of days to retain events for in the archive.
	RetentionDays *int32

	// The size of the archive in bytes.
	SizeBytes int64

	// The state of the archive.
	State types.ArchiveState

	// The reason that the archive is in the state.
	StateReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeArchiveMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeArchive{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeArchive{}, middleware.After)
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
	if err = addOpDescribeArchiveValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeArchive(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeArchive(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DescribeArchive",
	}
}
