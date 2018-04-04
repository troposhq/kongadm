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

func listConsumers(cmd *cobra.Command, args []string) {
	r, err := api.ListConsumers()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", r)
}

func listServices(cmd *cobra.Command, args []string) {
	r, err := api.ListServices()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", r)
}

func init() {
	// konga list
	rootCmd.AddCommand(listCmd)
	// konga list api
	listCmd.AddCommand(listAPICmd)
	// konga list consumer
	listCmd.AddCommand(listConsumersCmd)
	// konga list service
	listCmd.AddCommand(listServicesCmd)
}

var listCmd = &cobra.Command{
	Use: "list [resource]",
}

var listAPICmd = &cobra.Command{
	Use:   "api",
	Short: "List API objects",
	Run:   listAPI,
}

var listConsumersCmd = &cobra.Command{
	Use:   "consumer",
	Short: "List Consumer objects",
	Run:   listConsumers,
}

var listServicesCmd = &cobra.Command{
	Use:   "service",
	Short: "List Service objects",
	Run:   listServices,
}
