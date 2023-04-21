// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts a reference import job.
func (c *Client) StartReferenceImportJob(ctx context.Context, params *StartReferenceImportJobInput, optFns ...func(*Options)) (*StartReferenceImportJobOutput, error) {
	if params == nil {
		params = &StartReferenceImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartReferenceImportJob", params, optFns, c.addOperationStartReferenceImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartReferenceImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartReferenceImportJobInput struct {

	// The job's reference store ID.
	//
	// This member is required.
	ReferenceStoreId *string

	// A service role for the job.
	//
	// This member is required.
	RoleArn *string

	// The job's source files.
	//
	// This member is required.
	Sources []types.StartReferenceImportJobSourceItem

	// To ensure that jobs don't run multiple times, specify a unique token for each
	// job.
	ClientToken *string

	noSmithyDocumentSerde
}

type StartReferenceImportJobOutput struct {

	// When the job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The job's ID.
	//
	// This member is required.
	Id *string

	// The job's reference store ID.
	//
	// This member is required.
	ReferenceStoreId *string

	// The job's service role ARN.
	//
	// This member is required.
	RoleArn *string

	// The job's status.
	//
	// This member is required.
	Status types.ReferenceImportJobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartReferenceImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartReferenceImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartReferenceImportJob{}, middleware.After)
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
	if err = addEndpointPrefix_opStartReferenceImportJobMiddleware(stack); err != nil {
		return err
	}
	if err = addOpStartReferenceImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartReferenceImportJob(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opStartReferenceImportJobMiddleware struct {
}

func (*endpointPrefix_opStartReferenceImportJobMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opStartReferenceImportJobMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "control-storage-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opStartReferenceImportJobMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opStartReferenceImportJobMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opStartReferenceImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "omics",
		OperationName: "StartReferenceImportJob",
	}
}
