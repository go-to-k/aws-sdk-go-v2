// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about one or more Amazon Lightsail buckets. The information
// returned includes the synchronization status of the Amazon Simple Storage
// Service (Amazon S3) account-level block public access feature for your Lightsail
// buckets. For more information about buckets, see Buckets in Amazon Lightsail (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/buckets-in-amazon-lightsail)
// in the Amazon Lightsail Developer Guide.
func (c *Client) GetBuckets(ctx context.Context, params *GetBucketsInput, optFns ...func(*Options)) (*GetBucketsOutput, error) {
	if params == nil {
		params = &GetBucketsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBuckets", params, optFns, c.addOperationGetBucketsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketsInput struct {

	// The name of the bucket for which to return information. When omitted, the
	// response includes all of your buckets in the Amazon Web Services Region where
	// the request is made.
	BucketName *string

	// A Boolean value that indicates whether to include Lightsail instances that were
	// given access to the bucket using the SetResourceAccessForBucket (https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_SetResourceAccessForBucket.html)
	// action.
	IncludeConnectedResources *bool

	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetBuckets request. If your results are
	// paginated, the response will return a next page token that you can specify as
	// the page token in a subsequent request.
	PageToken *string

	noSmithyDocumentSerde
}

type GetBucketsOutput struct {

	// An object that describes the synchronization status of the Amazon S3
	// account-level block public access feature for your Lightsail buckets. For more
	// information about this feature and how it affects Lightsail buckets, see Block
	// public access for buckets in Amazon Lightsail (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-block-public-access-for-buckets)
	// .
	AccountLevelBpaSync *types.AccountLevelBpaSync

	// An array of objects that describe buckets.
	Buckets []types.Bucket

	// The token to advance to the next page of results from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetBuckets request and specify the next page
	// token using the pageToken parameter.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetBuckets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetBuckets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBuckets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBuckets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetBuckets",
	}
}
