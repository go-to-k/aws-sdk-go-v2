// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a thing to a thing group. Requires permission to access the
// AddThingToThingGroup (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) AddThingToThingGroup(ctx context.Context, params *AddThingToThingGroupInput, optFns ...func(*Options)) (*AddThingToThingGroupOutput, error) {
	if params == nil {
		params = &AddThingToThingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddThingToThingGroup", params, optFns, c.addOperationAddThingToThingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddThingToThingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddThingToThingGroupInput struct {

	// Override dynamic thing groups with static thing groups when 10-group limit is
	// reached. If a thing belongs to 10 thing groups, and one or more of those groups
	// are dynamic thing groups, adding a thing to a static group removes the thing
	// from the last dynamic group.
	OverrideDynamicGroups bool

	// The ARN of the thing to add to a group.
	ThingArn *string

	// The ARN of the group to which you are adding a thing.
	ThingGroupArn *string

	// The name of the group to which you are adding a thing.
	ThingGroupName *string

	// The name of the thing to add to a group.
	ThingName *string

	noSmithyDocumentSerde
}

type AddThingToThingGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddThingToThingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAddThingToThingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAddThingToThingGroup{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddThingToThingGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddThingToThingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "AddThingToThingGroup",
	}
}
