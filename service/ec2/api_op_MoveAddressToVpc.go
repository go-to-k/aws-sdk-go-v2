// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Moves an Elastic IP address from the EC2-Classic platform to the EC2-VPC
// platform. The Elastic IP address must be allocated to your account for more than
// 24 hours, and it must not be associated with an instance. After the Elastic IP
// address is moved, it is no longer available for use in the EC2-Classic platform,
// unless you move it back using the RestoreAddressToClassic request. You cannot
// move an Elastic IP address that was originally allocated for use in the EC2-VPC
// platform to the EC2-Classic platform. We are retiring EC2-Classic. We recommend
// that you migrate from EC2-Classic to a VPC. For more information, see Migrate
// from EC2-Classic to a VPC (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) MoveAddressToVpc(ctx context.Context, params *MoveAddressToVpcInput, optFns ...func(*Options)) (*MoveAddressToVpcOutput, error) {
	if params == nil {
		params = &MoveAddressToVpcInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MoveAddressToVpc", params, optFns, c.addOperationMoveAddressToVpcMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*MoveAddressToVpcOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MoveAddressToVpcInput struct {

	// The Elastic IP address.
	//
	// This member is required.
	PublicIp *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type MoveAddressToVpcOutput struct {

	// The allocation ID for the Elastic IP address.
	AllocationId *string

	// The status of the move of the IP address.
	Status types.Status

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationMoveAddressToVpcMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpMoveAddressToVpc{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpMoveAddressToVpc{}, middleware.After)
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
	if err = addOpMoveAddressToVpcValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opMoveAddressToVpc(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opMoveAddressToVpc(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "MoveAddressToVpc",
	}
}
