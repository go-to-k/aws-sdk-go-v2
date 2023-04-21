// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Import a new custom lens or update an existing custom lens. To update an
// existing custom lens, specify its ARN as the LensAlias . If no ARN is specified,
// a new custom lens is created. The new or updated lens will have a status of
// DRAFT . The lens cannot be applied to workloads or shared with other Amazon Web
// Services accounts until it's published with CreateLensVersion . Lenses are
// defined in JSON. For more information, see JSON format specification (https://docs.aws.amazon.com/wellarchitected/latest/userguide/lenses-format-specification.html)
// in the Well-Architected Tool User Guide. A custom lens cannot exceed 500 KB in
// size. Disclaimer Do not include or gather personal identifiable information
// (PII) of end users or other identifiable individuals in or via your custom
// lenses. If your custom lens or those shared with you and used in your account do
// include or collect PII you are responsible for: ensuring that the included PII
// is processed in accordance with applicable law, providing adequate privacy
// notices, and obtaining necessary consents for processing such data.
func (c *Client) ImportLens(ctx context.Context, params *ImportLensInput, optFns ...func(*Options)) (*ImportLensOutput, error) {
	if params == nil {
		params = &ImportLensInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportLens", params, optFns, c.addOperationImportLensMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportLensOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportLensInput struct {

	// A unique case-sensitive string used to ensure that this request is idempotent
	// (executes only once). You should not reuse the same token for other requests. If
	// you retry a request with the same client request token and the same parameters
	// after the original request has completed successfully, the result of the
	// original request is returned. This token is listed as required, however, if you
	// do not specify it, the Amazon Web Services SDKs automatically generate one for
	// you. If you are not using the Amazon Web Services SDK or the CLI, you must
	// provide this token or the request will fail.
	//
	// This member is required.
	ClientRequestToken *string

	// The JSON representation of a lens.
	//
	// This member is required.
	JSONString *string

	// The alias of the lens. For Amazon Web Services official lenses, this is either
	// the lens alias, such as serverless , or the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1::lens/serverless . Note that some operations
	// (such as ExportLens and CreateLensShare) are not permitted on Amazon Web
	// Services official lenses. For custom lenses, this is the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2:123456789012:lens/0123456789abcdef01234567890abcdef
	// . Each lens is identified by its LensSummary$LensAlias .
	LensAlias *string

	// Tags to associate to a lens.
	Tags map[string]string

	noSmithyDocumentSerde
}

type ImportLensOutput struct {

	// The ARN for the lens that was created or updated.
	LensArn *string

	// The status of the imported lens.
	Status types.ImportLensStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportLensMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpImportLens{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpImportLens{}, middleware.After)
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
	if err = addIdempotencyToken_opImportLensMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpImportLensValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportLens(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpImportLens struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpImportLens) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpImportLens) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ImportLensInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ImportLensInput ")
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
func addIdempotencyToken_opImportLensMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpImportLens{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opImportLens(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "ImportLens",
	}
}
