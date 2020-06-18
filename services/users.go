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
	CreateUser(c *gin.Context, user *models.User) error
	IsExisted(c *gin.Context, username string) (bool, error)
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

func (s *UserServiceImpl) CreateUser(c *gin.Context, user *models.User) error {
	db := c.MustGet("db").(*gorm.DB)
	db.NewRecord(*user)
	return db.Create(user).Error
}

func (a *UserServiceImpl) IsExisted(c *gin.Context, username string) (bool, error) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.Where("name = ?", username).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}
