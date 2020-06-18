package services

import (
	"restapi/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UserService interface defines user related API
type UserService interface {
	ListUser(*gin.Context) ([]models.User, error)
	GetUser(*gin.Context, int) (*models.User, error)
}

// UserServiceImpl is implementing API
type UserServiceImpl struct {
}

// NewUserService returns UserServiceImpl instance
func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{}
}

// ListUser returns the result from DB
func (s *UserServiceImpl) ListUser(c *gin.Context) ([]models.User, error) {
	db := c.MustGet("db").(*gorm.DB)
	users := []models.User{}
	err := db.Find(&users).Error
	return users, err
}

// GetUser returns the result from DB
func (s *UserServiceImpl) GetUser(c *gin.Context, userID int) (*models.User, error) {
	db := c.MustGet("db").(*gorm.DB)
	user := &models.User{}
	err := db.First(&user, userID).Error
	return user, err
}
