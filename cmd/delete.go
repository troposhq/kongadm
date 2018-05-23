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
	// DeleteConsumer Command
	deleteCmd.AddCommand(deleteConsumerCmd)
}

var deleteCmd = &cobra.Command{
	Use: "delete [resource]",
}

var deleteConsumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Delete a Consumer in Kong",
	Args:  cobra.ExactArgs(1),
	Run:   deleteConsumer,
}

var deleteServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Delete a Service in Kong",
	Args:  cobra.ExactArgs(1),
	Run:   deleteService,
}
