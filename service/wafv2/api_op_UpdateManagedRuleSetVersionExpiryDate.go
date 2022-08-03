// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the expiration information for your managed rule set. Use this to
// initiate the expiration of a managed rule group version. After you initiate
// expiration for a version, WAF excludes it from the response to
// ListAvailableManagedRuleGroupVersions for the managed rule group. This is
// intended for use only by vendors of managed rule sets. Vendors are Amazon Web
// Services and Amazon Web Services Marketplace sellers. Vendors, you can use the
// managed rule set APIs to provide controlled rollout of your versioned managed
// rule group offerings for your customers. The APIs are ListManagedRuleSets,
// GetManagedRuleSet, PutManagedRuleSetVersions, and
// UpdateManagedRuleSetVersionExpiryDate.
func (c *Client) UpdateManagedRuleSetVersionExpiryDate(ctx context.Context, params *UpdateManagedRuleSetVersionExpiryDateInput, optFns ...func(*Options)) (*UpdateManagedRuleSetVersionExpiryDateOutput, error) {
	if params == nil {
		params = &UpdateManagedRuleSetVersionExpiryDateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateManagedRuleSetVersionExpiryDate", params, optFns, c.addOperationUpdateManagedRuleSetVersionExpiryDateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateManagedRuleSetVersionExpiryDateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateManagedRuleSetVersionExpiryDateInput struct {

	// The time that you want the version to expire. Times are in Coordinated Universal
	// Time (UTC) format. UTC format includes the special designator, Z. For example,
	// "2016-09-27T14:50Z".
	//
	// This member is required.
	ExpiryTimestamp *time.Time

	// A unique identifier for the managed rule set. The ID is returned in the
	// responses to commands like list. You provide it to operations like get and
	// update.
	//
	// This member is required.
	Id *string

	// A token used for optimistic locking. WAF returns a token to your get and list
	// requests, to mark the state of the entity at the time of the request. To make
	// changes to the entity associated with the token, you provide the token to
	// operations like update and delete. WAF uses the token to ensure that no changes
	// have been made to the entity since you last retrieved it. If a change has been
	// made, the update fails with a WAFOptimisticLockException. If this happens,
	// perform another get, and use the new token returned by that operation.
	//
	// This member is required.
	LockToken *string

	// The name of the managed rule set. You use this, along with the rule set ID, to
	// identify the rule set. This name is assigned to the corresponding managed rule
	// group, which your customers can access and use.
	//
	// This member is required.
	Name *string

	// Specifies whether this is for an Amazon CloudFront distribution or for a
	// regional application. A regional application can be an Application Load Balancer
	// (ALB), an Amazon API Gateway REST API, an AppSync GraphQL API, or an Amazon
	// Cognito user pool. To work with CloudFront, you must also specify the Region US
	// East (N. Virginia) as follows:
	//
	// * CLI - Specify the Region when you use the
	// CloudFront scope: --scope=CLOUDFRONT --region=us-east-1.
	//
	// * API and SDKs - For
	// all calls, use the Region endpoint us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// The version that you want to remove from your list of offerings for the named
	// managed rule group.
	//
	// This member is required.
	VersionToExpire *string

	noSmithyDocumentSerde
}

type UpdateManagedRuleSetVersionExpiryDateOutput struct {

	// The version that is set to expire.
	ExpiringVersion *string

	// The time that the version will expire. Times are in Coordinated Universal Time
	// (UTC) format. UTC format includes the special designator, Z. For example,
	// "2016-09-27T14:50Z".
	ExpiryTimestamp *time.Time

	// A token used for optimistic locking. WAF returns a token to your get and list
	// requests, to mark the state of the entity at the time of the request. To make
	// changes to the entity associated with the token, you provide the token to
	// operations like update and delete. WAF uses the token to ensure that no changes
	// have been made to the entity since you last retrieved it. If a change has been
	// made, the update fails with a WAFOptimisticLockException. If this happens,
	// perform another get, and use the new token returned by that operation.
	NextLockToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateManagedRuleSetVersionExpiryDateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateManagedRuleSetVersionExpiryDate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateManagedRuleSetVersionExpiryDate{}, middleware.After)
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
	if err = addOpUpdateManagedRuleSetVersionExpiryDateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateManagedRuleSetVersionExpiryDate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateManagedRuleSetVersionExpiryDate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "UpdateManagedRuleSetVersionExpiryDate",
	}
}
