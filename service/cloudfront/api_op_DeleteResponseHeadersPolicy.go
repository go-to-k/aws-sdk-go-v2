// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a response headers policy. You cannot delete a response headers policy
// if it's attached to a cache behavior. First update your distributions to remove
// the response headers policy from all cache behaviors, then delete the response
// headers policy. To delete a response headers policy, you must provide the
// policy's identifier and version. To get these values, you can use
// ListResponseHeadersPolicies or GetResponseHeadersPolicy .
func (c *Client) DeleteResponseHeadersPolicy(ctx context.Context, params *DeleteResponseHeadersPolicyInput, optFns ...func(*Options)) (*DeleteResponseHeadersPolicyOutput, error) {
	if params == nil {
		params = &DeleteResponseHeadersPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteResponseHeadersPolicy", params, optFns, c.addOperationDeleteResponseHeadersPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteResponseHeadersPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteResponseHeadersPolicyInput struct {

	// The identifier for the response headers policy that you are deleting. To get
	// the identifier, you can use ListResponseHeadersPolicies .
	//
	// This member is required.
	Id *string

	// The version of the response headers policy that you are deleting. The version
	// is the response headers policy's ETag value, which you can get using
	// ListResponseHeadersPolicies , GetResponseHeadersPolicy , or
	// GetResponseHeadersPolicyConfig .
	IfMatch *string

	noSmithyDocumentSerde
}

type DeleteResponseHeadersPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteResponseHeadersPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteResponseHeadersPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteResponseHeadersPolicy{}, middleware.After)
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
	if err = addOpDeleteResponseHeadersPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteResponseHeadersPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteResponseHeadersPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "DeleteResponseHeadersPolicy",
	}
}
