package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "todos-apps/src/config"
	"todos-apps/src/middleware"
	restcontroller "todos-apps/src/restcontroller"
	"todos-apps/src/router"
)

//https://gist.github.com/mashingan/4212d447f857cfdfbbba4f5436b779ac
func main() {
	apps := gin.New()
	godotenv.Load(".env")
	apps.LoadHTMLGlob("templates/**/*.tmpl") //https://hoohoo.top/blog/20210530112304-golang-tutorial-introduction-gin-html-template-and-how-integration-with-bootstrap/
	apps.Static("/assets", "./assets")


	web := apps.Group("/")
	web.Use(middleware.WebMiddleware)
	{
		router.WebRouter(web)
	}



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
