// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The DisassociateQualificationFromWorker revokes a previously granted
// Qualification from a user. You can provide a text message explaining why the
// Qualification was revoked. The user who had the Qualification can see this
// message.
func (c *Client) DisassociateQualificationFromWorker(ctx context.Context, params *DisassociateQualificationFromWorkerInput, optFns ...func(*Options)) (*DisassociateQualificationFromWorkerOutput, error) {
	if params == nil {
		params = &DisassociateQualificationFromWorkerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateQualificationFromWorker", params, optFns, c.addOperationDisassociateQualificationFromWorkerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateQualificationFromWorkerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateQualificationFromWorkerInput struct {

	// The ID of the Qualification type of the Qualification to be revoked.
	//
	// This member is required.
	QualificationTypeId *string

	// The ID of the Worker who possesses the Qualification to be revoked.
	//
	// This member is required.
	WorkerId *string

	// A text message that explains why the Qualification was revoked. The user who
	// had the Qualification sees this message.
	Reason *string

	noSmithyDocumentSerde
}

type DisassociateQualificationFromWorkerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateQualificationFromWorkerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateQualificationFromWorker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateQualificationFromWorker{}, middleware.After)
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
	if err = addOpDisassociateQualificationFromWorkerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateQualificationFromWorker(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateQualificationFromWorker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "DisassociateQualificationFromWorker",
	}
}
