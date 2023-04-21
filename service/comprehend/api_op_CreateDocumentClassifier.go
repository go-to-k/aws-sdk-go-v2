// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new document classifier that you can use to categorize documents. To
// create a classifier, you provide a set of training documents that labeled with
// the categories that you want to use. After the classifier is trained you can use
// it to categorize a set of labeled documents into the categories. For more
// information, see Document Classification (https://docs.aws.amazon.com/comprehend/latest/dg/how-document-classification.html)
// in the Comprehend Developer Guide.
func (c *Client) CreateDocumentClassifier(ctx context.Context, params *CreateDocumentClassifierInput, optFns ...func(*Options)) (*CreateDocumentClassifierOutput, error) {
	if params == nil {
		params = &CreateDocumentClassifierInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDocumentClassifier", params, optFns, c.addOperationCreateDocumentClassifierMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDocumentClassifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDocumentClassifierInput struct {

	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend
	// read access to your input data.
	//
	// This member is required.
	DataAccessRoleArn *string

	// The name of the document classifier.
	//
	// This member is required.
	DocumentClassifierName *string

	// Specifies the format and location of the input data for the job.
	//
	// This member is required.
	InputDataConfig *types.DocumentClassifierInputDataConfig

	// The language of the input documents. You can specify any of the languages
	// supported by Amazon Comprehend. All documents must be in the same language.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// A unique identifier for the request. If you don't set the client request token,
	// Amazon Comprehend generates one.
	ClientRequestToken *string

	// Indicates the mode in which the classifier will be trained. The classifier can
	// be trained in multi-class mode, which identifies one and only one class for each
	// document, or multi-label mode, which identifies one or more labels for each
	// document. In multi-label mode, multiple labels for an individual document are
	// separated by a delimiter. The default delimiter between labels is a pipe (|).
	Mode types.DocumentClassifierMode

	// ID for the KMS key that Amazon Comprehend uses to encrypt trained custom
	// models. The ModelKmsKeyId can be either of the following formats:
	//   - KMS Key ID: "1234abcd-12ab-34cd-56ef-1234567890ab"
	//   - Amazon Resource Name (ARN) of a KMS Key:
	//   "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	ModelKmsKeyId *string

	// The resource-based policy to attach to your custom document classifier model.
	// You can use this policy to allow another Amazon Web Services account to import
	// your custom model. Provide your policy as a JSON body that you enter as a UTF-8
	// encoded string without line breaks. To provide valid JSON, enclose the attribute
	// names and values in double quotes. If the JSON body is also enclosed in double
	// quotes, then you must escape the double quotes that are inside the policy:
	// "{\"attribute\": \"value\", \"attribute\": [\"value\"]}" To avoid escaping
	// quotes, you can use single quotes to enclose the policy and double quotes to
	// enclose the JSON names and values: '{"attribute": "value", "attribute":
	// ["value"]}'
	ModelPolicy *string

	// Enables the addition of output results configuration parameters for custom
	// classifier jobs.
	OutputDataConfig *types.DocumentClassifierOutputDataConfig

	// Tags to associate with the document classifier. A tag is a key-value pair that
	// adds as a metadata to a resource used by Amazon Comprehend. For example, a tag
	// with "Sales" as the key might be added to a resource to indicate its use by the
	// sales department.
	Tags []types.Tag

	// The version name given to the newly created classifier. Version names can have
	// a maximum of 256 characters. Alphanumeric characters, hyphens (-) and
	// underscores (_) are allowed. The version name must be unique among all models
	// with the same classifier name in the Amazon Web Services account/Amazon Web
	// Services Region.
	VersionName *string

	// ID for the Amazon Web Services Key Management Service (KMS) key that Amazon
	// Comprehend uses to encrypt data on the storage volume attached to the ML compute
	// instance(s) that process the analysis job. The VolumeKmsKeyId can be either of
	// the following formats:
	//   - KMS Key ID: "1234abcd-12ab-34cd-56ef-1234567890ab"
	//   - Amazon Resource Name (ARN) of a KMS Key:
	//   "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	VolumeKmsKeyId *string

	// Configuration parameters for an optional private Virtual Private Cloud (VPC)
	// containing the resources you are using for your custom classifier. For more
	// information, see Amazon VPC (https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html)
	// .
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

type CreateDocumentClassifierOutput struct {

	// The Amazon Resource Name (ARN) that identifies the document classifier.
	DocumentClassifierArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDocumentClassifierMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDocumentClassifier{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDocumentClassifier{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateDocumentClassifierMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDocumentClassifierValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDocumentClassifier(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateDocumentClassifier struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDocumentClassifier) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDocumentClassifier) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDocumentClassifierInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDocumentClassifierInput ")
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
func addIdempotencyToken_opCreateDocumentClassifierMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateDocumentClassifier{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDocumentClassifier(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "CreateDocumentClassifier",
	}
}
