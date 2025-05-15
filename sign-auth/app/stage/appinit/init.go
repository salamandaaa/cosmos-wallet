// Package appinit provides method to Init all stages of app
package appinit

import (
	"github.com/salamandaaa/cosmos-wallet/sign-auth/app/stage/appinit/dbconinit"
	"github.com/salamandaaa/cosmos-wallet/sign-auth/app/stage/appinit/dbmigrate"
	"github.com/salamandaaa/cosmos-wallet/sign-auth/app/stage/appinit/logoinit"
)

func Init() {
	logoinit.Init()
	dbconinit.Init()
	dbmigrate.Migrate()
}
