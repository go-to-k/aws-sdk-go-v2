// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Inspects the text of a batch of documents for the syntax and part of speech of
// the words in the document and returns information about them. For more
// information, see Syntax (https://docs.aws.amazon.com/comprehend/latest/dg/how-syntax.html)
// in the Comprehend Developer Guide.
func (c *Client) BatchDetectSyntax(ctx context.Context, params *BatchDetectSyntaxInput, optFns ...func(*Options)) (*BatchDetectSyntaxOutput, error) {
	if params == nil {
		params = &BatchDetectSyntaxInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDetectSyntax", params, optFns, c.addOperationBatchDetectSyntaxMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDetectSyntaxOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDetectSyntaxInput struct {

	// The language of the input documents. You can specify any of the following
	// languages supported by Amazon Comprehend: German ("de"), English ("en"), Spanish
	// ("es"), French ("fr"), Italian ("it"), or Portuguese ("pt"). All documents must
	// be in the same language.
	//
	// This member is required.
	LanguageCode types.SyntaxLanguageCode

	// A list containing the UTF-8 encoded text of the input documents. The list can
	// contain a maximum of 25 documents. The maximum size for each document is 5 KB.
	//
	// This member is required.
	TextList []string

	noSmithyDocumentSerde
}

type BatchDetectSyntaxOutput struct {

	// A list containing one object for each document that contained an error. The
	// results are sorted in ascending order by the Index field and match the order of
	// the documents in the input list. If there are no errors in the batch, the
	// ErrorList is empty.
	//
	// This member is required.
	ErrorList []types.BatchItemError

	// A list of objects containing the results of the operation. The results are
	// sorted in ascending order by the Index field and match the order of the
	// documents in the input list. If all of the documents contain an error, the
	// ResultList is empty.
	//
	// This member is required.
	ResultList []types.BatchDetectSyntaxItemResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDetectSyntaxMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDetectSyntax{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDetectSyntax{}, middleware.After)
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
	if err = addOpBatchDetectSyntaxValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDetectSyntax(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDetectSyntax(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "BatchDetectSyntax",
	}
}
