// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes one or more sets of featured results. Features results are placed above
// all other results for certain queries. If there's an exact match of a query,
// then one or more specific documents are featured in the search results.
func (c *Client) BatchDeleteFeaturedResultsSet(ctx context.Context, params *BatchDeleteFeaturedResultsSetInput, optFns ...func(*Options)) (*BatchDeleteFeaturedResultsSetOutput, error) {
	if params == nil {
		params = &BatchDeleteFeaturedResultsSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteFeaturedResultsSet", params, optFns, c.addOperationBatchDeleteFeaturedResultsSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteFeaturedResultsSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteFeaturedResultsSetInput struct {

	// The identifiers of the featured results sets that you want to delete.
	//
	// This member is required.
	FeaturedResultsSetIds []string

	// The identifier of the index used for featuring results.
	//
	// This member is required.
	IndexId *string

	noSmithyDocumentSerde
}

type BatchDeleteFeaturedResultsSetOutput struct {

	// The list of errors for the featured results set IDs, explaining why they
	// couldn't be removed from the index.
	//
	// This member is required.
	Errors []types.BatchDeleteFeaturedResultsSetError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDeleteFeaturedResultsSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDeleteFeaturedResultsSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDeleteFeaturedResultsSet{}, middleware.After)
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
	if err = addOpBatchDeleteFeaturedResultsSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteFeaturedResultsSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDeleteFeaturedResultsSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "BatchDeleteFeaturedResultsSet",
	}
}
