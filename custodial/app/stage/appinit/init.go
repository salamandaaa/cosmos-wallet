// Package appinit provides method to Init all stages of app
package appinit

import (
	"github.com/salamandaaa/cosmos-wallet/custodial/app/stage/appinit/cosmosinit"
	"github.com/salamandaaa/cosmos-wallet/custodial/app/stage/appinit/dbconinit"
	"github.com/salamandaaa/cosmos-wallet/custodial/app/stage/appinit/dbmigrate"
	"github.com/salamandaaa/cosmos-wallet/custodial/app/stage/appinit/logoinit"
)

func Init() {
	cosmosinit.Init()
	logoinit.Init()
	dbconinit.Init()
	dbmigrate.Migrate()
}
