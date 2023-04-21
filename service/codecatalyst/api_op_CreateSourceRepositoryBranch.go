// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a branch in a specified source repository in Amazon CodeCatalyst. This
// API only creates a branch in a source repository hosted in Amazon CodeCatalyst.
// You cannot use this API to create a branch in a linked repository.
func (c *Client) CreateSourceRepositoryBranch(ctx context.Context, params *CreateSourceRepositoryBranchInput, optFns ...func(*Options)) (*CreateSourceRepositoryBranchOutput, error) {
	if params == nil {
		params = &CreateSourceRepositoryBranchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSourceRepositoryBranch", params, optFns, c.addOperationCreateSourceRepositoryBranchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSourceRepositoryBranchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSourceRepositoryBranchInput struct {

	// The name for the branch you're creating.
	//
	// This member is required.
	Name *string

	// The name of the project in the space.
	//
	// This member is required.
	ProjectName *string

	// The name of the repository where you want to create a branch.
	//
	// This member is required.
	SourceRepositoryName *string

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	// The commit ID in an existing branch from which you want to create the new
	// branch.
	HeadCommitId *string

	noSmithyDocumentSerde
}

type CreateSourceRepositoryBranchOutput struct {

	// The commit ID of the tip of the newly created branch.
	HeadCommitId *string

	// The time the branch was last updated, in coordinated universal time (UTC)
	// timestamp format as specified in RFC 3339 (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
	// .
	LastUpdatedTime *time.Time

	// The name of the newly created branch.
	Name *string

	// The Git reference name of the branch.
	Ref *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSourceRepositoryBranchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSourceRepositoryBranch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSourceRepositoryBranch{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addBearerAuthSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateSourceRepositoryBranchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSourceRepositoryBranch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSourceRepositoryBranch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSourceRepositoryBranch",
	}
}
