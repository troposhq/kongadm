package api

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
func ListPlugins() (results ListPluginsResult, err error) {
	req := buildRequest("GET", APIURLBase+"/plugins", nil, nil)
	err = makeRequest(req, &results)
	return results, err
}
