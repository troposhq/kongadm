package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func updateService(cmd *cobra.Command, args []string) {
	var err error

	service, err = client.UpdateService(service)
	if err != nil {
		fmt.Println("Error updating Service: ", err.Error())
		os.Exit(1)
	}

	prettyPrintStruct(service)
}

func updateRoute(cmd *cobra.Command, args []string) {
	var err error

	route, err = client.UpdateRoute(route)
	if err != nil {
		fmt.Println("Error updating Route: ", err.Error())
		os.Exit(1)
	}

	prettyPrintStruct(route)
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// UpdateService Command
	updateCmd.AddCommand(updateServiceCmd)
	// UpdateService Flags
	updateServiceCmd.Flags().StringVar(&service.ID, "id", "", "The ID of the service")
	updateServiceCmd.Flags().IntVar(&service.ConnectTimeout, "connect-timeout", 60000, "The timeout in milliseconds for establishing a connection to the upstream servere")
	updateServiceCmd.Flags().StringVar(&service.Protocol, "protocol", "http", "The protocol used to communicate with the upstream. It can be one of http (default) or https")
	updateServiceCmd.Flags().StringVar(&service.Host, "host", "", "The host of the upstream server")
	updateServiceCmd.Flags().IntVar(&service.Port, "port", 80, "The upstream server port")
	updateServiceCmd.Flags().StringVar(&service.Path, "path", "", "The path to be used in requests to the upstream server. Empty by default")
	updateServiceCmd.Flags().StringVar(&service.Name, "name", "", "The Service name")
	updateServiceCmd.Flags().IntVar(&service.Retries, "retries", 5, "The number of retries to execute upon failure to proxy. The default is 5")
	updateServiceCmd.Flags().IntVar(&service.ReadTimeout, "read-timeout", 60000, "The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server")
	updateServiceCmd.Flags().IntVar(&service.WriteTimeout, "write-timeout", 60000, "The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server")
	updateServiceCmd.Flags().StringVar(&service.URL, "url", "", "Shorthand attribute to set protocol, host, port and path at once. This attribute is write-only (the Admin API never \"returns\" the url)")
}

var updateCmd = &cobra.Command{
	Use: "update [resource]",
}

var updateServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Update a Service",
	Run:   updateService,
}

var updateRouteCmd = &cobra.Command{
	Use:   "route",
	Short: "Update a Route",
	Run:   updateRoute,
}
