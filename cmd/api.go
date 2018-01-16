package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(apiCmd)
	apiCmd.AddCommand(apiListCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Interact with API objects in Kong",
}

var apiListCmd = &cobra.Command{
	Use:   "list",
	Short: "List API objects",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world")
	},
}
