// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html)
// . With the latest version, AWS WAF has a single set of endpoints for regional
// and global use. Inserts or deletes Predicate objects in a rule and updates the
// RateLimit in the rule. Each Predicate object identifies a predicate, such as a
// ByteMatchSet or an IPSet , that specifies the web requests that you want to
// block or count. The RateLimit specifies the number of requests every five
// minutes that triggers the rule. If you add more than one predicate to a
// RateBasedRule , a request must match all the predicates and exceed the RateLimit
// to be counted or blocked. For example, suppose you add the following to a
// RateBasedRule :
//   - An IPSet that matches the IP address 192.0.2.44/32
//   - A ByteMatchSet that matches BadBot in the User-Agent header
//
// Further, you specify a RateLimit of 1,000. You then add the RateBasedRule to a
// WebACL and specify that you want to block requests that satisfy the rule. For a
// request to be blocked, it must come from the IP address 192.0.2.44 and the
// User-Agent header in the request must contain the value BadBot . Further,
// requests that match these two conditions much be received at a rate of more than
// 1,000 every five minutes. If the rate drops below this limit, AWS WAF no longer
// blocks the requests. As a second example, suppose you want to limit requests to
// a particular page on your site. To do this, you could add the following to a
// RateBasedRule :
//   - A ByteMatchSet with FieldToMatch of URI
//   - A PositionalConstraint of STARTS_WITH
//   - A TargetString of login
//
// Further, you specify a RateLimit of 1,000. By adding this RateBasedRule to a
// WebACL , you could limit requests to your login page without affecting the rest
// of your site.
func (c *Client) UpdateRateBasedRule(ctx context.Context, params *UpdateRateBasedRuleInput, optFns ...func(*Options)) (*UpdateRateBasedRuleOutput, error) {
	if params == nil {
		params = &UpdateRateBasedRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRateBasedRule", params, optFns, c.addOperationUpdateRateBasedRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRateBasedRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRateBasedRuleInput struct {

	// The value returned by the most recent call to GetChangeToken .
	//
	// This member is required.
	ChangeToken *string

	// The maximum number of requests, which have an identical value in the field
	// specified by the RateKey , allowed in a five-minute period. If the number of
	// requests exceeds the RateLimit and the other predicates specified in the rule
	// are also met, AWS WAF triggers the action that is specified for this rule.
	//
	// This member is required.
	RateLimit int64

	// The RuleId of the RateBasedRule that you want to update. RuleId is returned by
	// CreateRateBasedRule and by ListRateBasedRules .
	//
	// This member is required.
	RuleId *string

	// An array of RuleUpdate objects that you want to insert into or delete from a
	// RateBasedRule .
	//
	// This member is required.
	Updates []types.RuleUpdate

	noSmithyDocumentSerde
}

type UpdateRateBasedRuleOutput struct {

	// The ChangeToken that you used to submit the UpdateRateBasedRule request. You
	// can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus .
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRateBasedRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRateBasedRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRateBasedRule{}, middleware.After)
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
	if err = addOpUpdateRateBasedRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRateBasedRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRateBasedRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "UpdateRateBasedRule",
	}
}
