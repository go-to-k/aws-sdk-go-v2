// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Validates the syntax of a matchmaking rule or rule set. This operation checks
// that the rule set is using syntactically correct JSON and that it conforms to
// allowed property expressions. To validate syntax, provide a rule set JSON
// string. Learn more
//   - Build a rule set (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-rulesets.html)
func (c *Client) ValidateMatchmakingRuleSet(ctx context.Context, params *ValidateMatchmakingRuleSetInput, optFns ...func(*Options)) (*ValidateMatchmakingRuleSetOutput, error) {
	if params == nil {
		params = &ValidateMatchmakingRuleSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidateMatchmakingRuleSet", params, optFns, c.addOperationValidateMatchmakingRuleSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidateMatchmakingRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ValidateMatchmakingRuleSetInput struct {

	// A collection of matchmaking rules to validate, formatted as a JSON string.
	//
	// This member is required.
	RuleSetBody *string

	noSmithyDocumentSerde
}

type ValidateMatchmakingRuleSetOutput struct {

	// A response indicating whether the rule set is valid.
	Valid *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationValidateMatchmakingRuleSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpValidateMatchmakingRuleSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpValidateMatchmakingRuleSet{}, middleware.After)
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
	if err = addOpValidateMatchmakingRuleSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opValidateMatchmakingRuleSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opValidateMatchmakingRuleSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ValidateMatchmakingRuleSet",
	}
}
