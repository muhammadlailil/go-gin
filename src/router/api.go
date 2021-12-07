package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	todoscontroller "todos-apps/src/controller/todos"
	"todos-apps/src/restcontroller"
)

func ApiRouter(router *gin.RouterGroup)  {
	router.POST("/login",restcontroller.PostLogin)
	router.GET("/todo",todoscontroller.ListTodo)
	router.POST("/todo/create",todoscontroller.CreateTodo)
	router.PATCH("/todo/update",todoscontroller.UpdateTodo)
	router.DELETE("/todo/delete/:id",todoscontroller.DeleteTodo)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status" : 1,
			"message": "Hello World !",
		})
	})
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action" // true
		c.String(http.StatusOK, "%t", b)
	})

}
