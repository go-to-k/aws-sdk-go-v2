// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an analysis scheme. For more information, see Configuring Analysis
// Schemes (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-analysis-schemes.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DeleteAnalysisScheme(ctx context.Context, params *DeleteAnalysisSchemeInput, optFns ...func(*Options)) (*DeleteAnalysisSchemeOutput, error) {
	if params == nil {
		params = &DeleteAnalysisSchemeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAnalysisScheme", params, optFns, c.addOperationDeleteAnalysisSchemeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAnalysisSchemeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DeleteAnalysisScheme operation. Specifies
// the name of the domain you want to update and the analysis scheme you want to
// delete.
type DeleteAnalysisSchemeInput struct {

	// The name of the analysis scheme you want to delete.
	//
	// This member is required.
	AnalysisSchemeName *string

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start with a
	// letter or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

// The result of a DeleteAnalysisScheme request. Contains the status of the
// deleted analysis scheme.
type DeleteAnalysisSchemeOutput struct {

	// The status of the analysis scheme being deleted.
	//
	// This member is required.
	AnalysisScheme *types.AnalysisSchemeStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAnalysisSchemeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteAnalysisScheme{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteAnalysisScheme{}, middleware.After)
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
	if err = addOpDeleteAnalysisSchemeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAnalysisScheme(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAnalysisScheme(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DeleteAnalysisScheme",
	}
}
