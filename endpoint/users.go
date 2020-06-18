package endpoint

import (
	"net/http"
	"restapi/lib"
	"restapi/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserEndPoint struct {
	r   *gin.Engine
	svc services.UserService
}

func NewUserEndPoint(r *gin.Engine, svc services.UserService) *UserEndPoint {
	e := &UserEndPoint{
		r:   r,
		svc: svc,
	}
	e.InitRoutes()
	return e
}

func (e *UserEndPoint) InitRoutes() {
	e.r.GET("/users", e.ListUser)
	e.r.GET("/user/:id", e.GetUser)
}

func (e *UserEndPoint) ListUser(c *gin.Context) {
	users, err := e.svc.ListUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.NewInternalServiceErr(err))
		return
	}

	c.JSON(http.StatusOK, users)
}

func (e *UserEndPoint) GetUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.NewBadRequestErr("id"))
	}
	user, err := e.svc.GetUser(c, userID)
	if err != nil {
		if err.Error() != lib.NoRowFound {
			c.JSON(http.StatusInternalServerError, lib.NewInternalServiceErr(err))
			return
		}
		c.JSON(http.StatusNotFound, lib.NewNotFoundErr("user", userID))
		return
	}

	c.JSON(http.StatusOK, user)
}
