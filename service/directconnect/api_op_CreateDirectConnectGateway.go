// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Direct Connect gateway, which is an intermediate object that enables
// you to connect a set of virtual interfaces and virtual private gateways. A
// Direct Connect gateway is global and visible in any Amazon Web Services Region
// after it is created. The virtual interfaces and virtual private gateways that
// are connected through a Direct Connect gateway can be in different Amazon Web
// Services Regions. This enables you to connect to a VPC in any Region, regardless
// of the Region in which the virtual interfaces are located, and pass traffic
// between them.
func (c *Client) CreateDirectConnectGateway(ctx context.Context, params *CreateDirectConnectGatewayInput, optFns ...func(*Options)) (*CreateDirectConnectGatewayOutput, error) {
	if params == nil {
		params = &CreateDirectConnectGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDirectConnectGateway", params, optFns, c.addOperationCreateDirectConnectGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDirectConnectGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDirectConnectGatewayInput struct {

	// The name of the Direct Connect gateway.
	//
	// This member is required.
	DirectConnectGatewayName *string

	// The autonomous system number (ASN) for Border Gateway Protocol (BGP) to be
	// configured on the Amazon side of the connection. The ASN must be in the private
	// range of 64,512 to 65,534 or 4,200,000,000 to 4,294,967,294. The default is
	// 64512.
	AmazonSideAsn *int64

	noSmithyDocumentSerde
}

type CreateDirectConnectGatewayOutput struct {

	// The Direct Connect gateway.
	DirectConnectGateway *types.DirectConnectGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDirectConnectGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDirectConnectGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDirectConnectGateway{}, middleware.After)
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
	if err = addOpCreateDirectConnectGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDirectConnectGateway(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDirectConnectGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "CreateDirectConnectGateway",
	}
}
