package taskdataservices

import (
	"fmt"
	database "todo-app-backend/db"
	"todo-app-backend/internal/model"
	"todo-app-backend/internal/repository/taskdataqueries"
)

type IGetPreviewDataService interface {
	Exec(param GetTaskDataParam) (interface{}, error)
}

type GetPreviewTaskDataService struct {
	base  model.BaseAppservices
	query taskdataqueries.ITaskDataQuery
}

func NewGetPreviewDataService(dbContext *database.DbContext, q taskdataqueries.ITaskDataQuery) GetPreviewTaskDataService {
	base := model.NewBaseAppservices(dbContext)
	svc := &GetPreviewTaskDataService{
		base:  base,
		query: q,
	}
	return *svc
}

func (s GetPreviewTaskDataService) Exec(param GetTaskDataParam) (interface{}, error) {
	fmt.Println(param)
	p, err := taskdataqueries.NewGetTaskDataParam(
		param.Start,
		param.End,
		param.SearchQuery,
	)

	if err != nil {
		return nil, err
	}
	fmt.Println(p)

	data, err := s.query.DataQuery(*p)

	if err != nil {
		return nil, err
	}

	return data, nil
}
