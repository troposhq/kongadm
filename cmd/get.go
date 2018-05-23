package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getService(cmd *cobra.Command, args []string) {
	nameOrID := args[0]
	service, err := client.GetService(nameOrID)
	if err != nil {
		fmt.Println(err)
		return
	}

	prettyPrintStruct(service)
}

func getConsumer(cmd *cobra.Command, args []string) {
	usernameOrID := args[0]
	consumer, err := client.GetConsumer(usernameOrID)
	if err != nil {
		fmt.Println(err)
		return
	}

	prettyPrintStruct(consumer)
}

func init() {
	rootCmd.AddCommand(getCmd)

	// GetService Command
	getCmd.AddCommand(getServiceCmd)

	// GetConsumer Command
	getCmd.AddCommand(getConsumerCmd)
}

var getCmd = &cobra.Command{
	Use: "get [resource]",
}

var getServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Get a Service by name or id",
	Args:  cobra.ExactArgs(1),
	Run:   getService,
}

var getConsumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Get a Consumer by username or id",
	Args:  cobra.ExactArgs(1),
	Run:   getConsumer,
}
