package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/salamandaaa/cosmos-wallet/helpers/logo"
	"github.com/salamandaaa/cosmos-wallet/sign-auth/app/stage/appinit"
	"github.com/salamandaaa/cosmos-wallet/sign-auth/app/stage/apprun"
)

func main() {
	appinit.Init()
	logo.Info("Starting Sign Auth")
	apprun.Run()
}
