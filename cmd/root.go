package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/troposhq/konga/api"
)

var rootCmd = &cobra.Command{
	Use:   "konga",
	Short: "A CLI tool for interacting with the Kong API",
	Long:  "A CLI tool for interacting with the Kong API",
}

var client *api.KongAdminAPIClient

func init() {
	apiURLBase := rootCmd.PersistentFlags().String("url", "localhost:8001", "URL for the Kong Admin API")
	client = api.New(*apiURLBase)
}

// Execute is the entrypoint for the CLI called from the main function
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
