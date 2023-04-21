// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Assigns new properties to a user. Parameters you pass modify any or all of the
// following: the home directory, role, and policy for the UserName and ServerId
// you specify. The response returns the ServerId and the UserName for the updated
// user.
func (c *Client) UpdateUser(ctx context.Context, params *UpdateUserInput, optFns ...func(*Options)) (*UpdateUserOutput, error) {
	if params == nil {
		params = &UpdateUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUser", params, optFns, c.addOperationUpdateUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserInput struct {

	// A system-assigned unique identifier for a server instance that the user account
	// is assigned to.
	//
	// This member is required.
	ServerId *string

	// A unique string that identifies a user and is associated with a server as
	// specified by the ServerId . This user name must be a minimum of 3 and a maximum
	// of 100 characters long. The following are valid characters: a-z, A-Z, 0-9,
	// underscore '_', hyphen '-', period '.', and at sign '@'. The user name can't
	// start with a hyphen, period, or at sign.
	//
	// This member is required.
	UserName *string

	// The landing directory (folder) for a user when they log in to the server using
	// the client. A HomeDirectory example is /bucket_name/home/mydirectory .
	HomeDirectory *string

	// Logical directory mappings that specify what Amazon S3 or Amazon EFS paths and
	// keys should be visible to your user and how you want to make them visible. You
	// must specify the Entry and Target pair, where Entry shows how the path is made
	// visible and Target is the actual Amazon S3 or Amazon EFS path. If you only
	// specify a target, it is displayed as is. You also must ensure that your Identity
	// and Access Management (IAM) role provides access to paths in Target . This value
	// can be set only when HomeDirectoryType is set to LOGICAL. The following is an
	// Entry and Target pair example. [ { "Entry": "/directory1", "Target":
	// "/bucket_name/home/mydirectory" } ] In most cases, you can use this value
	// instead of the session policy to lock down your user to the designated home
	// directory (" chroot "). To do this, you can set Entry to '/' and set Target to
	// the HomeDirectory parameter value. The following is an Entry and Target pair
	// example for chroot . [ { "Entry": "/", "Target":
	// "/bucket_name/home/mydirectory" } ]
	HomeDirectoryMappings []types.HomeDirectoryMapEntry

	// The type of landing directory (folder) that you want your users' home directory
	// to be when they log in to the server. If you set it to PATH , the user will see
	// the absolute Amazon S3 bucket or EFS paths as is in their file transfer protocol
	// clients. If you set it LOGICAL , you need to provide mappings in the
	// HomeDirectoryMappings for how you want to make Amazon S3 or Amazon EFS paths
	// visible to your users.
	HomeDirectoryType types.HomeDirectoryType

	// A session policy for your user so that you can use the same Identity and Access
	// Management (IAM) role across multiple users. This policy scopes down a user's
	// access to portions of their Amazon S3 bucket. Variables that you can use inside
	// this policy include ${Transfer:UserName} , ${Transfer:HomeDirectory} , and
	// ${Transfer:HomeBucket} . This policy applies only when the domain of ServerId
	// is Amazon S3. Amazon EFS does not use session policies. For session policies,
	// Transfer Family stores the policy as a JSON blob, instead of the Amazon Resource
	// Name (ARN) of the policy. You save the policy as a JSON blob and pass it in the
	// Policy argument. For an example of a session policy, see Creating a session
	// policy (https://docs.aws.amazon.com/transfer/latest/userguide/session-policy) .
	// For more information, see AssumeRole (https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html)
	// in the Amazon Web Services Security Token Service API Reference.
	Policy *string

	// Specifies the full POSIX identity, including user ID ( Uid ), group ID ( Gid ),
	// and any secondary groups IDs ( SecondaryGids ), that controls your users' access
	// to your Amazon Elastic File Systems (Amazon EFS). The POSIX permissions that are
	// set on files and directories in your file system determines the level of access
	// your users get when transferring files into and out of your Amazon EFS file
	// systems.
	PosixProfile *types.PosixProfile

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// that controls your users' access to your Amazon S3 bucket or Amazon EFS file
	// system. The policies attached to this role determine the level of access that
	// you want to provide your users when transferring files into and out of your
	// Amazon S3 bucket or Amazon EFS file system. The IAM role should also contain a
	// trust relationship that allows the server to access your resources when
	// servicing your users' transfer requests.
	Role *string

	noSmithyDocumentSerde
}

// UpdateUserResponse returns the user name and identifier for the request to
// update a user's properties.
type UpdateUserOutput struct {

	// A system-assigned unique identifier for a server instance that the user account
	// is assigned to.
	//
	// This member is required.
	ServerId *string

	// The unique identifier for a user that is assigned to a server instance that was
	// specified in the request.
	//
	// This member is required.
	UserName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateUser{}, middleware.After)
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
	if err = addOpUpdateUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "UpdateUser",
	}
}
