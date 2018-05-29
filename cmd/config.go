package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

func setOption(cmd *cobra.Command, args []string) {
	var err error
	key, val := args[0], args[1]

	viper.Set(key, val)
	err = viper.WriteConfig()
	if err != nil {
		fmt.Println("Error setting configuration option: ", err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(setCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure options for kongadm",
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a configuration option",
	Args:  cobra.ExactArgs(2),
	Run:   setOption,
}
