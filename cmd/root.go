package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/troposhq/kongadm/api"
)

type PluginConfig struct {
	Options map[string]interface{}
}

func (c *PluginConfig) String() string {
	return fmt.Sprintln(c.Options)
}

func (c *PluginConfig) Set(s string) error {
	if c.Options == nil {
		c.Options = make(map[string]interface{})
	}

	kv := strings.Split(s, "=")
	if len(kv) != 2 {
		return errors.New("Specify --config in the format key=value")
	}

	k := kv[0]
	v := kv[1]
	c.Options[k] = v
	return nil
}

func (c *PluginConfig) Type() string {
	return "PluginConfig"
}

// These contains objects for storing values passed in through the CLI
//
// For example, if a user runs `kongadm create service --name hello --host google.com`
// the service variable below will have the values service.Name = "tropos-website" and service.Host = "troposhq.com"
//
// The same is true for any of the CLI methods including, Get, Update, and Delete commands. Any CLI flag values
// will be populated onto the appropriate object below.
//
// The objects are then used in the API calls to Kong

var service api.Service
var route api.Route
var consumer api.Consumer
var plugin api.Plugin
var pluginConfig PluginConfig

var filePath string

var client *api.KongAdminAPIClient

var rootCmd = &cobra.Command{
	Use:   "kongadm",
	Short: "A CLI tool for interacting with the Kong API",
	Long:  "A CLI tool for interacting with the Kong API",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		apiURLBase := viper.GetString("url")
		client = api.New(apiURLBase, nil)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.kongadm")
	// Create the config file if it does not exist
	if _, err := os.Stat("$HOME/.kongadm/config"); os.IsNotExist(err) {
		usr, _ := user.Current()
		dir := usr.HomeDir
		os.MkdirAll(dir+"/.kongadm", 0755)
		_, err = os.OpenFile(dir+"/.kongadm/config.yml", os.O_RDONLY|os.O_CREATE, 0644)
	}

	// Load the config file
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error reading config file: %s \n", err)
		os.Exit(1)
	}
}

// Execute is the entrypoint for the CLI called from the main function
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
