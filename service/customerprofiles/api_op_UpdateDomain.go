// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the properties of a domain, including creating or selecting a dead
// letter queue or an encryption key. After a domain is created, the name can’t be
// changed. Use this API or CreateDomain (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_CreateDomain.html)
// to enable identity resolution (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html)
// : set Matching to true. To prevent cross-service impersonation when you call
// this API, see Cross-service confused deputy prevention (https://docs.aws.amazon.com/connect/latest/adminguide/cross-service-confused-deputy-prevention.html)
// for sample policies that you should apply. To add or remove tags on an existing
// Domain, see TagResource (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_TagResource.html)
// / UntagResource (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_UntagResource.html)
// .
func (c *Client) UpdateDomain(ctx context.Context, params *UpdateDomainInput, optFns ...func(*Options)) (*UpdateDomainOutput, error) {
	if params == nil {
		params = &UpdateDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDomain", params, optFns, c.addOperationUpdateDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDomainInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The URL of the SQS dead letter queue, which is used for reporting errors
	// associated with ingesting data from third party applications. If specified as an
	// empty string, it will clear any existing value. You must set up a policy on the
	// DeadLetterQueue for the SendMessage operation to enable Amazon Connect Customer
	// Profiles to send messages to the DeadLetterQueue.
	DeadLetterQueueUrl *string

	// The default encryption key, which is an AWS managed key, is used when no
	// specific type of encryption key is specified. It is used to encrypt all data
	// before it is placed in permanent or semi-permanent storage. If specified as an
	// empty string, it will clear any existing value.
	DefaultEncryptionKey *string

	// The default number of days until the data within the domain expires.
	DefaultExpirationDays *int32

	// The process of matching duplicate profiles. If Matching = true , Amazon Connect
	// Customer Profiles starts a weekly batch process called Identity Resolution Job.
	// If you do not specify a date and time for Identity Resolution Job to run, by
	// default it runs every Saturday at 12AM UTC to detect duplicate profiles in your
	// domains. After the Identity Resolution Job completes, use the GetMatches (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html)
	// API to return and review the results. Or, if you have configured ExportingConfig
	// in the MatchingRequest , you can download the results from S3.
	Matching *types.MatchingRequest

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type UpdateDomainOutput struct {

	// The timestamp of when the domain was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The timestamp of when the domain was most recently edited.
	//
	// This member is required.
	LastUpdatedAt *time.Time

	// The URL of the SQS dead letter queue, which is used for reporting errors
	// associated with ingesting data from third party applications.
	DeadLetterQueueUrl *string

	// The default encryption key, which is an AWS managed key, is used when no
	// specific type of encryption key is specified. It is used to encrypt all data
	// before it is placed in permanent or semi-permanent storage.
	DefaultEncryptionKey *string

	// The default number of days until the data within the domain expires.
	DefaultExpirationDays *int32

	// The process of matching duplicate profiles. If Matching = true , Amazon Connect
	// Customer Profiles starts a weekly batch process called Identity Resolution Job.
	// If you do not specify a date and time for Identity Resolution Job to run, by
	// default it runs every Saturday at 12AM UTC to detect duplicate profiles in your
	// domains. After the Identity Resolution Job completes, use the GetMatches (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html)
	// API to return and review the results. Or, if you have configured ExportingConfig
	// in the MatchingRequest , you can download the results from S3.
	Matching *types.MatchingResponse

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDomain{}, middleware.After)
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
	if err = addOpUpdateDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "UpdateDomain",
	}
}
