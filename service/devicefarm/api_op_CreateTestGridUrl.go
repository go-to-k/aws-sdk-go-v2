// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a signed, short-term URL that can be passed to a Selenium
// RemoteWebDriver constructor.
func (c *Client) CreateTestGridUrl(ctx context.Context, params *CreateTestGridUrlInput, optFns ...func(*Options)) (*CreateTestGridUrlOutput, error) {
	if params == nil {
		params = &CreateTestGridUrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTestGridUrl", params, optFns, c.addOperationCreateTestGridUrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTestGridUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTestGridUrlInput struct {

	// Lifetime, in seconds, of the URL.
	//
	// This member is required.
	ExpiresInSeconds *int32

	// ARN (from CreateTestGridProject or ListTestGridProjects ) to associate with the
	// short-term URL.
	//
	// This member is required.
	ProjectArn *string

	noSmithyDocumentSerde
}

type CreateTestGridUrlOutput struct {

	// The number of seconds the URL from CreateTestGridUrlResult$url stays active.
	Expires *time.Time

	// A signed URL, expiring in CreateTestGridUrlRequest$expiresInSeconds seconds, to
	// be passed to a RemoteWebDriver .
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTestGridUrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTestGridUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTestGridUrl{}, middleware.After)
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
	if err = addOpCreateTestGridUrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTestGridUrl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTestGridUrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "CreateTestGridUrl",
	}
}
