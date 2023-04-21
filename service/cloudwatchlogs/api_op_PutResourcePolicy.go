// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a resource policy allowing other Amazon Web Services
// services to put log events to this account, such as Amazon Route 53. An account
// can have up to 10 resource policies per Amazon Web Services Region.
func (c *Client) PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) {
	if params == nil {
		params = &PutResourcePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutResourcePolicy", params, optFns, c.addOperationPutResourcePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResourcePolicyInput struct {

	// Details of the new policy, including the identity of the principal that is
	// enabled to put logs to this account. This is formatted as a JSON string. This
	// parameter is required. The following example creates a resource policy enabling
	// the Route 53 service to put DNS query logs in to the specified log group.
	// Replace "logArn" with the ARN of your CloudWatch Logs resource, such as a log
	// group or log stream. CloudWatch Logs also supports aws:SourceArn (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html#condition-keys-sourcearn)
	// and aws:SourceAccount (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html#condition-keys-sourceaccount)
	// condition context keys. In the example resource policy, you would replace the
	// value of SourceArn with the resource making the call from Route 53 to
	// CloudWatch Logs. You would also replace the value of SourceAccount with the
	// Amazon Web Services account ID making that call. { "Version": "2012-10-17",
	// "Statement": [ { "Sid": "Route53LogsToCloudWatchLogs", "Effect": "Allow",
	// "Principal": { "Service": [ "route53.amazonaws.com" ] }, "Action":
	// "logs:PutLogEvents", "Resource": "logArn", "Condition": { "ArnLike": {
	// "aws:SourceArn": "myRoute53ResourceArn" }, "StringEquals": {
	// "aws:SourceAccount": "myAwsAccountId" } } } ] }
	PolicyDocument *string

	// Name of the new policy. This parameter is required.
	PolicyName *string

	noSmithyDocumentSerde
}

type PutResourcePolicyOutput struct {

	// The new policy.
	ResourcePolicy *types.ResourcePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutResourcePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutResourcePolicy{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutResourcePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutResourcePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "PutResourcePolicy",
	}
}
