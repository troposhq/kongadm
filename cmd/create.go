package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aloisbarreras/konga/api"
	"github.com/spf13/cobra"
)

var filePath string

func createAPI(cmd *cobra.Command, args []string) {
	if filePath != "" {
		file, e := ioutil.ReadFile(filePath)
		if e != nil {
			fmt.Println("Error reading file. Are you sure the file exists and the path is correct?")
			os.Exit(1)
		}

		api, e := api.CreateAPI(string(file))
		if e != nil {
			fmt.Println("Error creating API: ", e)
		}

		fmt.Printf("%+v", api)
	}
}

func createConsumer(cmd *cobra.Command, args []string) {
	if filePath != "" {
		file, e := ioutil.ReadFile(filePath)
		if e != nil {
			fmt.Println("Error reading file. Are you sure the file exists and the path is correct?")
			os.Exit(1)
		}

		api, e := api.CreateConsumer(string(file))
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

		api, e := api.CreateService(string(file))
		if e != nil {
			fmt.Println("Error creating Service: ", e)
		}

		fmt.Printf("%+v", api)
	}
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createAPICmd)
	createCmd.AddCommand(createConsumerCmd)
	createCmd.AddCommand(createServiceCmd)

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
