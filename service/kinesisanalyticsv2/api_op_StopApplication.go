// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops the application from processing data. You can stop an application only if
// it is in the running status, unless you set the Force parameter to true . You
// can use the DescribeApplication operation to find the application status.
// Kinesis Data Analytics takes a snapshot when the application is stopped, unless
// Force is set to true .
func (c *Client) StopApplication(ctx context.Context, params *StopApplicationInput, optFns ...func(*Options)) (*StopApplicationOutput, error) {
	if params == nil {
		params = &StopApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopApplication", params, optFns, c.addOperationStopApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopApplicationInput struct {

	// The name of the running application to stop.
	//
	// This member is required.
	ApplicationName *string

	// Set to true to force the application to stop. If you set Force to true , Kinesis
	// Data Analytics stops the application without taking a snapshot. Force-stopping
	// your application may lead to data loss or duplication. To prevent data loss or
	// duplicate processing of data during application restarts, we recommend you to
	// take frequent snapshots of your application. You can only force stop a
	// Flink-based Kinesis Data Analytics application. You can't force stop a SQL-based
	// Kinesis Data Analytics application. The application must be in the STARTING ,
	// UPDATING , STOPPING , AUTOSCALING , or RUNNING status.
	Force *bool

	noSmithyDocumentSerde
}

type StopApplicationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopApplication{}, middleware.After)
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
	if err = addOpStopApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "StopApplication",
	}
}
