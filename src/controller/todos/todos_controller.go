package todoscontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	database "todos-apps/src/config"
	"todos-apps/src/entity"
	todosrepository "todos-apps/src/repository/todos"
)

func ListTodo(c *gin.Context)  {
	db := database.InitDB(c)
	title := c.Query("title")
	complate := c.Query("complate")
	fmt.Println(complate)

	//var todos []entity.Todos
	//db.Order("id desc").Find(&todos)
	page,_:=strconv.Atoi(c.Query("page"))
	todos := todosrepository.FindLimitFilter(db,page,title,complate)

	c.JSON(http.StatusOK, gin.H{
		"status" : http.StatusOK,
		"data" : todos,
	})
}

func CreateTodo(c *gin.Context)  {
	db := database.InitDB(c)

	title:= c.PostForm("title")
	db.Create(&entity.Todos{Title: title, Completed: false})
	c.JSON(http.StatusOK, gin.H{
		"status" : http.StatusOK,
	})
}

func UpdateTodo(c *gin.Context)  {
	db := database.InitDB(c)

	id := c.PostForm("id")
	title:= c.PostForm("title")

	todos := todosrepository.FindById(db,id)
	todos.Title = title;
	db.Save(&todos)

	c.JSON(http.StatusOK, gin.H{
		"status" : http.StatusOK,
	})
}


func DeleteTodo(c *gin.Context)  {
	db := database.InitDB(c)

	id := c.Param("id")

	todos := todosrepository.FindById(db,id)
	db.Delete(&todos)

	c.JSON(http.StatusOK, gin.H{
		"status" : http.StatusOK,
	})
}
