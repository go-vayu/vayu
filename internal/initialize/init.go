// Package initialize handles application initialization.
package initialize

import (
	"log/slog"
	"os"

	"github.com/go-vayu/vayu/internal/api/routes"
	"github.com/go-vayu/vayu/internal/db"
	"github.com/go-vayu/vayu/internal/db/migrations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// InitAPI initializes and starts the API server.
func InitAPI() {
	x, err := db.InitDB()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err = migrations.Migrate(x)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.RegisterRoutes(e)

	if err := e.Start("localhost:8000"); err != nil {
		slog.Info("shutting down...")
		os.Exit(1)
	}
}
