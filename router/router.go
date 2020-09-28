package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nitoc-ict/fushigidane-server/convertaddress"
	"github.com/nitoc-ict/fushigidane-server/gettransitpoints"
	"github.com/nitoc-ict/fushigidane-server/insertlocation"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/coordinate", convertaddress.ConvertAddress)
	e.POST("/route", gettransitpoints.GetTransitPoints)
	e.POST("/insertlocation", insertlocation.RecvLocationData)

	return e
}
