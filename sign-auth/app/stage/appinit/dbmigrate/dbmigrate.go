// Package dbmigrate provides method to migrate models into database
package dbmigrate

import (
	"github.com/salamandaaa/cosmos-wallet/helpers/logo"
	"github.com/salamandaaa/cosmos-wallet/sign-auth/models/flowid"
	"github.com/salamandaaa/cosmos-wallet/sign-auth/models/user"
	"github.com/salamandaaa/cosmos-wallet/sign-auth/pkg/store"
)

func Migrate() {
	db := store.DB
	err := db.AutoMigrate(&user.User{}, &flowid.FlowId{})
	if err != nil {
		logo.Fatalf("failed to migrate user into database: %s", err)
	}
}
