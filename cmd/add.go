package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func addService(cmd *cobra.Command, args []string) {
	var err error
	if filePath != "" {
		err := unmarshalFile(filePath, &service)
		if err != nil {
			fmt.Println("Error creating Service: ", err.Error())
			os.Exit(1)
		}
	}

	service, err = client.AddService(service)
	if err != nil {
		fmt.Println("Error creating Service: ", err.Error())
		os.Exit(1)
	}

	prettyPrintStruct(service)
}

func addRoute(cmd *cobra.Command, args []string) {
	var err error
	if filePath != "" {
		err = unmarshalFile(filePath, &route)
		if err != nil {
			fmt.Println("Error unmarshaling Route from file: ", err.Error())
			os.Exit(1)
		}
	}

	route, err := client.AddRoute(route)
	if err != nil {
		fmt.Println("Error creating Route: ", err.Error())
	}

	prettyPrintStruct(route)
}

func addPlugin(cmd *cobra.Command, args []string) {
	plugin.Config = pluginConfig.Options
	var err error
	if filePath != "" {
		err := unmarshalFile(filePath, &plugin)
		if err != nil {
			fmt.Println("Error creating Plugin: ", err.Error())
			os.Exit(1)
		}
	}

	plugin, err = client.AddPlugin(plugin)
	if err != nil {
		fmt.Println("Error creating Plugin: ", err.Error())
		os.Exit(1)
	}

	prettyPrintStruct(plugin)
}

func init() {
	addCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Filepath of a json representation of the resource to create.")

	rootCmd.AddCommand(addCmd)

	// AddService Command
	addCmd.AddCommand(addServiceCmd)
	// AddService Flags
	addServiceCmd.Flags().IntVar(&service.ConnectTimeout, "connect-timeout", 60000, "The timeout in milliseconds for establishing a connection to the upstream servere")
	addServiceCmd.Flags().StringVar(&service.Protocol, "protocol", "http", "The protocol used to communicate with the upstream. It can be one of http (default) or https")
	addServiceCmd.Flags().StringVar(&service.Host, "host", "", "The host of the upstream server")
	addServiceCmd.Flags().IntVar(&service.Port, "port", 80, "The upstream server port")
	addServiceCmd.Flags().StringVar(&service.Path, "path", "", "The path to be used in requests to the upstream server. Empty by default")
	addServiceCmd.Flags().StringVar(&service.Name, "name", "", "The Service name")
	addServiceCmd.Flags().IntVar(&service.Retries, "retries", 5, "The number of retries to execute upon failure to proxy. The default is 5")
	addServiceCmd.Flags().IntVar(&service.ReadTimeout, "read-timeout", 60000, "The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server")
	addServiceCmd.Flags().IntVar(&service.WriteTimeout, "write-timeout", 60000, "The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server")
	addServiceCmd.Flags().StringVar(&service.URL, "url", "", "Shorthand attribute to set protocol, host, port and path at once. This attribute is write-only (the Admin API never \"returns\" the url)")

	// AddRoute Command
	addCmd.AddCommand(addRouteCmd)
	// AddRoute Flags
	addRouteCmd.Flags().StringSliceVar(&route.Protocols, "protocols", []string{"http"}, "A list of the protocols this Route should allow. By default it is [\"http\", \"https\"], which means that the Route accepts both. When set to [\"https\"], HTTP requests are answered with a request to upgrade to HTTPS")
	addRouteCmd.Flags().StringSliceVar(&route.Methods, "methods", []string{"GET"}, "A list of HTTP methods that match this Route. For example: [\"GET\", \"POST\"]. At least one of hosts, paths, or methods must be set")
	addRouteCmd.Flags().StringSliceVar(&route.Hosts, "hosts", make([]string, 0), "A list of domain names that match this Route. For example: example.com. At least one of hosts, paths, or methods must be set")
	addRouteCmd.Flags().StringSliceVar(&route.Paths, "paths", make([]string, 0), "A list of paths that match this Route. For example: /my-path. At least one of hosts, paths, or methods must be set")
	addRouteCmd.Flags().BoolVar(&route.StripPath, "strip-path", true, "When matching a Route via one of the paths, strip the matching prefix from the upstream request URL. Defaults to true")
	addRouteCmd.Flags().BoolVar(&route.PreserveHost, "preserve-host", true, "When matching a Route via one of the hosts domain names, use the request Host header in the upstream request headers. By default set to false, and the upstream Host header will be that of the Service's host")
	addRouteCmd.Flags().StringVar(&route.Service.ID, "service", "", "The Service this Route is associated to. This is where the Route proxies traffic to")

	// AddPlugin Command
	addCmd.AddCommand(addPluginCmd)
	// AddPlugin Flags
	addPluginCmd.Flags().StringVar(&plugin.Name, "name", "", "The name of the Plugin that's going to be added")
	addPluginCmd.Flags().StringVarP(&plugin.ServiceID, "service", "s", "", "ID of service to configure plugin on top of")
	addPluginCmd.Flags().StringVarP(&plugin.RouteID, "route", "r", "", "ID of route to configure plugin on top of")
	addPluginCmd.Flags().StringVarP(&plugin.ConsumerID, "consumer", "c", "", "ID of consumer to configure plugin on top of")
	addPluginCmd.Flags().BoolVar(&plugin.Enabled, "enabled", true, "Whether the plugin is applied")
	addPluginCmd.Flags().Var(&pluginConfig, "config", "The configuration properties for the Plugin which can be found on the plugins documentation page in the Plugin Gallery")
}

var addCmd = &cobra.Command{
	Use: "add [resource]",
}

var addServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Add a Service",
	Run:   addService,
}

var addRouteCmd = &cobra.Command{
	Use:   "route",
	Short: "Add a Route",
	Run:   addRoute,
}

var addPluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Add a Plugin",
	Long: `
You can add a plugin in four different ways:
- For every Service/Route and Consumer. Don't set consumer_id and set service_id or route_id.
- For every Service/Route and a specific Consumer. Only set consumer_id.
- For every Consumer and a specific Service. Only set service_id (warning: some plugins only allow setting their route_id)
- For every Consumer and a specific Route. Only set route_id (warning: some plugins only allow setting their service_id)
- For a specific Service/Route and Consumer. Set both service_id/route_id and consumer_id.
Note that not all plugins allow to specify consumer_id. Check the plugin documentation.
	`,
	Run: addPlugin,
}
