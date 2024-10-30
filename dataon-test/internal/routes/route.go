package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/rzfd/dataon-test/internal/controller"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/kual/all", controller.GetAllKainKualitas)
}
