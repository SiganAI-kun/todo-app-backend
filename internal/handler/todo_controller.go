package handler

import (
	"fmt"
	"log"
	"net/http"
	"todo-app-backend/internal/model/taskdataservices"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	echo           *echo.Echo
	serviceFactory taskdataservices.ITaskDataServiceFactory
}

func NewTodoController(e *echo.Echo) *TodoController {
	tc := &TodoController{
		echo: e,
		serviceFactory: taskdataservices.NewTaskDataServiceFactory(),
	}

	return tc
}

func (tc TodoController) Handle() {
	tc.echo.POST("/todos", tc.Create_tasks).Name = "API0101"
	tc.echo.GET("/todos", tc.Get_tasks).Name = "API0102"
	tc.echo.PUT("/todos", tc.Update_tasks).Name = "API0103"
	tc.echo.DELETE("/todos", tc.Delete_tasks).Name = "API0104"
}

// ハンドラーを定義
func (tc TodoController) Get_tasks(c echo.Context) error {
	param := new(taskdataservices.GetTaskDataParam)
	fmt.Println(param.IsFilterd)
	err := c.Bind(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err = c.Validate(param); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	fmt.Println("ttt")

	svc := tc.serviceFactory.CreateGetTaskDataService(c)
	if svc == nil {
		log.Println("Error: GetTaskDataService is nil")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create GetTaskDataService"})
	}
	fmt.Println("check1")
	data, err := svc.Exec(*param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	// return c.JSON(http.StatusOK, "Get_tasks")
	return c.JSON(http.StatusOK, &data)
}

func (tc TodoController) Create_tasks(c echo.Context) error {
	return c.JSON(http.StatusOK, "Create_tasks")
}

func (tc TodoController) Update_tasks(c echo.Context) error {
	return c.JSON(http.StatusOK, "Update_tasks")
}

func (tc TodoController) Delete_tasks(c echo.Context) error {
	return c.JSON(http.StatusOK, "Delete_tasks")
}