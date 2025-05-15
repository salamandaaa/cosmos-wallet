// Package apprun provides method to Start http server of gin
package apprun

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/salamandaaa/cosmos-wallet/custodial/api"
	"github.com/salamandaaa/cosmos-wallet/custodial/pkg/env"
	"github.com/salamandaaa/cosmos-wallet/helpers/logo"
)

func Run() {
	ginApp := gin.Default()

	// Setup cors
	corsM := cors.New(cors.Config{AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowOrigins:     []string{env.MustGetEnv("ALLOWED_ORIGIN")}})
	ginApp.Use(corsM)

	// Apply /api
	api.ApplyRoutes(ginApp)
	port := env.MustGetEnv("APP_PORT")

	//Serve on APP_PORT
	err := ginApp.Run(":" + port)
	if err != nil {
		logo.Fatalf("failed to serve app on port %s: %s", port, err)
	}
}
