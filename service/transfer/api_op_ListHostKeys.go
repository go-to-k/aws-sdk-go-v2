// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of host keys for the server that's specified by the ServerId
// parameter.
func (c *Client) ListHostKeys(ctx context.Context, params *ListHostKeysInput, optFns ...func(*Options)) (*ListHostKeysOutput, error) {
	if params == nil {
		params = &ListHostKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHostKeys", params, optFns, c.addOperationListHostKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHostKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHostKeysInput struct {

	// The identifier of the server that contains the host keys that you want to view.
	//
	// This member is required.
	ServerId *string

	// The maximum number of host keys to return.
	MaxResults *int32

	// When there are additional results that were not returned, a NextToken parameter
	// is returned. You can use that value for a subsequent call to ListHostKeys to
	// continue listing results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListHostKeysOutput struct {

	// Returns an array, where each item contains the details of a host key.
	//
	// This member is required.
	HostKeys []types.ListedHostKey

	// Returns the server identifier that contains the listed host keys.
	//
	// This member is required.
	ServerId *string

	// Returns a token that you can use to call ListHostKeys again and receive
	// additional results, if there are any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHostKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListHostKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListHostKeys{}, middleware.After)
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
	if err = addOpListHostKeysValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHostKeys(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListHostKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "ListHostKeys",
	}
}
