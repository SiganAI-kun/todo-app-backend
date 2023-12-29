package middleware

import "github.com/labstack/echo/v4"

func SetHTTPErrorHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			e := c.Echo()
			e.HTTPErrorHandler = CustomHTTPErrorHandler
			return next(c)
		}
	}
}

func CustomHTTPErrorHandler(err error, c echo.Context) {

}
