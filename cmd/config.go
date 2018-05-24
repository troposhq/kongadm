package cmd

import "github.com/spf13/cobra"

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure options for kongadm",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
