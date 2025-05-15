// Package logoinit provides method to Init loging config
package logoinit

import (
	"github.com/salamandaaa/cosmos-wallet/custodial/pkg/environment"
	"github.com/salamandaaa/cosmos-wallet/helpers/logo"
	"github.com/sirupsen/logrus"
)

func Init() {
	logrusEntry := logrus.New().WithFields(logrus.Fields{})

	if environment.GetEnvironment() == environment.PROD {
		logrusEntry.Logger.SetFormatter(&logrus.JSONFormatter{})
	}

	logo.SetInstance(*logrusEntry)
}
