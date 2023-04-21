// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes an Amazon Forecast dataset created using the CreateDataset (https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDataset.html)
// operation. In addition to listing the parameters specified in the CreateDataset
// request, this operation includes the following dataset properties:
//   - CreationTime
//   - LastModificationTime
//   - Status
func (c *Client) DescribeDataset(ctx context.Context, params *DescribeDatasetInput, optFns ...func(*Options)) (*DescribeDatasetOutput, error) {
	if params == nil {
		params = &DescribeDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDataset", params, optFns, c.addOperationDescribeDatasetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDatasetInput struct {

	// The Amazon Resource Name (ARN) of the dataset.
	//
	// This member is required.
	DatasetArn *string

	noSmithyDocumentSerde
}

type DescribeDatasetOutput struct {

	// When the dataset was created.
	CreationTime *time.Time

	// The frequency of data collection. Valid intervals are Y (Year), M (Month), W
	// (Week), D (Day), H (Hour), 30min (30 minutes), 15min (15 minutes), 10min (10
	// minutes), 5min (5 minutes), and 1min (1 minute). For example, "M" indicates
	// every month and "30min" indicates every 30 minutes.
	DataFrequency *string

	// The Amazon Resource Name (ARN) of the dataset.
	DatasetArn *string

	// The name of the dataset.
	DatasetName *string

	// The dataset type.
	DatasetType types.DatasetType

	// The domain associated with the dataset.
	Domain types.Domain

	// The Key Management Service (KMS) key and the Identity and Access Management
	// (IAM) role that Amazon Forecast can assume to access the key.
	EncryptionConfig *types.EncryptionConfig

	// When you create a dataset, LastModificationTime is the same as CreationTime .
	// While data is being imported to the dataset, LastModificationTime is the
	// current time of the DescribeDataset call. After a CreateDatasetImportJob (https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetImportJob.html)
	// operation has finished, LastModificationTime is when the import job completed
	// or failed.
	LastModificationTime *time.Time

	// An array of SchemaAttribute objects that specify the dataset fields. Each
	// SchemaAttribute specifies the name and data type of a field.
	Schema *types.Schema

	// The status of the dataset. States include:
	//   - ACTIVE
	//   - CREATE_PENDING , CREATE_IN_PROGRESS , CREATE_FAILED
	//   - DELETE_PENDING , DELETE_IN_PROGRESS , DELETE_FAILED
	//   - UPDATE_PENDING , UPDATE_IN_PROGRESS , UPDATE_FAILED
	// The UPDATE states apply while data is imported to the dataset from a call to
	// the CreateDatasetImportJob (https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDatasetImportJob.html)
	// operation and reflect the status of the dataset import job. For example, when
	// the import job status is CREATE_IN_PROGRESS , the status of the dataset is
	// UPDATE_IN_PROGRESS . The Status of the dataset must be ACTIVE before you can
	// import training data.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDatasetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDataset{}, middleware.After)
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
	if err = addOpDescribeDatasetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDataset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDataset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DescribeDataset",
	}
}
