package repository

import database "todo-app-backend/db"

type BaseQueryService struct {
	DbClient *database.DbContext
}

func NewBaseQueryServices(d *database.DbContext) BaseQueryService {
	b := &BaseQueryService{
		DbClient: d,
	}

	return *b
}
