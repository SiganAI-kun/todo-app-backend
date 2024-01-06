package taskdataservices

import (
	database "todo-app-backend/db"
	"todo-app-backend/internal/repository/taskdataqueries"

	"github.com/labstack/echo/v4"
)

type ITaskDataServiceFactory interface {
	CreateGetTaskDataService(c echo.Context) IGetPreviewDataService
	CreateCreateTaskDataService(c echo.Context) ICreatePreviewDataService
	CreateMiddlewareTaskDataService(c echo.Context) ITaskDataMiddlewareService
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

func (f *TaskDataServiceFactory) CreateCreateTaskDataService(c echo.Context) ICreatePreviewDataService {
	dbContext := database.Current(c.Request().Context())
	// if dbContext == nil {
	//   fmt.Println("Error getting database context")
	//   // あるいはエラーハンドリングを適切に行う
	//   return nil
	// }
	q := taskdataqueries.NewCreateTaskDataQuery(dbContext)
	svc := NewCreatePreviewDataService(dbContext, q)
	return svc
}

func (f *TaskDataServiceFactory) CreateMiddlewareTaskDataService(c echo.Context) ITaskDataMiddlewareService {
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
