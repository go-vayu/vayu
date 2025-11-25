package cmd

import (
	"log/slog"
	"os"

	"github.com/go-vayu/vayu/internal/api/routes"
	"github.com/go-vayu/vayu/internal/initialize"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts the rest api web server",
	PreRun: func(cmd *cobra.Command, args []string) {
		initialize.FullInit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()

		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		routes.RegisterRoutes(e)

		if err := e.Start("localhost:8000"); err != nil {
			slog.Info("shutting down...")
			os.Exit(1)
		}
	},
}
