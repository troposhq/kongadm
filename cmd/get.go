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

func getConsumer(cmd *cobra.Command, args []string) {
	usernameOrID := args[0]
	consumer, err := api.GetConsumer(usernameOrID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", consumer)
}

func getService(cmd *cobra.Command, args []string) {
	nameOrID := args[0]
	service, err := api.GetService(nameOrID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", service)
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getAPICmd)
	getCmd.AddCommand(getConsumerCmd)
	getCmd.AddCommand(getServiceCmd)
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

var getConsumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Get a Consumer by username or id",
	Args:  cobra.ExactArgs(1),
	Run:   getConsumer,
}

var getServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Get a Service by name or id",
	Args:  cobra.ExactArgs(1),
	Run:   getService,
}
