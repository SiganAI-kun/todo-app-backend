package api

import (
	"todo-app-backend/internal/handler"

	"github.com/labstack/echo/v4"
)

func Routing(e *echo.Echo) {
	handler.NewSampleController(e).Handle()
	handler.NewTodoController(e).Handle()
}
