// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies a cluster subnet group to include the specified list of VPC subnets.
// The operation replaces the existing list of subnets with the new list of
// subnets.
func (c *Client) ModifyClusterSubnetGroup(ctx context.Context, params *ModifyClusterSubnetGroupInput, optFns ...func(*Options)) (*ModifyClusterSubnetGroupOutput, error) {
	if params == nil {
		params = &ModifyClusterSubnetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyClusterSubnetGroup", params, optFns, c.addOperationModifyClusterSubnetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyClusterSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyClusterSubnetGroupInput struct {

	// The name of the subnet group to be modified.
	//
	// This member is required.
	ClusterSubnetGroupName *string

	// An array of VPC subnet IDs. A maximum of 20 subnets can be modified in a single
	// request.
	//
	// This member is required.
	SubnetIds []string

	// A text description of the subnet group to be modified.
	Description *string

	noSmithyDocumentSerde
}

type ModifyClusterSubnetGroupOutput struct {

	// Describes a subnet group.
	ClusterSubnetGroup *types.ClusterSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyClusterSubnetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyClusterSubnetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyClusterSubnetGroup{}, middleware.After)
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
	if err = addOpModifyClusterSubnetGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyClusterSubnetGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyClusterSubnetGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ModifyClusterSubnetGroup",
	}
}
