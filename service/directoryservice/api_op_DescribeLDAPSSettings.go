// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the status of LDAP security for the specified directory.
func (c *Client) DescribeLDAPSSettings(ctx context.Context, params *DescribeLDAPSSettingsInput, optFns ...func(*Options)) (*DescribeLDAPSSettingsOutput, error) {
	if params == nil {
		params = &DescribeLDAPSSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLDAPSSettings", params, optFns, c.addOperationDescribeLDAPSSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLDAPSSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLDAPSSettingsInput struct {

	// The identifier of the directory.
	//
	// This member is required.
	DirectoryId *string

	// Specifies the number of items that should be displayed on one page.
	Limit *int32

	// The type of next token used for pagination.
	NextToken *string

	// The type of LDAP security to enable. Currently only the value Client is
	// supported.
	Type types.LDAPSType

	noSmithyDocumentSerde
}

type DescribeLDAPSSettingsOutput struct {

	// Information about LDAP security for the specified directory, including status
	// of enablement, state last updated date time, and the reason for the state.
	LDAPSSettingsInfo []types.LDAPSSettingInfo

	// The next token used to retrieve the LDAPS settings if the number of setting
	// types exceeds page limit and there is another page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLDAPSSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLDAPSSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLDAPSSettings{}, middleware.After)
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
	if err = addOpDescribeLDAPSSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLDAPSSettings(options.Region), middleware.Before); err != nil {
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

// DescribeLDAPSSettingsAPIClient is a client that implements the
// DescribeLDAPSSettings operation.
type DescribeLDAPSSettingsAPIClient interface {
	DescribeLDAPSSettings(context.Context, *DescribeLDAPSSettingsInput, ...func(*Options)) (*DescribeLDAPSSettingsOutput, error)
}

var _ DescribeLDAPSSettingsAPIClient = (*Client)(nil)

// DescribeLDAPSSettingsPaginatorOptions is the paginator options for
// DescribeLDAPSSettings
type DescribeLDAPSSettingsPaginatorOptions struct {
	// Specifies the number of items that should be displayed on one page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeLDAPSSettingsPaginator is a paginator for DescribeLDAPSSettings
type DescribeLDAPSSettingsPaginator struct {
	options   DescribeLDAPSSettingsPaginatorOptions
	client    DescribeLDAPSSettingsAPIClient
	params    *DescribeLDAPSSettingsInput
	nextToken *string
	firstPage bool
}

// NewDescribeLDAPSSettingsPaginator returns a new DescribeLDAPSSettingsPaginator
func NewDescribeLDAPSSettingsPaginator(client DescribeLDAPSSettingsAPIClient, params *DescribeLDAPSSettingsInput, optFns ...func(*DescribeLDAPSSettingsPaginatorOptions)) *DescribeLDAPSSettingsPaginator {
	if params == nil {
		params = &DescribeLDAPSSettingsInput{}
	}

	options := DescribeLDAPSSettingsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeLDAPSSettingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeLDAPSSettingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeLDAPSSettings page.
func (p *DescribeLDAPSSettingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeLDAPSSettingsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeLDAPSSettings(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeLDAPSSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DescribeLDAPSSettings",
	}
}
