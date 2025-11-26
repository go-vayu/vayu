package initialize

import (
	"log/slog"
	"os"

	"github.com/go-vayu/vayu/internal/api/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitAPI() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.RegisterRoutes(e)

	if err := e.Start("localhost:8000"); err != nil {
		slog.Info("shutting down...")
		os.Exit(1)
	}
}
