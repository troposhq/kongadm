package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func listServices(cmd *cobra.Command, args []string) {
	r, err := client.ListServices()
	if err != nil {
		fmt.Println(err)
		return
	}

	prettyPrintStruct(r)
}

func listRoutes(cmd *cobra.Command, args []string) {
	r, err := client.ListRoutes()
	if err != nil {
		fmt.Println(err)
		return
	}

	prettyPrintStruct(r)
}

func listConsumers(cmd *cobra.Command, args []string) {
	r, err := client.ListConsumers()
	if err != nil {
		fmt.Println(err)
		return
	}

	prettyPrintStruct(r)
}

func listPlugins(cmd *cobra.Command, args []string) {
	r, err := client.ListPlugins()
	if err != nil {
		fmt.Println(err)
		return
	}

	prettyPrintStruct(r)
}

func init() {
	rootCmd.AddCommand(listCmd)

	// ListServices Command
	listCmd.AddCommand(listServicesCmd)

	// ListRoutes Command
	listCmd.AddCommand(listRoutesCmd)

	// ListConsumers Command
	listCmd.AddCommand(listConsumersCmd)

	// ListPlugins Command
	listCmd.AddCommand(listPluginsCmd)
}

var listCmd = &cobra.Command{
	Use: "list [resource]",
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

var listConsumersCmd = &cobra.Command{
	Use:   "consumer",
	Short: "List Consumer objects",
	Run:   listConsumers,
}

var listPluginsCmd = &cobra.Command{
	Use:   "plugin",
	Short: "List Plugin objects",
	Run:   listPlugins,
}
