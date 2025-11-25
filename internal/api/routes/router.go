package routes

import (
	"github.com/go-vayu/vayu/internal/api/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	a := e.Group("/api/v1")

	aqiHandler := &handlers.AQIHandler{}

	a.GET("/aqi/{city}", aqiHandler.GetAQIByCity)
}
