// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Finalizes the cutover immediately for specific Source Servers. All AWS
// resources created by Application Migration Service for enabling the replication
// of these source servers will be terminated / deleted within 90 minutes. Launched
// Test or Cutover instances will NOT be terminated. The AWS Replication Agent will
// receive a command to uninstall itself (within 10 minutes). The following
// properties of the SourceServer will be changed immediately:
// dataReplicationInfo.dataReplicationState will be changed to DISCONNECTED; The
// SourceServer.lifeCycle.state will be changed to CUTOVER; The totalStorageBytes
// property fo each of dataReplicationInfo.replicatedDisks will be set to zero;
// dataReplicationInfo.lagDuration and dataReplicationInfo.lagDuration will be
// nullified.
func (c *Client) FinalizeCutover(ctx context.Context, params *FinalizeCutoverInput, optFns ...func(*Options)) (*FinalizeCutoverOutput, error) {
	if params == nil {
		params = &FinalizeCutoverInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "FinalizeCutover", params, optFns, c.addOperationFinalizeCutoverMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*FinalizeCutoverOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type FinalizeCutoverInput struct {

	// Request to finalize Cutover by Source Server ID.
	//
	// This member is required.
	SourceServerID *string

	noSmithyDocumentSerde
}

type FinalizeCutoverOutput struct {

	// Source server application ID.
	ApplicationID *string

	// Source server ARN.
	Arn *string

	// Source server data replication info.
	DataReplicationInfo *types.DataReplicationInfo

	// Source server fqdn for action framework.
	FqdnForActionFramework *string

	// Source server archived status.
	IsArchived *bool

	// Source server launched instance.
	LaunchedInstance *types.LaunchedInstance

	// Source server lifecycle state.
	LifeCycle *types.LifeCycle

	// Source server replication type.
	ReplicationType types.ReplicationType

	// Source server properties.
	SourceProperties *types.SourceProperties

	// Source server ID.
	SourceServerID *string

	// Source server Tags.
	Tags map[string]string

	// Source server user provided ID.
	UserProvidedID *string

	// Source server vCenter client id.
	VcenterClientID *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationFinalizeCutoverMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpFinalizeCutover{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpFinalizeCutover{}, middleware.After)
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
	if err = addOpFinalizeCutoverValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opFinalizeCutover(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opFinalizeCutover(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "FinalizeCutover",
	}
}
