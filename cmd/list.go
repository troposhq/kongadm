package cmd

import (
	"fmt"

	"github.com/aloisbarreras/konga/api"
	"github.com/spf13/cobra"
)

func listAPI(cmd *cobra.Command, args []string) {
	results, err := api.ListAPI()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Results are... %+v", results)
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listAPICmd)
}

var listCmd = &cobra.Command{
	Use: "list [resource]",
}

var listAPICmd = &cobra.Command{
	Use:   "api",
	Short: "List API objects",
	Run:   listAPI,
}
