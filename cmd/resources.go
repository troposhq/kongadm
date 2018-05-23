package cmd

import "github.com/troposhq/konga/api"

// resources.go contains object for storing values passed in through the CLI
//
// For example, if a user runs `konga create service --name hello --host google.com`
// the service variable below will have the values service.Name = "tropos-website" and service.Host = "troposhq.com"
//
// The same is true for any of the CLI methods including, Get, Update, and Delete commands. Any CLI flag values
// will be populated onto the appropriate object below.
//
// The objects are then used in the API calls to Kong

var service api.Service
var route api.Route
var plugin api.Plugin
var consumer api.Consumer

var filePath string
