package model

import database "todo-app-backend/db"

type BaseAppservices struct {
	DbClient database.IDbContext
}

func NewBaseAppservices(d database.IDbContext) BaseAppservices {
	svc := &BaseAppservices{
		DbClient: d,
	}
	return *svc
}
