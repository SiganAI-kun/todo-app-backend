package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SampleController struct {
	echo *echo.Echo
}

func NewSampleController(e *echo.Echo) *SampleController {
	sc := &SampleController{
		echo: e,
	}

	return sc
}

func (sc SampleController) Handle() {
	sc.echo.GET("/", sc.Hello).Name = "API0001"
}

// ハンドラーを定義
func (sc SampleController) Hello(c echo.Context) error {
	fmt.Println("aaa")
	return c.JSON(http.StatusOK, "Hello, World!")
}
