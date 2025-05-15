// Package cosmosinit provides method to Init cosmos config
package cosmosinit

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/salamandaaa/cosmos-wallet/custodial/pkg/env"
)

func Init() {
	config := sdkTypes.GetConfig()
	config.SetBech32PrefixForAccount(env.MustGetEnv("WALLET_ADDRESS_HRP"), env.MustGetEnv("PUBLIC_KEY_HRP"))
	config.Seal()
}
