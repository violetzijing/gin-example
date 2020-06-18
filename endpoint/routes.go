package endpoint

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	UserRoutes(r)
}
