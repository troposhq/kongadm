# kongadmdm

# Table of Contents

- [Overview](#overview)
- [Installing](#installing)
- [Getting Started](#getting-started)
  * [Services](#services)
- [License](#license)

# Overview

Konga is a CLI that provides a simple interface for interacting with the Kong Admin API.

# Installation

Download the binary from our Releases page.

# Getting Started

## Usage

```
Usage:
  kongadm [command]

Available Commands:
  add
  config      Configure options for kongadm
  create
  delete
  get
  help        Help about any command
  list
  update

Flags:
  -h, --help   help for kongadm

Use "kongadm [command] --help" for more information about a command.
```

## Services

### Add Service
```
Usage:
  kongadm add service [flags]

Flags:
      --connect-timeout int   The timeout in milliseconds for establishing a connection to the upstream servere (default 60000)
  -h, --help                  help for service
      --host string           The host of the upstream server
      --name string           The Service name
      --path string           The path to be used in requests to the upstream server. Empty by default
      --port int              The upstream server port (default 80)
      --protocol string       The protocol used to communicate with the upstream. It can be one of http (default) or https (default "http")
      --read-timeout int      The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server (default 60000)
      --retries int           The number of retries to execute upon failure to proxy. The default is 5 (default 5)
      --url string            Shorthand attribute to set protocol, host, port and path at once. This attribute is write-only (the Admin API never "returns" the url)
      --write-timeout int     The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server (default 60000)

Global Flags:
  -f, --file string   Filepath of a json representation of the resource to create.
```

#### Example
```
# Using the shorthand flag
kongadm add service --url https://www.example.com

# Using individual flags
kongadm add service --host example.com --protocol https
```

### Retrieve Service
```
Usage:
  kongadm add service [flags]

Flags:
      --connect-timeout int   The timeout in milliseconds for establishing a connection to the upstream servere (default 60000)
  -h, --help                  help for service
      --host string           The host of the upstream server
      --name string           The Service name
      --path string           The path to be used in requests to the upstream server. Empty by default
      --port int              The upstream server port (default 80)
      --protocol string       The protocol used to communicate with the upstream. It can be one of http (default) or https (default "http")
      --read-timeout int      The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server (default 60000)
      --retries int           The number of retries to execute upon failure to proxy. The default is 5 (default 5)
      --url string            Shorthand attribute to set protocol, host, port and path at once. This attribute is write-only (the Admin API never "returns" the url)
      --write-timeout int     The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server (default 60000)

Global Flags:
  -f, --file string   Filepath of a json representation of the resource to create.
```

#### Example
```
# Using the shorthand flag
kongadm add service --url https://www.example.com

# Using individual flags
kongadm add service --host example.com --protocol https
```
