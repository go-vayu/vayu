package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AQIHandler struct {
}

func (h *AQIHandler) GetAQIByCity(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}
