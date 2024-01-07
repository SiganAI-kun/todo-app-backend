package taskdataservices

import (
	database "todo-app-backend/db"
	"todo-app-backend/internal/model"
	"todo-app-backend/internal/repository/taskdataqueries"
)

type ITaskDataMiddlewareService interface {
	// Exec(param GetTaskDataParam) (interface{}, error)
	CheckSameTasks(param CreateTaskDataParam) (bool, error)
	CheckTaskId(param DeleteTaskDataParam) (bool, error)
	ChangeParamPutToDel(param UpdateTaskDataParam) (DeleteTaskDataParam)
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

func (s MiddlewareTaskDataService) CheckTaskId(param DeleteTaskDataParam) (bool, error) {
	p, err := taskdataqueries.NewDeleteTaskDataParam(
		param.TaskId,
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

func (s MiddlewareTaskDataService) ChangeParamPutToDel(param UpdateTaskDataParam) (DeleteTaskDataParam) {
	p := &DeleteTaskDataParam{
		TaskId: param.TaskId,
	}

	return *p
}
