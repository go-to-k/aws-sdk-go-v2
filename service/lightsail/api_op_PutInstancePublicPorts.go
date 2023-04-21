// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Opens ports for a specific Amazon Lightsail instance, and specifies the IP
// addresses allowed to connect to the instance through the ports, and the
// protocol. This action also closes all currently open ports that are not included
// in the request. Include all of the ports and the protocols you want to open in
// your PutInstancePublicPorts request. Or use the OpenInstancePublicPorts action
// to open ports without closing currently open ports. The PutInstancePublicPorts
// action supports tag-based access control via resource tags applied to the
// resource identified by instanceName . For more information, see the Amazon
// Lightsail Developer Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-controlling-access-using-tags)
// .
func (c *Client) PutInstancePublicPorts(ctx context.Context, params *PutInstancePublicPortsInput, optFns ...func(*Options)) (*PutInstancePublicPortsOutput, error) {
	if params == nil {
		params = &PutInstancePublicPortsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutInstancePublicPorts", params, optFns, c.addOperationPutInstancePublicPortsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutInstancePublicPortsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutInstancePublicPortsInput struct {

	// The name of the instance for which to open ports.
	//
	// This member is required.
	InstanceName *string

	// An array of objects to describe the ports to open for the specified instance.
	//
	// This member is required.
	PortInfos []types.PortInfo

	noSmithyDocumentSerde
}

type PutInstancePublicPortsOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operation *types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutInstancePublicPortsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutInstancePublicPorts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutInstancePublicPorts{}, middleware.After)
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
	if err = addOpPutInstancePublicPortsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutInstancePublicPorts(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutInstancePublicPorts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "PutInstancePublicPorts",
	}
}
