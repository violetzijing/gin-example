package endpoint

import (
	"net/http"
	"restapi/lib"
	"restapi/lib/middlewares"
	"restapi/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserEndPoint defines end_point for user
type UserEndPoint struct {
	r   *gin.Engine
	svc services.UserService
}

// NewUserEndPoint returns UserEndPoint and inits routes
func NewUserEndPoint(r *gin.Engine, svc services.UserService) *UserEndPoint {
	e := &UserEndPoint{
		r:   r,
		svc: svc,
	}
	e.initRoutes()
	return e
}

func (e *UserEndPoint) initRoutes() {
	e.r.GET("/users", middlewares.Authorized, e.ListUser)
	e.r.GET("/user/:id", e.GetUser)
}

// ListUser returns listing user json response
func (e *UserEndPoint) ListUser(c *gin.Context) {
	users, err := e.svc.ListUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.NewInternalServiceErr(err))
		return
	}
	result := []map[string]interface{}{}
	for _, u := range users {
		result = append(result, u.Serialize())
	}

	c.JSON(http.StatusOK, result)
}

// GetUser returns getting user json response
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

	c.JSON(http.StatusOK, user.Serialize())
}
