package cmd

import (
	"log/slog"
	"os"

	"github.com/go-vayu/vayu/internal/api/routes"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts the rest api web server",
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()
		routes.RegisterRoutes(e)

		if err := e.Start("localhost:8000"); err != nil {
			slog.Info("shutting down...")
			os.Exit(1)
		}
	},
}
