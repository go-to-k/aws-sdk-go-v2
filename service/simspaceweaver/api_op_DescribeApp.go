// Code generated by smithy-go-codegen DO NOT EDIT.

package simspaceweaver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/simspaceweaver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the state of the given custom app.
func (c *Client) DescribeApp(ctx context.Context, params *DescribeAppInput, optFns ...func(*Options)) (*DescribeAppOutput, error) {
	if params == nil {
		params = &DescribeAppInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeApp", params, optFns, c.addOperationDescribeAppMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAppOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAppInput struct {

	// The name of the app.
	//
	// This member is required.
	App *string

	// The name of the domain of the app.
	//
	// This member is required.
	Domain *string

	// The name of the simulation of the app.
	//
	// This member is required.
	Simulation *string

	noSmithyDocumentSerde
}

type DescribeAppOutput struct {

	// The description of the app.
	Description *string

	// The name of the domain of the app.
	Domain *string

	// Information about the network endpoint for the custom app. You can use the
	// endpoint to connect to the custom app.
	EndpointInfo *types.SimulationAppEndpointInfo

	// Options that apply when the app starts. These options override default behavior.
	LaunchOverrides *types.LaunchOverrides

	// The name of the app.
	Name *string

	// The name of the simulation of the app.
	Simulation *string

	// The current lifecycle state of the custom app.
	Status types.SimulationAppStatus

	// The desired lifecycle state of the custom app.
	TargetStatus types.SimulationAppTargetStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAppMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeApp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeApp{}, middleware.After)
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
	if err = addOpDescribeAppValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApp(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeApp(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "simspaceweaver",
		OperationName: "DescribeApp",
	}
}
