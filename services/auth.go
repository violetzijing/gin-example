package services

import (
	"io/ioutil"
	"restapi/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AuthService interface {
	Authorized(*gin.Context, string) (*models.User, error)
	GenerateToken(data map[string]interface{}) (string, error)
}

type AuthServiceImpl struct{}

func NewAuthService() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

func (a *AuthServiceImpl) Authorized(c *gin.Context, username string) (*models.User, error) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	err := db.Where("name = ?", username).First(&user).Error
	return &user, err
}

func (a *AuthServiceImpl) GenerateToken(data map[string]interface{}) (string, error) {
	//  token is valid for 7days
	date := time.Now().Add(time.Hour * 24 * 7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": data,
		"exp":  date.Unix(),
	})

	keyPath := "config/development/jwtsecret.key"

	key, readErr := ioutil.ReadFile(keyPath)
	if readErr != nil {
		return "", readErr
	}
	tokenString, err := token.SignedString(key)
	return tokenString, err
}
