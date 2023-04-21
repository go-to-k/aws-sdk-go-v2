// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an Amazon FSx for OpenZFS volume to the state saved by the specified
// snapshot.
func (c *Client) RestoreVolumeFromSnapshot(ctx context.Context, params *RestoreVolumeFromSnapshotInput, optFns ...func(*Options)) (*RestoreVolumeFromSnapshotOutput, error) {
	if params == nil {
		params = &RestoreVolumeFromSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreVolumeFromSnapshot", params, optFns, c.addOperationRestoreVolumeFromSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreVolumeFromSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreVolumeFromSnapshotInput struct {

	// The ID of the source snapshot. Specifies the snapshot that you are restoring
	// from.
	//
	// This member is required.
	SnapshotId *string

	// The ID of the volume that you are restoring.
	//
	// This member is required.
	VolumeId *string

	// (Optional) An idempotency token for resource creation, in a string of up to 63
	// ASCII characters. This token is automatically filled on your behalf when you use
	// the Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	// The settings used when restoring the specified volume from snapshot.
	//   - DELETE_INTERMEDIATE_SNAPSHOTS - Deletes snapshots between the current state
	//   and the specified snapshot. If there are intermediate snapshots and this option
	//   isn't used, RestoreVolumeFromSnapshot fails.
	//   - DELETE_CLONED_VOLUMES - Deletes any dependent clone volumes created from
	//   intermediate snapshots. If there are any dependent clone volumes and this option
	//   isn't used, RestoreVolumeFromSnapshot fails.
	Options []types.RestoreOpenZFSVolumeOption

	noSmithyDocumentSerde
}

type RestoreVolumeFromSnapshotOutput struct {

	// A list of administrative actions for the file system that are in process or
	// waiting to be processed. Administrative actions describe changes to the Amazon
	// FSx system.
	AdministrativeActions []types.AdministrativeAction

	// The lifecycle state of the volume being restored.
	Lifecycle types.VolumeLifecycle

	// The ID of the volume that you restored.
	VolumeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreVolumeFromSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRestoreVolumeFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRestoreVolumeFromSnapshot{}, middleware.After)
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
	if err = addIdempotencyToken_opRestoreVolumeFromSnapshotMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRestoreVolumeFromSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreVolumeFromSnapshot(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpRestoreVolumeFromSnapshot struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRestoreVolumeFromSnapshot) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRestoreVolumeFromSnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RestoreVolumeFromSnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RestoreVolumeFromSnapshotInput ")
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
func addIdempotencyToken_opRestoreVolumeFromSnapshotMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpRestoreVolumeFromSnapshot{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opRestoreVolumeFromSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "RestoreVolumeFromSnapshot",
	}
}
