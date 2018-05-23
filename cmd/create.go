package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// func createPlugin(cmd *cobra.Command, args []string) {
// 	if filePath != "" {
// 		file, e := ioutil.ReadFile(filePath)
// 		if e != nil {
// 			fmt.Println("Error reading file. Are you sure the file exists and the path is correct?")
// 			os.Exit(1)
// 		}

// 		plugin, e := client.CreatePlugin(string(file), service, route)
// 		if e != nil {
// 			fmt.Println("Error creating Plugin: ", e.Error())
// 		}

// 		fmt.Printf("%+v", plugin)
// 	}
// }

// func createConsumer(cmd *cobra.Command, args []string) {
// 	if filePath != "" {
// 		file, e := ioutil.ReadFile(filePath)
// 		if e != nil {
// 			fmt.Println("Error reading file. Are you sure the file exists and the path is correct?")
// 			os.Exit(1)
// 		}

// 		api, e := client.CreateConsumer(string(file))
// 		if e != nil {
// 			fmt.Println("Error creating Consumer: ", e)
// 		}

// 		fmt.Printf("%+v", api)
// 	}
// }

func createService(cmd *cobra.Command, args []string) {
	var err error
	if filePath != "" {
		err := unmarshalFile(filePath, &service)
		if err != nil {
			fmt.Println("Error creating Service: ", err.Error())
			os.Exit(1)
		}
	}

	service, err = client.CreateService(service)
	if err != nil {
		fmt.Println("Error creating Service: ", err.Error())
		os.Exit(1)
	}

	prettyPrintStruct(service)
}

func createRoute(cmd *cobra.Command, args []string) {
	var err error
	if filePath != "" {
		err = unmarshalFile(filePath, &route)
		if err != nil {
			fmt.Println("Error unmarshaling Route from file: ", err.Error())
			os.Exit(1)
		}
	}

	route, err := client.CreateRoute(route)
	if err != nil {
		fmt.Println("Error creating Route: ", err.Error())
	}

	prettyPrintStruct(route)
}

func init() {
	createCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Filepath of a json representation of the resource to create.")

	rootCmd.AddCommand(createCmd)
	// createCmd.AddCommand(createConsumerCmd)

	// CreateService Command
	createCmd.AddCommand(createServiceCmd)
	// CreateService Flags
	createServiceCmd.Flags().IntVar(&service.ConnectTimeout, "connect-timeout", 60000, "The timeout in milliseconds for establishing a connection to the upstream servere")
	createServiceCmd.Flags().StringVar(&service.Protocol, "protocol", "http", "The protocol used to communicate with the upstream. It can be one of http (default) or https")
	createServiceCmd.Flags().StringVar(&service.Host, "host", "", "The host of the upstream server")
	createServiceCmd.Flags().IntVar(&service.Port, "port", 80, "The upstream server port")
	createServiceCmd.Flags().StringVar(&service.Path, "path", "", "The path to be used in requests to the upstream server. Empty by default")
	createServiceCmd.Flags().StringVar(&service.Name, "name", "", "The Service name")
	createServiceCmd.Flags().IntVar(&service.Retries, "retries", 5, "The number of retries to execute upon failure to proxy. The default is 5")
	createServiceCmd.Flags().IntVar(&service.ReadTimeout, "read-timeout", 60000, "The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server")
	createServiceCmd.Flags().IntVar(&service.WriteTimeout, "write-timeout", 60000, "The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server")
	createServiceCmd.Flags().StringVar(&service.URL, "url", "", "Shorthand attribute to set protocol, host, port and path at once. This attribute is write-only (the Admin API never \"returns\" the url)")

	// CreateRoute Command
	createCmd.AddCommand(createRouteCmd)
	// CreateRoute Flags
	createRouteCmd.Flags().StringSliceVar(&route.Protocols, "protocols", []string{"http"}, "A list of the protocols this Route should allow. By default it is [\"http\", \"https\"], which means that the Route accepts both. When set to [\"https\"], HTTP requests are answered with a request to upgrade to HTTPS")
	createRouteCmd.Flags().StringSliceVar(&route.Methods, "methods", []string{"GET"}, "A list of HTTP methods that match this Route. For example: [\"GET\", \"POST\"]. At least one of hosts, paths, or methods must be set")
	createRouteCmd.Flags().StringSliceVar(&route.Hosts, "hosts", make([]string, 0), "A list of domain names that match this Route. For example: example.com. At least one of hosts, paths, or methods must be set")
	createRouteCmd.Flags().StringSliceVar(&route.Paths, "paths", make([]string, 0), "A list of paths that match this Route. For example: /my-path. At least one of hosts, paths, or methods must be set")
	createRouteCmd.Flags().BoolVar(&route.StripPath, "strip-path", true, "When matching a Route via one of the paths, strip the matching prefix from the upstream request URL. Defaults to true")
	createRouteCmd.Flags().BoolVar(&route.PreserveHost, "preserve-host", true, "When matching a Route via one of the hosts domain names, use the request Host header in the upstream request headers. By default set to false, and the upstream Host header will be that of the Service's host")
	createRouteCmd.Flags().StringVar(&route.Service.ID, "service", "", "The Service this Route is associated to. This is where the Route proxies traffic to")

	// createPluginCmd.Flags().StringVarP(&serviceNameOrID, "service", "s", "", "ID of service to configure plugin on top of")
	// createPluginCmd.Flags().StringVarP(&route, "route", "r", "", "ID of route to configure plugin on top of")
	// createCmd.AddCommand(createPluginCmd)
}

var createCmd = &cobra.Command{
	Use: "create [resource]",
}

// var createConsumerCmd = &cobra.Command{
// 	Use:   "consumer",
// 	Short: "Create a Consumer",
// 	Run:   createConsumer,
// }

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

// var createPluginCmd = &cobra.Command{
// 	Use:   "plugin",
// 	Short: "Create a Plugin",
// 	Run:   createPlugin,
// }
