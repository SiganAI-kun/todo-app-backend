package taskdataservices

import (
	database "todo-app-backend/db"
	"todo-app-backend/internal/model"
	"todo-app-backend/internal/repository/taskdataqueries"
)

type ITaskDataMiddlewareService interface {
	// Exec(param GetTaskDataParam) (interface{}, error)
	CheckSameTasks(param CreateTaskDataParam) (bool, error)
	CheckTaskId(param UpdateTaskDataParam) (bool, error)
}

type MiddlewareTaskDataService struct {
	base  model.BaseAppservices
	query taskdataqueries.ITaskDataQuery
}

func NewMiddlewareTaskDataService(dbContext *database.DbContext, q taskdataqueries.ITaskDataQuery) MiddlewareTaskDataService {
	base := model.NewBaseAppservices(dbContext)
	svc := &MiddlewareTaskDataService{
		base:  base,
		query: q,
	}
	return *svc
}

func (s MiddlewareTaskDataService) CheckSameTasks(param CreateTaskDataParam) (bool, error) {
	p, err := taskdataqueries.NewCreateTaskDataParam(
		param.TaskName,
		param.TaskDeadline,
		param.TaskDetails,
	)

	if err != nil {
		return false, err
	}

	res, err := s.query.CheckSameTasksDataQuery(*p)

	if err != nil {
		return false, err
	}

	return res, nil
}

func (s MiddlewareTaskDataService) CheckTaskId(param UpdateTaskDataParam) (bool, error) {
	p, err := taskdataqueries.NewUpdateTaskDataParam(
		param.TaskId,
		param.TaskName,
		param.TaskDeadline,
		param.TaskDetails,
	)

	if err != nil {
		return false, err
	}

	res, err := s.query.CheckTaskExistDataQuery(*p)

	if err != nil {
		return false, err
	}

	return res, nil
}
