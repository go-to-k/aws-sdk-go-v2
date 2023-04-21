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

// Creates a customization of a GatewayResponse of a specified response type and
// status code on the given RestApi.
func (c *Client) PutGatewayResponse(ctx context.Context, params *PutGatewayResponseInput, optFns ...func(*Options)) (*PutGatewayResponseOutput, error) {
	if params == nil {
		params = &PutGatewayResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutGatewayResponse", params, optFns, c.addOperationPutGatewayResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutGatewayResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a customization of a GatewayResponse of a specified response type and
// status code on the given RestApi.
type PutGatewayResponseInput struct {

	// The response type of the associated GatewayResponse
	//
	// This member is required.
	ResponseType types.GatewayResponseType

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// Response parameters (paths, query strings and headers) of the GatewayResponse
	// as a string-to-string map of key-value pairs.
	ResponseParameters map[string]string

	// Response templates of the GatewayResponse as a string-to-string map of
	// key-value pairs.
	ResponseTemplates map[string]string

	// The HTTP status code of the GatewayResponse.
	StatusCode *string

	noSmithyDocumentSerde
}

// A gateway response of a given response type and status code, with optional
// response parameters and mapping templates.
type PutGatewayResponseOutput struct {

	// A Boolean flag to indicate whether this GatewayResponse is the default gateway
	// response ( true ) or not ( false ). A default gateway response is one generated
	// by API Gateway without any customization by an API developer.
	DefaultResponse bool

	// Response parameters (paths, query strings and headers) of the GatewayResponse
	// as a string-to-string map of key-value pairs.
	ResponseParameters map[string]string

	// Response templates of the GatewayResponse as a string-to-string map of
	// key-value pairs.
	ResponseTemplates map[string]string

	// The response type of the associated GatewayResponse.
	ResponseType types.GatewayResponseType

	// The HTTP status code for this GatewayResponse.
	StatusCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutGatewayResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutGatewayResponse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutGatewayResponse{}, middleware.After)
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
	if err = addOpPutGatewayResponseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutGatewayResponse(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutGatewayResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "PutGatewayResponse",
	}
}
