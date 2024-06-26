// Package crypto handles all cryptographic types and operations associated with Aptos.  It mainly handles signing,
// verification, parsing, and key generation.
package crypto

import "github.com/aptos-labs/aptos-go-sdk/bcs"

// Signer a generic interface for any kind of signing
type Signer interface {
	// Sign signs a transaction and returns an associated authenticator
	Sign(msg []byte) (authenticator *AccountAuthenticator, err error)

	// SignMessage signs a message and returns the raw signature without a key
	SignMessage(msg []byte) (signature Signature, err error)

	// AuthKey gives the AuthenticationKey associated with the signer
	AuthKey() *AuthenticationKey

	// PubKey Retrieve the public key for signature verification
	PubKey() PublicKey
}

// MessageSigner a generic interface for a signing private key, a private key isn't always a signer, see SingleSender
// This is not BCS serializable, because this doesn't go on-chain
type MessageSigner interface {
	// SignMessage signs a message and returns the raw signature without a key
	SignMessage(msg []byte) (signature Signature, err error)

	// VerifyingKey Retrieve the public key for signature verification
	VerifyingKey() VerifyingKey
}

// PublicKey is an interface for a public key that can be used to verify transactions in a TransactionAuthenticator
type PublicKey interface {
	VerifyingKey

	// AuthKey gives the AuthenticationKey associated with the public key
	AuthKey() *AuthenticationKey

	// Scheme The scheme used for address derivation
	Scheme() DeriveScheme
}

// VerifyingKey a generic interface for a public key associated with the private key, but it cannot necessarily stand on
// its own as a public key for authentication on Aptos
type VerifyingKey interface {
	bcs.Struct
	CryptoMaterial

	// Verify verifies a message with the public key
	Verify(msg []byte, sig Signature) bool
}

// Signature is an identifier for a serializable Signature for on-chain representation
type Signature interface {
	bcs.Struct
	CryptoMaterial
}

// CryptoMaterial is a set of functions for serializing and deserializing a key to and from bytes and hex
// This mirrors the trait in Rust
type CryptoMaterial interface {
	// Bytes outputs the raw byte representation of the CryptoMaterial
	Bytes() []byte

	// FromBytes loads the CryptoMaterial from the raw bytes
	FromBytes([]byte) error

	// ToHex outputs the hex representation of the CryptoMaterial with a leading `0x`
	ToHex() string

	// FromHex parses the hex representation of the CryptoMaterial with or without a leading `0x`
	FromHex(string) error
}
