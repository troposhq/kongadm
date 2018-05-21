package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func deleteAPI(cmd *cobra.Command, args []string) {
	err := client.DeleteAPI(args[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Successfully deleted API")
}

func deleteConsumer(cmd *cobra.Command, args []string) {
	err := client.DeleteConsumer(args[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Successfully deleted Consumer")
}

func deleteService(cmd *cobra.Command, args []string) {
	err := client.DeleteService(args[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Successfully deleted Service")
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteAPICmd)
	deleteCmd.AddCommand(deleteConsumerCmd)
	deleteCmd.AddCommand(deleteServiceCmd)
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
