package taskdataservices

import (
	database "todo-app-backend/db"
	"todo-app-backend/internal/model"
	"todo-app-backend/internal/repository/taskdataqueries"
)

type ICreatePreviewDataService interface {
	Exec(param CreateTaskDataParam) (error, error)
}

type CreatePreviewTaskDataService struct {
	base  model.BaseAppservices
	query taskdataqueries.ITaskDataQuery
}

func NewCreatePreviewDataService(dbContext *database.DbContext, q taskdataqueries.ITaskDataQuery) CreatePreviewTaskDataService {
	base := model.NewBaseAppservices(dbContext)
	svc := &CreatePreviewTaskDataService{
		base:  base,
		query: q,
	}
	return *svc
}

func (s CreatePreviewTaskDataService) Exec(param CreateTaskDataParam) (error, error) {
	p, err := taskdataqueries.NewCreateTaskDataParam(
		param.TaskName,
		param.TaskDeadline,
		param.TaskDetails,
	)

	if err != nil {
		return nil, err
	}

	res, err := s.query.CreateDataQuery(*p)

	if err != nil {
		return nil, err
	}

	return res, nil
}
