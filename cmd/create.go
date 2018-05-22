package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var filePath string
var service string
var route string

func createAPI(cmd *cobra.Command, args []string) {
	if filePath != "" {
		file, e := ioutil.ReadFile(filePath)
		if e != nil {
			fmt.Println("Error reading file. Are you sure the file exists and the path is correct?")
			os.Exit(1)
		}

		api, e := client.CreateAPI(string(file))
		if e != nil {
			fmt.Println("Error creating API: ", e)
		}

		fmt.Printf("%+v", api)
	}
}

func createPlugin(cmd *cobra.Command, args []string) {
	if filePath != "" {
		file, e := ioutil.ReadFile(filePath)
		if e != nil {
			fmt.Println("Error reading file. Are you sure the file exists and the path is correct?")
			os.Exit(1)
		}

		plugin, e := client.CreatePlugin(string(file), service, route)
		if e != nil {
			fmt.Println("Error creating Plugin: ", e.Error())
		}

		fmt.Printf("%+v", plugin)
	}
}

func createConsumer(cmd *cobra.Command, args []string) {
	if filePath != "" {
		file, e := ioutil.ReadFile(filePath)
		if e != nil {
			fmt.Println("Error reading file. Are you sure the file exists and the path is correct?")
			os.Exit(1)
		}

		api, e := client.CreateConsumer(string(file))
		if e != nil {
			fmt.Println("Error creating Consumer: ", e)
		}

		fmt.Printf("%+v", api)
	}
}

func createService(cmd *cobra.Command, args []string) {
	if filePath != "" {
		file, e := ioutil.ReadFile(filePath)
		if e != nil {
			fmt.Println("Error reading file. Are you sure the file exists and the path is correct?")
			os.Exit(1)
		}

		service, e := client.CreateService(string(file))
		if e != nil {
			fmt.Println("Error creating Service: ", e)
		}

		fmt.Printf("%+v", service)
	}
}

func createRoute(cmd *cobra.Command, args []string) {
	if filePath != "" {
		file, e := ioutil.ReadFile(filePath)
		if e != nil {
			fmt.Println("Error reading file. Are you sure the file exists and the path is correct?")
			os.Exit(1)
		}

		route, e := client.CreateRoute(string(file))
		if e != nil {
			fmt.Println("Error creating Route: ", e)
		}

		fmt.Printf("%+v", route)
	}
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createAPICmd)
	createCmd.AddCommand(createConsumerCmd)
	createCmd.AddCommand(createServiceCmd)
	createCmd.AddCommand(createRouteCmd)

	createPluginCmd.Flags().StringVarP(&service, "service", "s", "", "ID of service to configure plugin on top of")
	createPluginCmd.Flags().StringVarP(&route, "route", "r", "", "ID of route to configure plugin on top of")
	createCmd.AddCommand(createPluginCmd)

	createCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Filepath of a json representation of the resource to create.")
}

var createCmd = &cobra.Command{
	Use: "create [resource]",
}

var createAPICmd = &cobra.Command{
	Use:   "api",
	Short: "Create an API",
	Run:   createAPI,
}

var createConsumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Create a Consumer",
	Run:   createConsumer,
}

var createServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Create a Service",
	Run:   createService,
}

var createRouteCmd = &cobra.Command{
	Use:   "route",
	Short: "Create a Route",
	Run:   createRoute,
}

var createPluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Create a Plugin",
	Run:   createPlugin,
}
