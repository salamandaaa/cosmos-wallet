package arbitraryverify

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/ethereum/go-ethereum/crypto"
)

func toECDSAPubKey(pk secp256k1.PubKey) (*ecdsa.PublicKey, error) {
	pubBytes := pk.Key
	x, y := elliptic.Unmarshal(secp256k1.S256(), pubBytes)
	if x == nil || y == nil {
		return nil, errors.New("invalid public key")
	}
	return &ecdsa.PublicKey{Curve: secp256k1.S256(), X: x, Y: y}, nil
}

// VerifyArbitraryMsg verifies arbitrary Adr036 message by first
// composing it with the 0 values required and
// then verifying it against public key
func VerifyArbitraryMsg(signer string, msg string, signature []byte, publicKey secp256k1.PubKey) (bool, error) {
	composedArbitraryMsg, err := ComposeArbitraryMsg(signer, msg)
	if err != nil {
		return false, fmt.Errorf("failed to compose arbitrary msg: %w", err)
	}

	hash := crypto.Keccak256Hash(composedArbitraryMsg)

	ecdsaPubKey, err := toECDSAPubKey(publicKey)
	if err != nil {
		return false, err
	}

	r := new(big.Int).SetBytes(signature[0:32])
	s := new(big.Int).SetBytes(signature[32:64])

	verified := ecdsa.Verify(ecdsaPubKey, hash[:], r, s)

	return verified, nil
}
