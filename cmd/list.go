package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func listAPI(cmd *cobra.Command, args []string) {
	results, err := client.ListAPI()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Results are... %+v", results)
}

func listConsumers(cmd *cobra.Command, args []string) {
	r, err := client.ListConsumers()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", r)
}

func listServices(cmd *cobra.Command, args []string) {
	r, err := client.ListServices()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", r)
}

func listRoutes(cmd *cobra.Command, args []string) {
	r, err := client.ListRoutes()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", r)
}

func listPlugins(cmd *cobra.Command, args []string) {
	r, err := client.ListPlugins()
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
	// konga list route
	listCmd.AddCommand(listRoutesCmd)
	// konga list plugin
	listCmd.AddCommand(listPluginsCmd)
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

var listRoutesCmd = &cobra.Command{
	Use:   "route",
	Short: "List Route objects",
	Run:   listRoutes,
}

var listPluginsCmd = &cobra.Command{
	Use:   "plugin",
	Short: "List Plugin objects",
	Run:   listPlugins,
}
