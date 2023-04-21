// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a program within a channel. For information about programs, see
// Working with programs (https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-programs.html)
// in the MediaTailor User Guide.
func (c *Client) DescribeProgram(ctx context.Context, params *DescribeProgramInput, optFns ...func(*Options)) (*DescribeProgramOutput, error) {
	if params == nil {
		params = &DescribeProgramInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProgram", params, optFns, c.addOperationDescribeProgramMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProgramOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProgramInput struct {

	// The name of the channel associated with this Program.
	//
	// This member is required.
	ChannelName *string

	// The name of the program.
	//
	// This member is required.
	ProgramName *string

	noSmithyDocumentSerde
}

type DescribeProgramOutput struct {

	// The ad break configuration settings.
	AdBreaks []types.AdBreak

	// The ARN of the program.
	Arn *string

	// The name of the channel that the program belongs to.
	ChannelName *string

	// The clip range configuration settings.
	ClipRange *types.ClipRange

	// The timestamp of when the program was created.
	CreationTime *time.Time

	// The duration of the live program in milliseconds.
	DurationMillis *int64

	// The name of the LiveSource for this Program.
	LiveSourceName *string

	// The name of the program.
	ProgramName *string

	// The date and time that the program is scheduled to start in ISO 8601 format and
	// Coordinated Universal Time (UTC). For example, the value
	// 2021-03-27T17:48:16.751Z represents March 27, 2021 at 17:48:16.751 UTC.
	ScheduledStartTime *time.Time

	// The source location name.
	SourceLocationName *string

	// The name that's used to refer to a VOD source.
	VodSourceName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeProgramMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeProgram{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeProgram{}, middleware.After)
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
	if err = addOpDescribeProgramValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProgram(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeProgram(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "DescribeProgram",
	}
}
