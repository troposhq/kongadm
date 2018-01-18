package cmd

import (
	"fmt"

	"github.com/aloisbarreras/konga/api"
	"github.com/spf13/cobra"
)

func getAPI(cmd *cobra.Command, args []string) {
	nameOrID := args[0]
	api, err := api.GetAPI(nameOrID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", api)
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getAPICmd)
}

var getCmd = &cobra.Command{
	Use: "get [resource]",
}

var getAPICmd = &cobra.Command{
	Use:   "api",
	Short: "Get an API by name or id",
	Args:  cobra.ExactArgs(1),
	Run:   getAPI,
}
