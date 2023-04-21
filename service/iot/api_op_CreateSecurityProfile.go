// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Device Defender security profile. Requires permission to access the
// CreateSecurityProfile (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) CreateSecurityProfile(ctx context.Context, params *CreateSecurityProfileInput, optFns ...func(*Options)) (*CreateSecurityProfileOutput, error) {
	if params == nil {
		params = &CreateSecurityProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSecurityProfile", params, optFns, c.addOperationCreateSecurityProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSecurityProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSecurityProfileInput struct {

	// The name you are giving to the security profile.
	//
	// This member is required.
	SecurityProfileName *string

	// Please use CreateSecurityProfileRequest$additionalMetricsToRetainV2 instead. A
	// list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors , but it is also retained for any
	// metric specified here. Can be used with custom metrics; cannot be used with
	// dimensions.
	//
	// Deprecated: Use additionalMetricsToRetainV2.
	AdditionalMetricsToRetain []string

	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors , but it is also retained for any
	// metric specified here. Can be used with custom metrics; cannot be used with
	// dimensions.
	AdditionalMetricsToRetainV2 []types.MetricToRetain

	// Specifies the destinations to which alerts are sent. (Alerts are always sent to
	// the console.) Alerts are generated when a device (thing) violates a behavior.
	AlertTargets map[string]types.AlertTarget

	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	Behaviors []types.Behavior

	// A description of the security profile.
	SecurityProfileDescription *string

	// Metadata that can be used to manage the security profile.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateSecurityProfileOutput struct {

	// The ARN of the security profile.
	SecurityProfileArn *string

	// The name you gave to the security profile.
	SecurityProfileName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSecurityProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSecurityProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSecurityProfile{}, middleware.After)
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
	if err = addOpCreateSecurityProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSecurityProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSecurityProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateSecurityProfile",
	}
}
