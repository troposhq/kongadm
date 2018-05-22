package api

import "strings"

// ListPluginsResult is the structure returned from the /plugins route in Kong
type ListPluginsResult struct {
	Data []Plugin `json:"data"`
}

// Plugin is an individual Plugin object
type Plugin struct {
	ID        string `json:"id"`
	ServiceID string `json:"service_id"`
	Name      string `json:"name"`
	Enabled   bool   `json:"enabled"`
	CreatedAt int    `json:"created_at"`
}

// ListPlugins lists the Plugin objects in Kong
func (c *KongAdminAPIClient) ListPlugins() (results ListPluginsResult, err error) {
	req := c.buildRequest("GET", "/plugins", nil, nil)
	err = c.makeRequest(req, &results)
	return results, err
}

// CreatePlugin adds a new Plugin object in Kong
func (c *KongAdminAPIClient) CreatePlugin(b string, serviceId string, routeId string) (r Plugin, err error) {
	var url string
	if serviceId != "" {
		url = c.APIURL + "/services/" + serviceId + "/plugins"
	} else if routeId != "" {
		url = c.APIURL + "/routes/" + routeId + "/plugins"
	} else {
		url = c.APIURL + "/plugins"
	}
	req := c.buildRequest("POST", url, nil, strings.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	err = c.makeRequest(req, &r)
	return r, err
}
