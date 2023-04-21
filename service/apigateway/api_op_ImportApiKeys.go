// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Import API keys from an external source, such as a CSV-formatted file.
func (c *Client) ImportApiKeys(ctx context.Context, params *ImportApiKeysInput, optFns ...func(*Options)) (*ImportApiKeysOutput, error) {
	if params == nil {
		params = &ImportApiKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportApiKeys", params, optFns, c.addOperationImportApiKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportApiKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The POST request to import API keys from an external source, such as a
// CSV-formatted file.
type ImportApiKeysInput struct {

	// The payload of the POST request to import API keys. For the payload format, see
	// API Key File Format.
	//
	// This member is required.
	Body []byte

	// A query parameter to specify the input format to imported API keys. Currently,
	// only the csv format is supported.
	//
	// This member is required.
	Format types.ApiKeysFormat

	// A query parameter to indicate whether to rollback ApiKey importation ( true ) or
	// not ( false ) when error is encountered.
	FailOnWarnings bool

	noSmithyDocumentSerde
}

// The identifier of an ApiKey used in a UsagePlan.
type ImportApiKeysOutput struct {

	// A list of all the ApiKey identifiers.
	Ids []string

	// A list of warning messages.
	Warnings []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportApiKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpImportApiKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpImportApiKeys{}, middleware.After)
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
	if err = addOpImportApiKeysValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportApiKeys(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opImportApiKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "ImportApiKeys",
	}
}
