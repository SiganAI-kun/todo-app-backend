package middleware

import (
	"context"
	"fmt"
	"net/http"
	database "todo-app-backend/db"
	"todo-app-backend/middleware/users"

	"github.com/labstack/echo/v4"
)

// CustomKeyType は独自の型を定義した例です
type CustomKeyType string

const (
	currentUserNameKey CustomKeyType = "CURRENT_USER_NAME"
	dbContextKey      CustomKeyType = "DB"
)

func Database(db database.IDbContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
					profile := c.Get("PROFILE")
					if profile == nil {
							fmt.Println("PROFILE is nil")
							return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
					}

					userPtr, ok := profile.(*users.User)
					if !ok {
							fmt.Printf("Failed to convert PROFILE to *users.User, got type %T\n", profile)
							return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
					}

					user := *userPtr  // ポインタから値に変換

					ctxWithUser := context.WithValue(c.Request().Context(), "CURRENT_USER_NAME", user.Name)
					dbContext := db.WithContext(ctxWithUser)
					ctx := context.WithValue(c.Request().Context(), "DB", dbContext)
					req := c.Request().WithContext(ctx)
					c.SetRequest(req)
					return next(c)
			}
	}
}
