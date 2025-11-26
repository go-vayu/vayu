// Package handlers provides HTTP handlers for the API.
package handlers

import (
	"net/http"

	"github.com/go-vayu/vayu/internal/config"
	"github.com/labstack/echo/v4"
)

// AQIHandler handles AQI-related HTTP requests.
type AQIHandler struct {
}

// GetAQIByCityID retrieves AQI data for a specific city.
// @Summary Get AQI data by city
// @tags aqi
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Router /aqi/{city} [get]
func (h *AQIHandler) GetAQIByCityID(c echo.Context) error {
	return c.JSON(http.StatusOK, config.OpenAQAPIKey.GetString())
}
