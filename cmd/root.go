package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/troposhq/konga/api"
)

var rootCmd = &cobra.Command{
	Use:   "konga",
	Short: "A CLI tool for interacting the the Kong API",
	Long:  "A CLI tool for interacting the the Kong API",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&api.APIURLBase, "url", "localhost:8001", "URL for the Kong Admin API")
}

// Execute is the entrypoint for the CLI called from the main function
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
