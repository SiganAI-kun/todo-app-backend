package taskdataservices

import (
	database "todo-app-backend/db"
	"todo-app-backend/internal/model"
	"todo-app-backend/internal/repository/taskdataqueries"
)

type IDeletePreviewDataService interface {
	Exec(param DeleteTaskDataParam) (error, error)
}

type DeletePreviewTaskDataService struct {
	base  model.BaseAppservices
	query taskdataqueries.ITaskDataQuery
}

func NewDeletePreviewDataService(dbContext *database.DbContext, q taskdataqueries.ITaskDataQuery) DeletePreviewTaskDataService {
	base := model.NewBaseAppservices(dbContext)
	svc := &DeletePreviewTaskDataService{
		base:  base,
		query: q,
	}
	return *svc
}

func (s DeletePreviewTaskDataService) Exec(param DeleteTaskDataParam) (error, error) {
	p, err := taskdataqueries.NewDeleteTaskDataParam(
		param.TaskId,
	)

	if err != nil {
		return nil, err
	}

	res, err := s.query.DeleteDataQuery(*p)

	if err != nil {
		return nil, err
	}

	return res, nil
}
