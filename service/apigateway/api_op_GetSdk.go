// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates a client SDK for a RestApi and Stage.
func (c *Client) GetSdk(ctx context.Context, params *GetSdkInput, optFns ...func(*Options)) (*GetSdkOutput, error) {
	if params == nil {
		params = &GetSdkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSdk", params, optFns, c.addOperationGetSdkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSdkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request a new generated client SDK for a RestApi and Stage.
type GetSdkInput struct {

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// The language for the generated SDK. Currently java , javascript , android ,
	// objectivec (for iOS), swift (for iOS), and ruby are supported.
	//
	// This member is required.
	SdkType *string

	// The name of the Stage that the SDK will use.
	//
	// This member is required.
	StageName *string

	// A string-to-string key-value map of query parameters sdkType -dependent
	// properties of the SDK. For sdkType of objectivec or swift , a parameter named
	// classPrefix is required. For sdkType of android , parameters named groupId ,
	// artifactId , artifactVersion , and invokerPackage are required. For sdkType of
	// java , parameters named serviceName and javaPackageName are required.
	Parameters map[string]string

	noSmithyDocumentSerde
}

// The binary blob response to GetSdk, which contains the generated SDK.
type GetSdkOutput struct {

	// The binary blob response to GetSdk, which contains the generated SDK.
	Body []byte

	// The content-disposition header value in the HTTP response.
	ContentDisposition *string

	// The content-type header value in the HTTP response.
	ContentType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSdkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSdk{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSdk{}, middleware.After)
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
	if err = addOpGetSdkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSdk(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSdk(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetSdk",
	}
}
