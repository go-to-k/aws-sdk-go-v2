// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Task runners call ReportTaskProgress when assigned a task to acknowledge that
// it has the task. If the web service does not receive this acknowledgement within
// 2 minutes, it assigns the task in a subsequent PollForTask call. After this
// initial acknowledgement, the task runner only needs to report progress every 15
// minutes to maintain its ownership of the task. You can change this reporting
// time from 15 minutes by specifying a reportProgressTimeout field in your
// pipeline. If a task runner does not report its status after 5 minutes, AWS Data
// Pipeline assumes that the task runner is unable to process the task and
// reassigns the task in a subsequent response to PollForTask . Task runners should
// call ReportTaskProgress every 60 seconds. POST / HTTP/1.1 Content-Type:
// application/x-amz-json-1.1 X-Amz-Target: DataPipeline.ReportTaskProgress
// Content-Length: 832 Host: datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon,
// 12 Nov 2012 17:49:52 GMT Authorization: AuthParams {"taskId":
// "aaGgHT4LuH0T0Y0oLrJRjas5qH0d8cDPADxqq3tn+zCWGELkCdV2JprLreXm1oxeP5EFZHFLJ69kjSsLYE0iYHYBYVGBrB+E/pYq7ANEEeGJFnSBMRiXZVA+8UJ3OzcInvXeinqBmBaKwii7hnnKb/AXjXiNTXyxgydX1KAyg1AxkwBYG4cfPYMZbuEbQJFJvv5C/2+GVXz1w94nKYTeUeepwUOFOuRLS6JVtZoYwpF56E+Yfk1IcGpFOvCZ01B4Bkuu7x3J+MD/j6kJgZLAgbCJQtI3eiW3kdGmX0p0I2BdY1ZsX6b4UiSvM3OMj6NEHJCJL4E0ZfitnhCoe24Kvjo6C2hFbZq+ei/HPgSXBQMSagkr4vS9c0ChzxH2+LNYvec6bY4kymkaZI1dvOzmpa0FcnGf5AjSK4GpsViZ/ujz6zxFv81qBXzjF0/4M1775rjV1VUdyKaixiA/sJiACNezqZqETidp8d24BDPRhGsj6pBCrnelqGFrk/gXEXUsJ+xwMifRC8UVwiKekpAvHUywVk7Ku4jH/n3i2VoLRP6FXwpUbelu34iiZ9czpXyLtyPKwxa87dlrnRVURwkcVjOt2Mcrcaqe+cbWHvNRhyrPkkdfSF3ac8/wfgVbXvLEB2k9mKc67aD9rvdc1PKX09Tk8BKklsMTpZ3TRCd4NzQlJKigMe8Jat9+1tKj4Ole5ZzW6uyTu2s2iFjEV8KXu4MaiRJyNKCdKeGhhZWY37Qk4NBK4Ppgu+C6Y41dpfOh288SLDEVx0/UySlqOEdhba7c6BiPp5r3hKj3mk9lFy5OYp1aoGLeeFmjXveTnPdf2gkWqXXg7AUbJ7jEs1F0lKZQg4szep2gcKyAJXgvXLfJJHcha8Lfb/Ee7wYmyOcAaRpDBoFNSbtoVXar46teIrpho+ZDvynUXvU0grHWGOk=:wn3SgymHZM99bEXAMPLE",
// "fields": [ {"key": "percentComplete", "stringValue": "50"} ] }
// x-amzn-RequestId: 640bd023-0775-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 18 Date: Mon, 12 Nov 2012 17:50:53
// GMT {"canceled": false}
func (c *Client) ReportTaskProgress(ctx context.Context, params *ReportTaskProgressInput, optFns ...func(*Options)) (*ReportTaskProgressOutput, error) {
	if params == nil {
		params = &ReportTaskProgressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReportTaskProgress", params, optFns, c.addOperationReportTaskProgressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReportTaskProgressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ReportTaskProgress.
type ReportTaskProgressInput struct {

	// The ID of the task assigned to the task runner. This value is provided in the
	// response for PollForTask .
	//
	// This member is required.
	TaskId *string

	// Key-value pairs that define the properties of the ReportTaskProgressInput
	// object.
	Fields []types.Field

	noSmithyDocumentSerde
}

// Contains the output of ReportTaskProgress.
type ReportTaskProgressOutput struct {

	// If true, the calling task runner should cancel processing of the task. The task
	// runner does not need to call SetTaskStatus for canceled tasks.
	//
	// This member is required.
	Canceled bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReportTaskProgressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpReportTaskProgress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpReportTaskProgress{}, middleware.After)
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
	if err = addOpReportTaskProgressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReportTaskProgress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReportTaskProgress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "ReportTaskProgress",
	}
}
