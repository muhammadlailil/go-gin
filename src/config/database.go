package database

import (
	"database/sql"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	hostname   = "127.0.0.1"
	port       = "3306"
	username   = "laililmahfud"
	password   = "123456"
	database   = "belajar"
	connection = username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + database
)

func InitDB(c *gin.Context) *gorm.DB {

	dsn := connection + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		c.JSON(500, gin.H{"status": 0, "message": "Can't connect to the database!"})
		c.Abort()
		return nil
	}

	return db
}

func NativeDB() *sql.DB {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}
