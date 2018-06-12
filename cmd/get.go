package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getServices(cmd *cobra.Command, args []string) {
	var r interface{}
	var err error

	if len(args) > 0 {
		nameOrID := args[0]
		r, err = client.GetService(nameOrID)
	} else {
		r, err = client.ListServices()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	prettyPrintStruct(r)
}

func getRoutes(cmd *cobra.Command, args []string) {
	var r interface{}
	var err error

	if len(args) > 0 {
		nameOrID := args[0]
		r, err = client.GetRoute(nameOrID)
	} else {
		r, err = client.ListRoutes()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	prettyPrintStruct(r)
}

func getConsumers(cmd *cobra.Command, args []string) {
	var r interface{}
	var err error

	if len(args) > 0 {
		nameOrID := args[0]
		r, err = client.GetConsumer(nameOrID)
	} else {
		r, err = client.ListConsumers()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	prettyPrintStruct(r)
}

func getPlugins(cmd *cobra.Command, args []string) {
	r, err := client.ListPlugins()
	if err != nil {
		fmt.Println(err)
		return
	}

	prettyPrintStruct(r)
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.AddCommand(getServicesCmd)
	getCmd.AddCommand(getRoutesCmd)
	getCmd.AddCommand(getConsumersCmd)
	getCmd.AddCommand(getPluginsCmd)
}

var getCmd = &cobra.Command{
	Use: "get [resource]",
}

var getServicesCmd = &cobra.Command{
	Use:     "services [id]",
	Aliases: []string{"service", "svc"},
	Short:   "Get Service objects. Specify a name or id to get one service.",
	Run:     getServices,
}

var getRoutesCmd = &cobra.Command{
	Use:     "routes [id]",
	Aliases: []string{"route", "r"},
	Short:   "Get Route objects. Specify an id to get one route.",
	Run:     getRoutes,
}

var getConsumersCmd = &cobra.Command{
	Use:     "consumers [id]",
	Aliases: []string{"consumer", "c"},
	Short:   "Get Consumer objects. Specify an id or username to get one consumer.",
	Run:     getConsumers,
}

var getPluginsCmd = &cobra.Command{
	Use:     "plugins [id]",
	Aliases: []string{"plugins", "p"},
	Short:   "Get Plugin objects. Specify an id to get one plugin.",
	Run:     getPlugins,
}
