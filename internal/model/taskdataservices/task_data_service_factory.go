package taskdataservices

import (
	"fmt"
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
	fmt.Println("test1")
	dbContext := database.Current(c.Request().Context())
	fmt.Println(dbContext)
	// if dbContext == nil {
  //   fmt.Println("Error getting database context")
  //   // あるいはエラーハンドリングを適切に行う
  //   return nil
	// }
	fmt.Println("test2")
	q := taskdataqueries.NewGetTaskDataQuery(dbContext)
	fmt.Println("test3")
	svc := NewGetPreviewDataService(dbContext, q)
	fmt.Println("test4")
	return svc
}
