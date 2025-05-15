package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/salamandaaa/cosmos-wallet/custodial/app/stage/appinit"
	"github.com/salamandaaa/cosmos-wallet/custodial/app/stage/apprun"
	"github.com/salamandaaa/cosmos-wallet/helpers/logo"
)

func main() {
	appinit.Init()
	logo.Info("Starting Custodial Wallet")
	apprun.Run()
}
