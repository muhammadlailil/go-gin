package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	_ "todos-apps/src/config"
	logincontroller "todos-apps/src/controller/login"
	"todos-apps/src/helper"
	"todos-apps/src/middleware"
	restcontroller "todos-apps/src/restcontroller"
	"todos-apps/src/router"
)

//https://gist.github.com/mashingan/4212d447f857cfdfbbba4f5436b779ac
func main() {
	apps := gin.New()
	godotenv.Load(".env")
	store := cookie.NewStore([]byte("BWUIRY023##$$%6740KMXKCZM"))
	apps.Use(sessions.Sessions("mysession", store))

	apps.LoadHTMLGlob("templates/**/*.html") //https://hoohoo.top/blog/20210530112304-golang-tutorial-introduction-gin-html-template-and-how-integration-with-bootstrap/
	apps.Static("/assets", "./assets")

	apps.GET("/login", logincontroller.GetIndex)
	apps.POST("/login", logincontroller.PostLogin)
	web := apps.Group("/")
	web.Use(middleware.WebMiddleware)
	{
		router.WebRouter(web)
	}
	apps.GET("/mail", func(c *gin.Context) {
		helper.SendEmail()
		c.JSON(http.StatusOK, gin.H{
			"status": "oke",
		})
	})

	api := apps.Group("/api")
	api.GET("/get-token", restcontroller.GenerateToken)
	api.Use(middleware.ApiMiddleware)
	{
		router.ApiRouter(api)
	}

	err := apps.Run(":7000")
	if err != nil {
		return
	}
}
