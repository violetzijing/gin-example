package main

import (
	"fmt"
	"restapi/endpoint"
	"restapi/lib"
	"restapi/lib/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := lib.ParseConfig()

	db := database.Init(cfg.DBConfig)

	app := gin.Default()
	app.Use(database.Inject(db))
	app.Use(gin.Recovery())
	endpoint.ApplyRoutes(app)
	app.Run(fmt.Sprintf(":%d", cfg.Port))
}
