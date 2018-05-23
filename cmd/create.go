package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func createConsumer(cmd *cobra.Command, args []string) {
	var err error
	if filePath != "" {
		err = unmarshalFile(filePath, &consumer)
		if err != nil {
			fmt.Println("Error unmarshaling Consumer from file: ", err.Error())
			os.Exit(1)
		}
	}

	consumer, err := client.CreateConsumer(consumer)
	if err != nil {
		fmt.Println("Error creating Consumer: ", err.Error())
	}

	prettyPrintStruct(consumer)
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Filepath of a json representation of the resource to create.")

	// CreateConsumer Command
	createCmd.AddCommand(createConsumerCmd)
	// CreateConsumer Flags
	createConsumerCmd.Flags().StringVarP(&consumer.Username, "username", "U", "", "The unique username of the consumer. You must send either this field or custom_id with the request")
	createConsumerCmd.Flags().StringVar(&consumer.CustomID, "custom-id", "", "Field for storing an existing unique ID for the consumer - useful for mapping Kong with users in your existing database. You must send either this field or username with the request")
}

var createCmd = &cobra.Command{
	Use: "create [resource]",
}

var createConsumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Create a Consumer",
	Run:   createConsumer,
}
