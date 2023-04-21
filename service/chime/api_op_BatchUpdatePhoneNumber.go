// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates phone number product types or calling names. You can update one
// attribute at a time for each UpdatePhoneNumberRequestItem . For example, you can
// update the product type or the calling name. For toll-free numbers, you cannot
// use the Amazon Chime Business Calling product type. For numbers outside the
// U.S., you must use the Amazon Chime SIP Media Application Dial-In product type.
// Updates to outbound calling names can take up to 72 hours to complete. Pending
// updates to outbound calling names must be complete before you can request
// another update.
func (c *Client) BatchUpdatePhoneNumber(ctx context.Context, params *BatchUpdatePhoneNumberInput, optFns ...func(*Options)) (*BatchUpdatePhoneNumberOutput, error) {
	if params == nil {
		params = &BatchUpdatePhoneNumberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdatePhoneNumber", params, optFns, c.addOperationBatchUpdatePhoneNumberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdatePhoneNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdatePhoneNumberInput struct {

	// The request containing the phone number IDs and product types or calling names
	// to update.
	//
	// This member is required.
	UpdatePhoneNumberRequestItems []types.UpdatePhoneNumberRequestItem

	noSmithyDocumentSerde
}

type BatchUpdatePhoneNumberOutput struct {

	// If the action fails for one or more of the phone numbers in the request, a list
	// of the phone numbers is returned, along with error codes and error messages.
	PhoneNumberErrors []types.PhoneNumberError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchUpdatePhoneNumberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdatePhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdatePhoneNumber{}, middleware.After)
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
	if err = addOpBatchUpdatePhoneNumberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdatePhoneNumber(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchUpdatePhoneNumber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "BatchUpdatePhoneNumber",
	}
}
