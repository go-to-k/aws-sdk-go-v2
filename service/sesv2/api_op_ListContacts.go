// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the contacts present in a specific contact list.
func (c *Client) ListContacts(ctx context.Context, params *ListContactsInput, optFns ...func(*Options)) (*ListContactsOutput, error) {
	if params == nil {
		params = &ListContactsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListContacts", params, optFns, c.addOperationListContactsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListContactsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListContactsInput struct {

	// The name of the contact list.
	//
	// This member is required.
	ContactListName *string

	// A filter that can be applied to a list of contacts.
	Filter *types.ListContactsFilter

	// A string token indicating that there might be additional contacts available to
	// be listed. Use the token provided in the Response to use in the subsequent call
	// to ListContacts with the same parameters to retrieve the next page of contacts.
	NextToken *string

	// The number of contacts that may be returned at once, which is dependent on if
	// there are more or less contacts than the value of the PageSize. Use this
	// parameter to paginate results. If additional contacts exist beyond the specified
	// limit, the NextToken element is sent in the response. Use the NextToken value
	// in subsequent requests to retrieve additional contacts.
	PageSize *int32

	noSmithyDocumentSerde
}

type ListContactsOutput struct {

	// The contacts present in a specific contact list.
	Contacts []types.Contact

	// A string token indicating that there might be additional contacts available to
	// be listed. Copy this token to a subsequent call to ListContacts with the same
	// parameters to retrieve the next page of contacts.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListContactsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListContacts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListContacts{}, middleware.After)
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
	if err = addOpListContactsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListContacts(options.Region), middleware.Before); err != nil {
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

// ListContactsAPIClient is a client that implements the ListContacts operation.
type ListContactsAPIClient interface {
	ListContacts(context.Context, *ListContactsInput, ...func(*Options)) (*ListContactsOutput, error)
}

var _ ListContactsAPIClient = (*Client)(nil)

// ListContactsPaginatorOptions is the paginator options for ListContacts
type ListContactsPaginatorOptions struct {
	// The number of contacts that may be returned at once, which is dependent on if
	// there are more or less contacts than the value of the PageSize. Use this
	// parameter to paginate results. If additional contacts exist beyond the specified
	// limit, the NextToken element is sent in the response. Use the NextToken value
	// in subsequent requests to retrieve additional contacts.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListContactsPaginator is a paginator for ListContacts
type ListContactsPaginator struct {
	options   ListContactsPaginatorOptions
	client    ListContactsAPIClient
	params    *ListContactsInput
	nextToken *string
	firstPage bool
}

// NewListContactsPaginator returns a new ListContactsPaginator
func NewListContactsPaginator(client ListContactsAPIClient, params *ListContactsInput, optFns ...func(*ListContactsPaginatorOptions)) *ListContactsPaginator {
	if params == nil {
		params = &ListContactsInput{}
	}

	options := ListContactsPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListContactsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListContactsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListContacts page.
func (p *ListContactsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListContactsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListContacts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListContacts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListContacts",
	}
}
