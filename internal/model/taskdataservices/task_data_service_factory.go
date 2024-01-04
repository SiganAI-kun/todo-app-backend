package taskdataservices

import (
	database "todo-app-backend/db"
	"todo-app-backend/internal/repository/taskdataqueries"

	"github.com/labstack/echo/v4"
)

type ITaskDataServiceFactory interface {
	CreateGetTaskDataService(c echo.Context) IGetPreviewDataService
}

type TaskDataServiceFactory struct {
}

func NewTaskDataServiceFactory() ITaskDataServiceFactory {
	return &TaskDataServiceFactory{}
}

func (f *TaskDataServiceFactory) CreateGetTaskDataService(c echo.Context) IGetPreviewDataService {
	dbContext := database.Current(c.Request().Context())
	// if dbContext == nil {
	//   fmt.Println("Error getting database context")
	//   // あるいはエラーハンドリングを適切に行う
	//   return nil
	// }
	q := taskdataqueries.NewGetTaskDataQuery(dbContext)
	svc := NewGetPreviewDataService(dbContext, q)
	return svc
}
