package router

import (
	"github.com/labstack/echo"
	"github.com/nitoc-ict/fushigidane-server/convertaddress"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.POST("/coordinate", convertaddress.ConvertAddress)

	return e
}
