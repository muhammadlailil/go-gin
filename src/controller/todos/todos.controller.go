package todoscontroller

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	database "todos-apps/src/config"
	"todos-apps/src/entity"
	"todos-apps/src/helper"
	"todos-apps/src/repository"
)

func GetListTodo(c *gin.Context) {
	c.HTML(http.StatusOK, "todo.index.html", gin.H{
		"page_title": "Todos",
	})
}

func TodoListDataTable(c *gin.Context) {
	var (
		db            = database.InitDB(c)
		draw, _       = strconv.Atoi(c.PostForm("draw"))
		start, _      = strconv.Atoi(c.PostForm("start"))
		length, _     = strconv.Atoi(c.PostForm("length"))
		search        = c.PostForm("search[value]")
		orderColum, _ = strconv.Atoi(c.PostForm("order[0][column]"))
		orderDir      = c.PostForm("order[0][dir]")
		columns       = [](string){
			"id", "title", "completed",
		}
		sortColumns = columns[orderColum]
		todos       []entity.Todos
		paginate    helper.Pagination
	)
	page := (start / length) + 1
	paginate.Page = page
	paginate.Sort = sortColumns + " " + orderDir

	db.Scopes(
		helper.Paginate(todos, &paginate, db, length),
	).Where("title like ?", "%"+search+"%").Find(&todos)
	c.JSON(http.StatusOK, gin.H{
		"draw":            draw,
		"recordsTotal":    paginate.TotalData,
		"recordsFiltered": paginate.TotalData,
		"data":            todos,
	})
}

func PostSaveTodo(c *gin.Context) {
	db := database.InitDB(c)
	var (
		id           = c.PostForm("id")
		title        = c.PostForm("title")
		completed, _ = strconv.ParseBool(c.PostForm("completed"))
		message      = "Berhasil menambahkan data baru"
	)

	if id != "" {
		todos := repository.TodosFindById(db, id)
		todos.Title = title
		todos.Completed = completed
		db.Save(&todos)

		message = "Berhasil mengubah data yang dipilih"
	} else {
		db.Create(&entity.Todos{Title: title, Completed: completed})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": message,
	})
}

func PostDeleteTodo(c *gin.Context) {
	db := database.InitDB(c)
	var (
		id = c.PostForm("id")
	)
	todos := repository.TodosFindById(db, id)
	db.Delete(&todos)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil menghapus data yang dipilih",
	})
}

func PostImportTodo(c *gin.Context) {
	var (
		db       = database.InitDB(c)
		file, _  = c.FormFile("file")
		filename = "./assets/uploads/import_file." + filepath.Ext(file.Filename)
		err      = c.SaveUploadedFile(file, filename)
	)

	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	rows := f.GetRows("Sheet1")
	var todo []entity.Todos
	for i := 1; i < len(rows); i++ {
		var (
			row   = rows[i]
			title = row[0]
		)
		todo = append(todo, entity.Todos{Title: title, Completed: false})
	}
	db.Create(todo)
	c.Redirect(http.StatusFound, "/todo")
}

func GetExportTodo(c *gin.Context) {
	var (
		db          = database.InitDB(c)
		import_type = c.Query("type")
		todos       []entity.Todos
	)
	if import_type == "excel" {
		db.Order("id desc").Find(&todos)
		xlsx := excelize.NewFile()
		sheet1Name := "Sheet1"

		// create header
		xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)
		xlsx.SetCellValue(sheet1Name, "A1", "ID")
		xlsx.SetCellValue(sheet1Name, "B1", "Todo")
		xlsx.SetCellValue(sheet1Name, "C1", "Status")

		err := xlsx.AutoFilter(sheet1Name, "A1", "C1", "")
		if err != nil {
			log.Fatal("ERROR", err.Error())
		}
		for i, row := range todos {
			//status := strconv.FormatBool(row.Completed)
			//if !row.Completed {
			//	status = "Progress"
			//}
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), row.ID)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), row.Title)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), row.Completed)
		}
		fmt.Println("done loop")
		var w = c.Writer
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment;filename=todo_list.xlsx")
		w.Header().Set("File-Name", "todo_list.xlsx")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		w.Header().Set("Expires", "0")
		err = xlsx.Write(w)
	}
}
