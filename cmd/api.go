package cmd

import (
	"github.com/go-vayu/vayu/internal/config"
	"github.com/go-vayu/vayu/internal/initialize"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts the rest api web server",
	PreRun: func(_ *cobra.Command, _ []string) {
		config.InitConfig()
	},
	Run: func(_ *cobra.Command, _ []string) {
		initialize.InitAPI()
	},
}
