package cmd

import (
	"fmt"

	"github.com/aloisbarreras/konga/api"
	"github.com/spf13/cobra"
)

func deleteAPI(cmd *cobra.Command, args []string) {
	err := api.DeleteAPI(args[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Successfully deleted API")
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteAPICmd)
}

var deleteCmd = &cobra.Command{
	Use: "delete [resource]",
}

var deleteAPICmd = &cobra.Command{
	Use:   "api",
	Short: "Delete an API in Kong",
	Args:  cobra.ExactArgs(1),
	Run:   deleteAPI,
}
