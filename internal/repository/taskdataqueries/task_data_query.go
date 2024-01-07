package taskdataqueries

import (
	database "todo-app-backend/db"
	"todo-app-backend/internal/repository"
)

type ITaskDataQuery interface {
	GetDataQuery(GetTaskDataParam) ([]TaskDataResponse, error)
	CreateDataQuery(CreateTaskDataParam) (error, error)
	UpdateDataQuery(UpdateTaskDataParam) (error, error)
	CheckSameTasksDataQuery(CreateTaskDataParam) (bool, error)
	CheckTaskExistDataQuery(UpdateTaskDataParam) (bool, error)
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

func NewCreateTaskDataQuery(dbContext *database.DbContext) ITaskDataQuery {
	base := repository.NewBaseQueryServices(dbContext)
	q := &TaskDataQuery{
		base: base,
	}

	return *q
}

func NewUpdateTaskDataQuery(dbContext *database.DbContext) ITaskDataQuery {
	base := repository.NewBaseQueryServices(dbContext)
	q := &TaskDataQuery{
		base: base,
	}

	return *q
}

func NewMiddlewareTaskDataQuery(dbContext *database.DbContext) ITaskDataQuery {
	base := repository.NewBaseQueryServices(dbContext)
	q := &TaskDataQuery{
		base: base,
	}

	return *q
}
