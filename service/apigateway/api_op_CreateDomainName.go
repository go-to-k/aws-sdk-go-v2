// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new domain name.
func (c *Client) CreateDomainName(ctx context.Context, params *CreateDomainNameInput, optFns ...func(*Options)) (*CreateDomainNameOutput, error) {
	if params == nil {
		params = &CreateDomainNameInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDomainName", params, optFns, c.addOperationCreateDomainNameMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDomainNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to create a new domain name.
type CreateDomainNameInput struct {

	// The name of the DomainName resource.
	//
	// This member is required.
	DomainName *string

	// The reference to an AWS-managed certificate that will be used by edge-optimized
	// endpoint for this domain name. AWS Certificate Manager is the only supported
	// source.
	CertificateArn *string

	// [Deprecated] The body of the server certificate that will be used by
	// edge-optimized endpoint for this domain name provided by your certificate
	// authority.
	CertificateBody *string

	// [Deprecated] The intermediate certificates and optionally the root certificate,
	// one after the other without any blank lines, used by an edge-optimized endpoint
	// for this domain name. If you include the root certificate, your certificate
	// chain must start with intermediate certificates and end with the root
	// certificate. Use the intermediate certificates that were provided by your
	// certificate authority. Do not include any intermediaries that are not in the
	// chain of trust path.
	CertificateChain *string

	// The user-friendly name of the certificate that will be used by edge-optimized
	// endpoint for this domain name.
	CertificateName *string

	// [Deprecated] Your edge-optimized endpoint's domain name certificate's private
	// key.
	CertificatePrivateKey *string

	// The endpoint configuration of this DomainName showing the endpoint types of the
	// domain name.
	EndpointConfiguration *types.EndpointConfiguration

	// The mutual TLS authentication configuration for a custom domain name. If
	// specified, API Gateway performs two-way authentication between the client and
	// the server. Clients must present a trusted certificate to access your API.
	MutualTlsAuthentication *types.MutualTlsAuthenticationInput

	// The ARN of the public certificate issued by ACM to validate ownership of your
	// custom domain. Only required when configuring mutual TLS and using an ACM
	// imported or private CA certificate ARN as the regionalCertificateArn.
	OwnershipVerificationCertificateArn *string

	// The reference to an AWS-managed certificate that will be used by regional
	// endpoint for this domain name. AWS Certificate Manager is the only supported
	// source.
	RegionalCertificateArn *string

	// The user-friendly name of the certificate that will be used by regional
	// endpoint for this domain name.
	RegionalCertificateName *string

	// The Transport Layer Security (TLS) version + cipher suite for this DomainName.
	// The valid values are TLS_1_0 and TLS_1_2 .
	SecurityPolicy types.SecurityPolicy

	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/]. The
	// tag key can be up to 128 characters and must not start with aws: . The tag value
	// can be up to 256 characters.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Represents a custom domain name as a user-friendly host name of an API
// (RestApi).
type CreateDomainNameOutput struct {

	// The reference to an AWS-managed certificate that will be used by edge-optimized
	// endpoint for this domain name. AWS Certificate Manager is the only supported
	// source.
	CertificateArn *string

	// The name of the certificate that will be used by edge-optimized endpoint for
	// this domain name.
	CertificateName *string

	// The timestamp when the certificate that was used by edge-optimized endpoint for
	// this domain name was uploaded.
	CertificateUploadDate *time.Time

	// The domain name of the Amazon CloudFront distribution associated with this
	// custom domain name for an edge-optimized endpoint. You set up this association
	// when adding a DNS record pointing the custom domain name to this distribution
	// name. For more information about CloudFront distributions, see the Amazon
	// CloudFront documentation.
	DistributionDomainName *string

	// The region-agnostic Amazon Route 53 Hosted Zone ID of the edge-optimized
	// endpoint. The valid value is Z2FDTNDATAQYW2 for all the regions. For more
	// information, see Set up a Regional Custom Domain Name and AWS Regions and
	// Endpoints for API Gateway.
	DistributionHostedZoneId *string

	// The custom domain name as an API host name, for example, my-api.example.com .
	DomainName *string

	// The status of the DomainName migration. The valid values are AVAILABLE and
	// UPDATING . If the status is UPDATING , the domain cannot be modified further
	// until the existing operation is complete. If it is AVAILABLE , the domain can be
	// updated.
	DomainNameStatus types.DomainNameStatus

	// An optional text message containing detailed information about status of the
	// DomainName migration.
	DomainNameStatusMessage *string

	// The endpoint configuration of this DomainName showing the endpoint types of the
	// domain name.
	EndpointConfiguration *types.EndpointConfiguration

	// The mutual TLS authentication configuration for a custom domain name. If
	// specified, API Gateway performs two-way authentication between the client and
	// the server. Clients must present a trusted certificate to access your API.
	MutualTlsAuthentication *types.MutualTlsAuthentication

	// The ARN of the public certificate issued by ACM to validate ownership of your
	// custom domain. Only required when configuring mutual TLS and using an ACM
	// imported or private CA certificate ARN as the regionalCertificateArn.
	OwnershipVerificationCertificateArn *string

	// The reference to an AWS-managed certificate that will be used for validating
	// the regional domain name. AWS Certificate Manager is the only supported source.
	RegionalCertificateArn *string

	// The name of the certificate that will be used for validating the regional
	// domain name.
	RegionalCertificateName *string

	// The domain name associated with the regional endpoint for this custom domain
	// name. You set up this association by adding a DNS record that points the custom
	// domain name to this regional domain name. The regional domain name is returned
	// by API Gateway when you create a regional endpoint.
	RegionalDomainName *string

	// The region-specific Amazon Route 53 Hosted Zone ID of the regional endpoint.
	// For more information, see Set up a Regional Custom Domain Name and AWS Regions
	// and Endpoints for API Gateway.
	RegionalHostedZoneId *string

	// The Transport Layer Security (TLS) version + cipher suite for this DomainName.
	// The valid values are TLS_1_0 and TLS_1_2 .
	SecurityPolicy types.SecurityPolicy

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDomainNameMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDomainName{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDomainName{}, middleware.After)
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
	if err = addOpCreateDomainNameValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomainName(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateDomainName(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateDomainName",
	}
}
