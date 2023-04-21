// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Rekognition stream processor that you can use to detect and
// recognize faces or to detect labels in a streaming video. Amazon Rekognition
// Video is a consumer of live video from Amazon Kinesis Video Streams. There are
// two different settings for stream processors in Amazon Rekognition: detecting
// faces and detecting labels.
//   - If you are creating a stream processor for detecting faces, you provide as
//     input a Kinesis video stream ( Input ) and a Kinesis data stream ( Output )
//     stream for receiving the output. You must use the FaceSearch option in
//     Settings , specifying the collection that contains the faces you want to
//     recognize. After you have finished analyzing a streaming video, use
//     StopStreamProcessor to stop processing.
//   - If you are creating a stream processor to detect labels, you provide as
//     input a Kinesis video stream ( Input ), Amazon S3 bucket information ( Output
//     ), and an Amazon SNS topic ARN ( NotificationChannel ). You can also provide a
//     KMS key ID to encrypt the data sent to your Amazon S3 bucket. You specify what
//     you want to detect by using the ConnectedHome option in settings, and
//     selecting one of the following: PERSON , PET , PACKAGE , ALL You can also
//     specify where in the frame you want Amazon Rekognition to monitor with
//     RegionsOfInterest . When you run the StartStreamProcessor operation on a label
//     detection stream processor, you input start and stop information to determine
//     the length of the processing time.
//
// Use Name to assign an identifier for the stream processor. You use Name to
// manage the stream processor. For example, you can start processing the source
// video by calling StartStreamProcessor with the Name field. This operation
// requires permissions to perform the rekognition:CreateStreamProcessor action.
// If you want to tag your stream processor, you also require permission to perform
// the rekognition:TagResource operation.
func (c *Client) CreateStreamProcessor(ctx context.Context, params *CreateStreamProcessorInput, optFns ...func(*Options)) (*CreateStreamProcessorOutput, error) {
	if params == nil {
		params = &CreateStreamProcessorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStreamProcessor", params, optFns, c.addOperationCreateStreamProcessorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStreamProcessorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStreamProcessorInput struct {

	// Kinesis video stream stream that provides the source streaming video. If you
	// are using the AWS CLI, the parameter name is StreamProcessorInput . This is
	// required for both face search and label detection stream processors.
	//
	// This member is required.
	Input *types.StreamProcessorInput

	// An identifier you assign to the stream processor. You can use Name to manage
	// the stream processor. For example, you can get the current status of the stream
	// processor by calling DescribeStreamProcessor . Name is idempotent. This is
	// required for both face search and label detection stream processors.
	//
	// This member is required.
	Name *string

	// Kinesis data stream stream or Amazon S3 bucket location to which Amazon
	// Rekognition Video puts the analysis results. If you are using the AWS CLI, the
	// parameter name is StreamProcessorOutput . This must be a S3Destination of an
	// Amazon S3 bucket that you own for a label detection stream processor or a
	// Kinesis data stream ARN for a face search stream processor.
	//
	// This member is required.
	Output *types.StreamProcessorOutput

	// The Amazon Resource Number (ARN) of the IAM role that allows access to the
	// stream processor. The IAM role provides Rekognition read permissions for a
	// Kinesis stream. It also provides write permissions to an Amazon S3 bucket and
	// Amazon Simple Notification Service topic for a label detection stream processor.
	// This is required for both face search and label detection stream processors.
	//
	// This member is required.
	RoleArn *string

	// Input parameters used in a streaming video analyzed by a stream processor. You
	// can use FaceSearch to recognize faces in a streaming video, or you can use
	// ConnectedHome to detect labels.
	//
	// This member is required.
	Settings *types.StreamProcessorSettings

	// Shows whether you are sharing data with Rekognition to improve model
	// performance. You can choose this option at the account level or on a per-stream
	// basis. Note that if you opt out at the account level this setting is ignored on
	// individual streams.
	DataSharingPreference *types.StreamProcessorDataSharingPreference

	// The identifier for your AWS Key Management Service key (AWS KMS key). This is
	// an optional parameter for label detection stream processors and should not be
	// used to create a face search stream processor. You can supply the Amazon
	// Resource Name (ARN) of your KMS key, the ID of your KMS key, an alias for your
	// KMS key, or an alias ARN. The key is used to encrypt results and data published
	// to your Amazon S3 bucket, which includes image frames and hero images. Your
	// source images are unaffected.
	KmsKeyId *string

	// The Amazon Simple Notification Service topic to which Amazon Rekognition
	// publishes the object detection results and completion status of a video analysis
	// operation. Amazon Rekognition publishes a notification the first time an object
	// of interest or a person is detected in the video stream. For example, if Amazon
	// Rekognition detects a person at second 2, a pet at second 4, and a person again
	// at second 5, Amazon Rekognition sends 2 object class detected notifications, one
	// for a person at second 2 and one for a pet at second 4. Amazon Rekognition also
	// publishes an an end-of-session notification with a summary when the stream
	// processing session is complete.
	NotificationChannel *types.StreamProcessorNotificationChannel

	// Specifies locations in the frames where Amazon Rekognition checks for objects
	// or people. You can specify up to 10 regions of interest, and each region has
	// either a polygon or a bounding box. This is an optional parameter for label
	// detection stream processors and should not be used to create a face search
	// stream processor.
	RegionsOfInterest []types.RegionOfInterest

	// A set of tags (key-value pairs) that you want to attach to the stream processor.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateStreamProcessorOutput struct {

	// Amazon Resource Number for the newly created stream processor.
	StreamProcessorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStreamProcessorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateStreamProcessor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateStreamProcessor{}, middleware.After)
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
	if err = addOpCreateStreamProcessorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStreamProcessor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateStreamProcessor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "CreateStreamProcessor",
	}
}
