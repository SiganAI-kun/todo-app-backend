package main

import (
	"todo-app-backend/api"
	"todo-app-backend/cmd"
	database "todo-app-backend/db"
	myMiddleware "todo-app-backend/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var DB *database.DbContext

func main() {
	env := cmd.GetEnvironmentVariables()
	DB = database.NewDatabase(env.Dsn)
	// インスタンスを作成
	e := echo.New()

	// if os.Getenv("APP_ENV") != "local" {
	// 	ecs.Init()
	// 	e.Use(echo.WrapMiddleware(func(h http.Handler) http.Handler {
	// 		return xray.Handler(xray.NewFixedSegmentNamer("todoapp-api"), h)
	// 	}))
	// }

	// ミドルウェアを設定
	e.Use(middleware.CORS())
	// e.Use(middleware.JWTWithConfig())
	e.Use(myMiddleware.SetProfile(DB))
	e.Use(myMiddleware.Database(DB))
	e.Use(myMiddleware.SetHTTPErrorHandler())
	e.Validator = myMiddleware.NewCustomValidator()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	api.Routing(e)
	e.Logger.Fatal(e.Start(":5000"))
}

// func main() {
// 	http.HandleFunc("/", handler.RestHandler)
// 	http.ListenAndServe(":5001", nil)
// }
