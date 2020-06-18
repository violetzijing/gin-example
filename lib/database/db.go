package database

import (
	"fmt"
	"restapi/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init(config string) *gorm.DB {
	db, err := gorm.Open("mysql", config)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to DB. %s", err.Error()))
	}
	db.LogMode(true)
	db.AutoMigrate(&models.User{})

	return db
}

func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
