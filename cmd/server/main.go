package main

import (
	"todo-app-backend/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  // インスタンスを作成
  e := echo.New()

  // ミドルウェアを設定
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // ルートを設定
  e.GET("/", handler.Hello)
	e.GET("/todos", handler.Get_tasks)
	e.POST("/todos", handler.Create_tasks)
	e.PUT("/todos", handler.Update_tasks)
	e.DELETE("/todos", handler.Delete_tasks)

  e.Logger.Fatal(e.Start(":5001"))
}

// func main() {
// 	http.HandleFunc("/", handler.RestHandler)
// 	http.ListenAndServe(":5001", nil)
// }

