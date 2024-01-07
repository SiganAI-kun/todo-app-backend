package taskdataservices

import (
	database "todo-app-backend/db"
	"todo-app-backend/internal/model"
	"todo-app-backend/internal/repository/taskdataqueries"
)

type IUpdatePreviewDataService interface {
	Exec(param UpdateTaskDataParam) (error, error)
}

type UpdatePreviewTaskDataService struct {
	base  model.BaseAppservices
	query taskdataqueries.ITaskDataQuery
}

func NewUpdatePreviewDataService(dbContext *database.DbContext, q taskdataqueries.ITaskDataQuery) UpdatePreviewTaskDataService {
	base := model.NewBaseAppservices(dbContext)
	svc := &UpdatePreviewTaskDataService{
		base:  base,
		query: q,
	}
	return *svc
}

func (s UpdatePreviewTaskDataService) Exec(param UpdateTaskDataParam) (error, error) {
	p, err := taskdataqueries.NewUpdateTaskDataParam(
		param.TaskId,
		param.TaskName,
		param.TaskDeadline,
		param.TaskDetails,
	)

	if err != nil {
		return nil, err
	}

	res, err := s.query.UpdateDataQuery(*p)

	if err != nil {
		return nil, err
	}

	return res, nil
}
