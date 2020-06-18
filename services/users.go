package services

import (
	"restapi/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserService interface {
	ListUser(*gin.Context) ([]models.User, error)
	GetUser(*gin.Context, int) (*models.User, error)
}

type UserServiceImpl struct {
}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (s *UserServiceImpl) ListUser(c *gin.Context) ([]models.User, error) {
	db := c.MustGet("db").(*gorm.DB)
	users := []models.User{}
	err := db.Find(&users).Error
	return users, err
}

func (s *UserServiceImpl) GetUser(c *gin.Context, userID int) (*models.User, error) {
	db := c.MustGet("db").(*gorm.DB)
	user := &models.User{}
	err := db.First(&user, userID).Error
	return user, err
}
