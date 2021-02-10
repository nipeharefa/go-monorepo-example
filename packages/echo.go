package packages

import "github.com/labstack/echo/v4"

//go:noinline
func NewEchoServer() *echo.Echo {
	e := echo.New()

	return e
}
