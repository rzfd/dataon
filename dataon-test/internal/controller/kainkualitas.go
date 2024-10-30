package controller

import (
	"net/http"

	"github.com/rzfd/dataon-test/internal/config"
	"github.com/rzfd/dataon-test/internal/exceptions"
	"github.com/rzfd/dataon-test/internal/model"

	"github.com/labstack/echo/v4"
)

func GetAllKainKualitas(c echo.Context) error {
	var kainKualitasList []model.KainKualitas

	if err := config.DB.Preload("Kain").Preload("Kualitas").Preload("Kain.JenisKain").Find(&kainKualitasList).Error; err != nil {
		exceptions.AppException(c, err.Error())
	}

	return c.JSON(http.StatusOK, kainKualitasList)
}
