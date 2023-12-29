package middleware

import (
	database "todo-app-backend/db"
	"todo-app-backend/middleware/users"

	"github.com/labstack/echo/v4"
)

const profileKey = "PROFILE"

func SetProfile(db *database.DbContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// token := c.Get("user").(*jwt.Token)
			// claims := token.Claims.(jwt.MapClaims)
			// sub := claims["sub"].(string)
			user := &users.User{
				Name: "SiganAI",
			}
			c.Set(profileKey, user)
			// DB をコンテキストにセット
			// Cognito 認証を統合するためのコード（コメントアウト）
			// 例:
			// userID, err := getCognitoUserID(c)
			// if err != nil {
			//     return err
			// }
			// c.Set("userID", userID)

			// 他にも必要な処理があればここで行う			c.Set(profileKey, user)

			// 他にも必要な処理があればここで行う

			return next(c)
		}
	}
}

// func GetProfile(c echo.Context) users.User {
// 	return c.Get(profileKey).(users.User)
// }
