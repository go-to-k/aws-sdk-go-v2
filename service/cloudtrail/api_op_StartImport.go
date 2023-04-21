// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts an import of logged trail events from a source S3 bucket to a
// destination event data store. By default, CloudTrail only imports events
// contained in the S3 bucket's CloudTrail prefix and the prefixes inside the
// CloudTrail prefix, and does not check prefixes for other Amazon Web Services
// services. If you want to import CloudTrail events contained in another prefix,
// you must include the prefix in the S3LocationUri . For more considerations about
// importing trail events, see Considerations (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-copy-trail-to-lake.html#cloudtrail-trail-copy-considerations)
// . When you start a new import, the Destinations and ImportSource parameters are
// required. Before starting a new import, disable any access control lists (ACLs)
// attached to the source S3 bucket. For more information about disabling ACLs, see
// Controlling ownership of objects and disabling ACLs for your bucket (https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html)
// . When you retry an import, the ImportID parameter is required. If the
// destination event data store is for an organization, you must use the management
// account to import trail events. You cannot use the delegated administrator
// account for the organization.
func (c *Client) StartImport(ctx context.Context, params *StartImportInput, optFns ...func(*Options)) (*StartImportOutput, error) {
	if params == nil {
		params = &StartImportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartImport", params, optFns, c.addOperationStartImportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartImportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartImportInput struct {

	// The ARN of the destination event data store. Use this parameter for a new
	// import.
	Destinations []string

	// Use with StartEventTime to bound a StartImport request, and limit imported
	// trail events to only those events logged within a specified time period. When
	// you specify a time range, CloudTrail checks the prefix and log file names to
	// verify the names contain a date between the specified StartEventTime and
	// EndEventTime before attempting to import events.
	EndEventTime *time.Time

	// The ID of the import. Use this parameter when you are retrying an import.
	ImportId *string

	// The source S3 bucket for the import. Use this parameter for a new import.
	ImportSource *types.ImportSource

	// Use with EndEventTime to bound a StartImport request, and limit imported trail
	// events to only those events logged within a specified time period. When you
	// specify a time range, CloudTrail checks the prefix and log file names to verify
	// the names contain a date between the specified StartEventTime and EndEventTime
	// before attempting to import events.
	StartEventTime *time.Time

	noSmithyDocumentSerde
}

type StartImportOutput struct {

	// The timestamp for the import's creation.
	CreatedTimestamp *time.Time

	// The ARN of the destination event data store.
	Destinations []string

	// Used with StartEventTime to bound a StartImport request, and limit imported
	// trail events to only those events logged within a specified time period.
	EndEventTime *time.Time

	// The ID of the import.
	ImportId *string

	// The source S3 bucket for the import.
	ImportSource *types.ImportSource

	// Shows the status of the import after a StartImport request. An import finishes
	// with a status of COMPLETED if there were no failures, or FAILED if there were
	// failures.
	ImportStatus types.ImportStatus

	// Used with EndEventTime to bound a StartImport request, and limit imported trail
	// events to only those events logged within a specified time period.
	StartEventTime *time.Time

	// The timestamp of the import's last update, if applicable.
	UpdatedTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartImportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartImport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartImport{}, middleware.After)
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
	if err = addOpStartImportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartImport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartImport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "StartImport",
	}
}
