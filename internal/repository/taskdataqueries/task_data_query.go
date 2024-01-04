package taskdataqueries

import (
	database "todo-app-backend/db"
	"todo-app-backend/internal/repository"
)

type ITaskDataQuery interface {
	DataQuery(TaskDataParam) ([]TaskDataResponse, error)
}

type TaskDataQuery struct {
	base repository.BaseQueryService
}

func NewGetTaskDataQuery(dbContext *database.DbContext) ITaskDataQuery {
	base := repository.NewBaseQueryServices(dbContext)
	q := &TaskDataQuery{
		base: base,
	}

	return *q
}
