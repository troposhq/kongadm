package cmd

import (
	"fmt"

	"github.com/aloisbarreras/konga/api"
	"github.com/spf13/cobra"
)

func init() {
	// konga api
	rootCmd.AddCommand(apiCmd)

	// konga api list
	apiCmd.AddCommand(apiListCmd)
}

func apiList(cmd *cobra.Command, args []string) {
	results, err := api.ListAPI()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Results are...", results)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Interact with API objects in Kong",
}

var apiListCmd = &cobra.Command{
	Use:   "list",
	Short: "List API objects",
	Run:   apiList,
}
