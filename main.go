package main

import (
	"fmt"

	"restapi/endpoint"
	"restapi/lib"
	"restapi/lib/database"
	"restapi/lib/middlewares"
	"restapi/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := lib.ParseConfig()

	db := database.Init(cfg.DBConfig)

	app := gin.Default()

	app.Use(database.Inject(db))
	app.Use(gin.Recovery())
	app.Use(middlewares.JWTMiddleware())

	registerEndpoint(app)

	app.Run(fmt.Sprintf(":%d", cfg.Port))
}

func registerEndpoint(r *gin.Engine) {
	userSVC := services.NewUserService()
	authSVC := services.NewAuthService()
	endpoint.NewUserEndPoint(r, userSVC)
	endpoint.NewAuthEndPoint(r, authSVC, userSVC)
}
