package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ハンドラーを定義
func Get_tasks(c echo.Context) error {
  return c.String(http.StatusOK, "Get_tasks")
}

func Create_tasks(c echo.Context) error {
  return c.String(http.StatusOK, "Create_tasks")
}

func Update_tasks(c echo.Context) error {
  return c.String(http.StatusOK, "Update_tasks")
}

func Delete_tasks(c echo.Context) error {
  return c.String(http.StatusOK, "Delete_tasks")
}
