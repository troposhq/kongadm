package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func updatePlugin(cmd *cobra.Command, args []string) {
	var err error
	if filePath != "" {
		err = unmarshalFile(filePath, &plugin)
		if err != nil {
			fmt.Println("Error unmarshaling Plugin from file: ", err.Error())
			os.Exit(1)
		}
	}

	plugin, err := client.UpdatePlugin(plugin)
	if err != nil {
		fmt.Println("Error updating Plugin: ", err.Error())
	}

	prettyPrintStruct(plugin)
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Filepath of a json representation of the resource to create.")

	// UpdatePlugin Command
	updateCmd.AddCommand(updatePluginCmd)
	// UpdatePlugin Flags
	updatePluginCmd.Flags().StringVarP(&plugin.ID, "id", "i", "", "ID of plugin to update")
}

var updateCmd = &cobra.Command{
	Use: "update [resource]",
}

var updatePluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Update a Plugin",
	Run:   updatePlugin,
}
