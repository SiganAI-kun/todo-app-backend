package taskdataservices

import (
	database "todo-app-backend/db"
	"todo-app-backend/internal/repository/taskdataqueries"

	"github.com/labstack/echo/v4"
)

type ITaskDataMiddlewareFactory interface {
	// CreateGetTaskDataService(c echo.Context) IGetPreviewDataService
	// CreateCreateTaskDataService(c echo.Context) ICreatePreviewDataService
	CreateTaskDataMiddlewareService(c echo.Context) ITaskDataMiddlewareService
}

type TaskDataMiddlewareFactory struct {
}

func NewTaskDataMiddlewareFactory() ITaskDataMiddlewareFactory {
	return &TaskDataMiddlewareFactory{}
}

func (f *TaskDataMiddlewareFactory) CreateTaskDataMiddlewareService(c echo.Context) ITaskDataMiddlewareService {
	dbContext := database.Current(c.Request().Context())
	// if dbContext == nil {
	//   fmt.Println("Error getting database context")
	//   // あるいはエラーハンドリングを適切に行う
	//   return nil
	// }
	q := taskdataqueries.NewMiddlewareTaskDataQuery(dbContext)
	svc := NewMiddlewareTaskDataService(dbContext, q)
	return svc
}
