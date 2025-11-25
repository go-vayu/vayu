package handlers

import (
	"net/http"

	"github.com/go-vayu/vayu/internal/config"
	"github.com/labstack/echo/v4"
)

type AQIHandler struct {
}

// @Summary Get AQI data by city
// @tags aqi
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router /aqi/{city} [get]
func (h *AQIHandler) GetAQIByCity(c echo.Context) error {
	return c.JSON(http.StatusOK, config.OPENAQ_API_KEY.GetString())
}
