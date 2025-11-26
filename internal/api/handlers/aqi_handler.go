// Package handlers provides HTTP handlers for the API.
package handlers

import (
	"encoding/json"
	"log/slog"
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
// @Param cityID path int true "The City ID"
// @Success 200 {string} string
// @Router /aqi/{cityID} [get]
func (h *AQIHandler) GetAQIByCityID(c echo.Context) error {
	req, err := http.NewRequestWithContext(c.Request().Context(), http.MethodGet, "https://api.openaq.org/v3/locations?countries_id=9", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	req.Header.Add("X-API-Key", config.OpenAQAPIKey.GetString())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			slog.Info(err.Error())
		}
	}()

	var result any
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, result)
}
