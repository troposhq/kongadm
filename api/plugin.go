package api

import (
	"bytes"
	"encoding/json"
)

// ListPluginsResult is the structure returned from the /plugins route in Kong
type ListPluginsResult struct {
	Data []Plugin `json:"data"`
}

// Plugin is an individual Plugin object
type Plugin struct {
	ID         string                 `json:"id,omitempty"`
	ServiceID  string                 `json:"service_id,omitempty"`
	ConsumerID string                 `json:"consumer_id,omitempty"`
	RouteID    string                 `json:"route_id,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Enabled    bool                   `json:"enabled,omitempty"`
	Config     map[string]interface{} `json:"config,omitempty"`
	CreatedAt  int                    `json:"created_at,omitempty"`
}

// ListPlugins lists the Plugin objects in Kong
func (c *KongAdminAPIClient) ListPlugins() (results ListPluginsResult, err error) {
	req := c.buildRequest("GET", "/plugins", nil, nil)
	err = c.makeRequest(req, &results)
	return results, err
}

// AddPlugin adds a new Plugin object in Kong
func (c *KongAdminAPIClient) AddPlugin(p Plugin) (r Plugin, err error) {
	b, err := json.Marshal(p)
	req := c.buildRequest("POST", "/plugins", nil, bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	err = c.makeRequest(req, &r)
	return r, err
}
