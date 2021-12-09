package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	logincontroller "todos-apps/src/controller/login"
	todoscontroller "todos-apps/src/controller/todos"
)

func WebRouter(router *gin.RouterGroup) {
	router.GET("/logout", logincontroller.GetLogout)
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/todo")
	})

	todo := router.Group("/todo")
	{
		todo.GET("", todoscontroller.GetListTodo)
		todo.POST("save", todoscontroller.PostSaveTodo)
		todo.POST("delete", todoscontroller.PostDeleteTodo)
		todo.POST("import", todoscontroller.PostImportTodo)
		todo.GET("export", todoscontroller.GetExportTodo)
		todo.POST("/data-table", todoscontroller.TodoListDataTable)
	}
}
