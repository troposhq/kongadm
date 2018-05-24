package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func deleteService(cmd *cobra.Command, args []string) {
	if !askForConfirmation("Are you sure?") {
		fmt.Println("User cancelled operation")
		return
	}

	err := client.DeleteService(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Successfully deleted Service")
}

func deleteRoute(cmd *cobra.Command, args []string) {
	if !askForConfirmation("Are you sure?") {
		fmt.Println("User cancelled operation")
		return
	}

	err := client.DeleteRoute(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Successfully deleted Route")
}

func deleteConsumer(cmd *cobra.Command, args []string) {
	if !askForConfirmation("Are you sure?") {
		fmt.Println("User cancelled operation")
		return
	}

	err := client.DeleteConsumer(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Successfully deleted Consumer")
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// DeleteService Command
	deleteCmd.AddCommand(deleteServiceCmd)
	// DeleteRoute Command
	deleteCmd.AddCommand(deleteRouteCmd)
	// DeleteConsumer Command
	deleteCmd.AddCommand(deleteConsumerCmd)
}

var deleteCmd = &cobra.Command{
	Use: "delete [resource]",
}

var deleteServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Delete a Service in Kong",
	Args:  cobra.ExactArgs(1),
	Run:   deleteService,
}

var deleteRouteCmd = &cobra.Command{
	Use:   "route",
	Short: "Delete a Route in Kong",
	Args:  cobra.ExactArgs(1),
	Run:   deleteRoute,
}

var deleteConsumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Delete a Consumer in Kong",
	Args:  cobra.ExactArgs(1),
	Run:   deleteConsumer,
}
