package database

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func InitDB(c *gin.Context) *gorm.DB  {
	hostname := "127.0.0.1"
	port	   := "3306"
	username := "laililmahfud"
	password := "123456"
	database := "belajar"


	dsn := username + ":" + password + "@tcp("+hostname+":"+port+")/"+database+"?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})


	if err != nil {
		c.JSON(500, gin.H{"status":0,"message":"Can't connect to the database!"})
		c.Abort()
		return nil
	}

	return db
}
