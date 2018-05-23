package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
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

var rootCmd = &cobra.Command{
	Use:   "kongadm",
	Short: "A CLI tool for interacting with the Kong API",
	Long:  "A CLI tool for interacting with the Kong API",
}

var client *api.KongAdminAPIClient

func init() {
	apiURLBase := rootCmd.PersistentFlags().String("url", "localhost:8001", "URL for the Kong Admin API")
	client = api.New(*apiURLBase)
}

// Execute is the entrypoint for the CLI called from the main function
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
