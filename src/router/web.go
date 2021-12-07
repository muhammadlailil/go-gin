package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebRouter(router *gin.RouterGroup) {

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "todo/index.tmpl", gin.H{
			"title": "Main website",
		})
	})
}
