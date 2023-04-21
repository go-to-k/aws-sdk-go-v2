// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a batch inference job. The operation can handle up to 50 million
// records and the input file must be in JSON format. For more information, see
// Creating a batch inference job (https://docs.aws.amazon.com/personalize/latest/dg/creating-batch-inference-job.html)
// .
func (c *Client) CreateBatchInferenceJob(ctx context.Context, params *CreateBatchInferenceJobInput, optFns ...func(*Options)) (*CreateBatchInferenceJobOutput, error) {
	if params == nil {
		params = &CreateBatchInferenceJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBatchInferenceJob", params, optFns, c.addOperationCreateBatchInferenceJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBatchInferenceJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBatchInferenceJobInput struct {

	// The Amazon S3 path that leads to the input file to base your recommendations
	// on. The input material must be in JSON format.
	//
	// This member is required.
	JobInput *types.BatchInferenceJobInput

	// The name of the batch inference job to create.
	//
	// This member is required.
	JobName *string

	// The path to the Amazon S3 bucket where the job's output will be stored.
	//
	// This member is required.
	JobOutput *types.BatchInferenceJobOutput

	// The ARN of the Amazon Identity and Access Management role that has permissions
	// to read and write to your input and output Amazon S3 buckets respectively.
	//
	// This member is required.
	RoleArn *string

	// The Amazon Resource Name (ARN) of the solution version that will be used to
	// generate the batch inference recommendations.
	//
	// This member is required.
	SolutionVersionArn *string

	// The configuration details of a batch inference job.
	BatchInferenceJobConfig *types.BatchInferenceJobConfig

	// The ARN of the filter to apply to the batch inference job. For more information
	// on using filters, see Filtering batch recommendations (https://docs.aws.amazon.com/personalize/latest/dg/filter-batch.html)
	// .
	FilterArn *string

	// The number of recommendations to retrieve.
	NumResults *int32

	// A list of tags (https://docs.aws.amazon.com/personalize/latest/dev/tagging-resources.html)
	// to apply to the batch inference job.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateBatchInferenceJobOutput struct {

	// The ARN of the batch inference job.
	BatchInferenceJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBatchInferenceJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBatchInferenceJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBatchInferenceJob{}, middleware.After)
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
	if err = addOpCreateBatchInferenceJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBatchInferenceJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBatchInferenceJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateBatchInferenceJob",
	}
}
