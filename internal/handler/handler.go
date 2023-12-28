package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ハンドラーを定義
func Hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}


// func RestHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 			fmt.Fprintln(w, "GET called!!")
// 	} else if r.Method == "POST" {
// 			fmt.Fprintln(w, "POST called!!")
// 	} else if r.Method == "PUT" {
// 			fmt.Fprintln(w, "PUT called!!")
// 	} else if r.Method == "DELETE" {
// 			fmt.Fprintln(w, "DELETE called!!")
// 	}
// }