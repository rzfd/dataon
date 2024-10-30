package exceptions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Success bool        `json:"Success"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

func AppException(c echo.Context, message string) {
	res := Error{
		Success: false,
		Message: message,
		Data:    nil,
	}
	c.JSON(http.StatusInternalServerError, res)
}
