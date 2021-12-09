package repository

import (
	"gorm.io/gorm"
	"todos-apps/src/entity"
	"todos-apps/src/helper"
)

func TodosFindById(db *gorm.DB, id string) entity.Todos {
	var todos entity.Todos
	db.Where("id=?", id).Find(&todos)
	return todos
}

func TodosFindLimitFilter(db *gorm.DB, page int, title string, complate string) helper.Pagination {
	var todos []entity.Todos
	var paginate helper.Pagination
	paginate.Page = page

	wherecondition := "title like '%" + title + "%'"
	if complate == "true" {
		wherecondition += " and completed=true"
	} else if complate == "false" {
		wherecondition += " and completed=false"
	}

	db.Scopes(
		helper.Paginate(todos, &paginate, db, 10),
	).Where(wherecondition).Find(&todos)
	paginate.Items = todos

	return paginate
}
