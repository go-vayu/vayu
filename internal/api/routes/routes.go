// Package routes defines API route registration.
package routes

import (
	"github.com/go-vayu/vayu/internal/api/handlers"
	_ "github.com/go-vayu/vayu/internal/api/swagger" // registers swagger docs
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// RegisterRoutes sets up all API routes.
// @title Vayu API
// @version 1.0
// @description This is the documentation for the Vayu API. Vayu is a open-source environmental monitoring platform tracking air quality, emissions, and climate data across India.
// @BasePath /api/v1
// @license.name AGPL-3.0-or-later
func RegisterRoutes(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	a := e.Group("/api/v1")

	aqiHandler := &handlers.AQIHandler{}

	a.GET("/aqi/{cityID}", aqiHandler.GetAQIByCityID)
}
