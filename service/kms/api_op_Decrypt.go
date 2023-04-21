// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Decrypts ciphertext that was encrypted by a KMS key using any of the following
// operations:
//   - Encrypt
//   - GenerateDataKey
//   - GenerateDataKeyPair
//   - GenerateDataKeyWithoutPlaintext
//   - GenerateDataKeyPairWithoutPlaintext
//
// You can use this operation to decrypt ciphertext that was encrypted under a
// symmetric encryption KMS key or an asymmetric encryption KMS key. When the KMS
// key is asymmetric, you must specify the KMS key and the encryption algorithm
// that was used to encrypt the ciphertext. For information about asymmetric KMS
// keys, see Asymmetric KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
// in the Key Management Service Developer Guide. The Decrypt operation also
// decrypts ciphertext that was encrypted outside of KMS by the public key in an
// KMS asymmetric KMS key. However, it cannot decrypt symmetric ciphertext produced
// by other libraries, such as the Amazon Web Services Encryption SDK (https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/)
// or Amazon S3 client-side encryption (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html)
// . These libraries return a ciphertext format that is incompatible with KMS. If
// the ciphertext was encrypted under a symmetric encryption KMS key, the KeyId
// parameter is optional. KMS can get this information from metadata that it adds
// to the symmetric ciphertext blob. This feature adds durability to your
// implementation by ensuring that authorized users can decrypt ciphertext decades
// after it was encrypted, even if they've lost track of the key ID. However,
// specifying the KMS key is always recommended as a best practice. When you use
// the KeyId parameter to specify a KMS key, KMS only uses the KMS key you
// specify. If the ciphertext was encrypted under a different KMS key, the Decrypt
// operation fails. This practice ensures that you use the KMS key that you intend.
// Whenever possible, use key policies to give users permission to call the Decrypt
// operation on a particular KMS key, instead of using &IAM; policies. Otherwise,
// you might create an &IAM; policy that gives the user Decrypt permission on all
// KMS keys. This user could decrypt ciphertext that was encrypted by KMS keys in
// other accounts if the key policy for the cross-account KMS key permits it. If
// you must use an IAM policy for Decrypt permissions, limit the user to
// particular KMS keys or particular trusted accounts. For details, see Best
// practices for IAM policies (https://docs.aws.amazon.com/kms/latest/developerguide/iam-policies.html#iam-policies-best-practices)
// in the Key Management Service Developer Guide. Applications in Amazon Web
// Services Nitro Enclaves can call this operation by using the Amazon Web
// Services Nitro Enclaves Development Kit (https://github.com/aws/aws-nitro-enclaves-sdk-c)
// . For information about the supporting parameters, see How Amazon Web Services
// Nitro Enclaves use KMS (https://docs.aws.amazon.com/kms/latest/developerguide/services-nitro-enclaves.html)
// in the Key Management Service Developer Guide. The KMS key that you use for this
// operation must be in a compatible key state. For details, see Key states of KMS
// keys (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in
// the Key Management Service Developer Guide. Cross-account use: Yes. If you use
// the KeyId parameter to identify a KMS key in a different Amazon Web Services
// account, specify the key ARN or the alias ARN of the KMS key. Required
// permissions: kms:Decrypt (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (key policy) Related operations:
//   - Encrypt
//   - GenerateDataKey
//   - GenerateDataKeyPair
//   - ReEncrypt
func (c *Client) Decrypt(ctx context.Context, params *DecryptInput, optFns ...func(*Options)) (*DecryptOutput, error) {
	if params == nil {
		params = &DecryptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Decrypt", params, optFns, c.addOperationDecryptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DecryptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DecryptInput struct {

	// Ciphertext to be decrypted. The blob includes metadata.
	//
	// This member is required.
	CiphertextBlob []byte

	// Specifies the encryption algorithm that will be used to decrypt the ciphertext.
	// Specify the same algorithm that was used to encrypt the data. If you specify a
	// different algorithm, the Decrypt operation fails. This parameter is required
	// only when the ciphertext was encrypted under an asymmetric KMS key. The default
	// value, SYMMETRIC_DEFAULT , represents the only supported algorithm that is valid
	// for symmetric encryption KMS keys.
	EncryptionAlgorithm types.EncryptionAlgorithmSpec

	// Specifies the encryption context to use when decrypting the data. An encryption
	// context is valid only for cryptographic operations (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
	// with a symmetric encryption KMS key. The standard asymmetric encryption
	// algorithms and HMAC algorithms that KMS uses do not support an encryption
	// context. An encryption context is a collection of non-secret key-value pairs
	// that represent additional authenticated data. When you use an encryption context
	// to encrypt data, you must specify the same (an exact case-sensitive match)
	// encryption context to decrypt the data. An encryption context is supported only
	// on operations with symmetric encryption KMS keys. On operations with symmetric
	// encryption KMS keys, an encryption context is optional, but it is strongly
	// recommended. For more information, see Encryption context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the Key Management Service Developer Guide.
	EncryptionContext map[string]string

	// A list of grant tokens. Use a grant token when your permission to call this
	// operation comes from a new grant that has not yet achieved eventual consistency.
	// For more information, see Grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token)
	// and Using a grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token)
	// in the Key Management Service Developer Guide.
	GrantTokens []string

	// Specifies the KMS key that KMS uses to decrypt the ciphertext. Enter a key ID
	// of the KMS key that was used to encrypt the ciphertext. If you identify a
	// different KMS key, the Decrypt operation throws an IncorrectKeyException . This
	// parameter is required only when the ciphertext was encrypted under an asymmetric
	// KMS key. If you used a symmetric encryption KMS key, KMS can get the KMS key
	// from metadata that it adds to the symmetric ciphertext blob. However, it is
	// always recommended as a best practice. This practice ensures that you use the
	// KMS key that you intend. To specify a KMS key, use its key ID, key ARN, alias
	// name, or alias ARN. When using an alias name, prefix it with "alias/" . To
	// specify a KMS key in a different Amazon Web Services account, you must use the
	// key ARN or alias ARN. For example:
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Alias name: alias/ExampleAlias
	//   - Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey . To
	// get the alias name and alias ARN, use ListAliases .
	KeyId *string

	noSmithyDocumentSerde
}

type DecryptOutput struct {

	// The encryption algorithm that was used to decrypt the ciphertext.
	EncryptionAlgorithm types.EncryptionAlgorithmSpec

	// The Amazon Resource Name ( key ARN (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN)
	// ) of the KMS key that was used to decrypt the ciphertext.
	KeyId *string

	// Decrypted plaintext data. When you use the HTTP API or the Amazon Web Services
	// CLI, the value is Base64-encoded. Otherwise, it is not Base64-encoded.
	Plaintext []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDecryptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDecrypt{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDecrypt{}, middleware.After)
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
	if err = addOpDecryptValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDecrypt(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDecrypt(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "Decrypt",
	}
}
