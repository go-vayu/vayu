// Package cmd provides CLI commands for the application.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vayu",
	Short: "Environmental data API and monitoring platform for India",
	Long: `Vayu tracks air quality, emissions, and climate data across India.

An open-source platform aggregating environmental data from public sources,
making it freely accessible through a REST API and web interface.

Empowering citizens with transparent environmental information.`,
	PreRun: apiCmd.PreRun,
	Run:    apiCmd.Run,
}

// Execute starts the application
func Execute() {
	rootCmd.AddCommand(apiCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
