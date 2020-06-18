package middlewares

import (
	"net/http"
	"restapi/lib"

	"github.com/gin-gonic/gin"
)

func Authorized(c *gin.Context) {
	_, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, lib.NewUnauthorized())
		c.AbortWithStatus(401)
		return
	}
}
