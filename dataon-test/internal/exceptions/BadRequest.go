package exceptions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func BadRequest(c echo.Context, message string) {
	res := Error{
		Success: false,
		Message: message,
		Data:    nil,
	}
	c.JSON(http.StatusBadRequest, res)
}
