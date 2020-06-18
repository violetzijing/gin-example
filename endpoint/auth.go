package endpoint

import (
	"net/http"
	"restapi/lib"
	"restapi/models"
	"restapi/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// AuthEndPoint defines auth endpoint
type AuthEndPoint struct {
	r       *gin.Engine
	svc     services.AuthService
	userSVC services.UserService
}

// NewAuthEndPoint returns an instance of endpoint and inits routes
func NewAuthEndPoint(r *gin.Engine, svc services.AuthService, userSVC services.UserService) *AuthEndPoint {
	e := &AuthEndPoint{
		r:       r,
		svc:     svc,
		userSVC: userSVC,
	}
	e.initRoutes()
	return e
}

func (e *AuthEndPoint) initRoutes() {
	e.r.POST("/register", e.Register)
	e.r.POST("/login", e.Login)
}

type RequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register is for registering user and returning a token
func (e *AuthEndPoint) Register(c *gin.Context) {
	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, lib.NewBadRequestErr("body"))
		return
	}
	isExisted, err := e.userSVC.IsExisted(c, body.Username)
	if err != nil && err.Error() != lib.NoRowFound {
		c.JSON(http.StatusInternalServerError, lib.NewInternalServiceErr(err))
		return
	}
	if isExisted {
		c.JSON(http.StatusConflict, lib.NewConflict())
		return
	}
	hashedP, err := hash(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.NewInternalServiceErr(err))
		return
	}
	user := &models.User{
		Name:         body.Username,
		PasswordHash: hashedP,
	}
	if err := e.userSVC.CreateUser(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, lib.NewInternalServiceErr(err))
		return
	}

	token, err := e.svc.GenerateToken(user.Serialize())
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.NewInternalServiceErr(err))
		return
	}
	c.SetCookie("token", token, 60*60*24*7, "/", "", false, true)

	c.JSON(200, map[string]interface{}{
		"user":  user.Serialize(),
		"token": token,
	})
}

// Login will check the authentication and return a token
func (e *AuthEndPoint) Login(c *gin.Context) {
	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, lib.NewBadRequestErr("body"))
		return
	}
	user, err := e.svc.Authorized(c, body.Username)
	if err != nil {
		if err.Error() != lib.NoRowFound {
			c.JSON(http.StatusInternalServerError, lib.NewInternalServiceErr(err))
			return
		}
		c.JSON(http.StatusNotFound, lib.NewNotFoundErr("user", body.Username))
		return

	}
	if !checkHash(body.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, lib.NewUnauthorized())
		return
	}
	token, err := e.svc.GenerateToken(user.Serialize())
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.NewInternalServiceErr(err))
		return
	}
	c.SetCookie("token", token, 60*60*24*7, "/", "", false, true)
	c.JSON(200, map[string]interface{}{
		"user":  user.Serialize(),
		"token": token,
	})
}

func checkHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}
