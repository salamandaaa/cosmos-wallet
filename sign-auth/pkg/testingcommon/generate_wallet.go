package testingcommon

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/salamandaaa/cosmos-wallet/helpers/logo"
	walletaddress "github.com/salamandaaa/cosmos-wallet/sign-auth/pkg/cosmos_blockchain/wallet_address"
	"github.com/salamandaaa/cosmos-wallet/sign-auth/pkg/env"
)

type TestWallet struct {
	PrivKey       *secp256k1.PrivKey
	WalletAddress string
}

func GenerateWallet() TestWallet {
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()
	secp256k1PubKey := secp256k1.PubKey{
		Key: pubKey.Bytes(),
	}
	walletAddr, err := walletaddress.GetWalletAddrFromPubKey(env.MustGetEnv("WALLET_ADDRESS_HRP"), secp256k1PubKey)
	if err != nil {
		logo.Fatal(err)
	}
	return TestWallet{
		PrivKey:       privKey,
		WalletAddress: walletAddr,
	}
}
