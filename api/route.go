package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// ListRoutesResult is the structure returned from the /routes path in Kong
type ListRoutesResult struct {
	Data []Route `json:"data"`
}

// Route is an individual Route object
type Route struct {
	ID            string   `json:"id,omitempty"`
	CreatedAt     int      `json:"created_at,omitempty"`
	UpdatedAt     int      `json:"updated_at,omitempty"`
	Protocols     []string `json:"protocols,omitempty"`
	Methods       []string `json:"methods,omitempty"`
	Hosts         []string `json:"hosts,omitempty"`
	Paths         []string `json:"paths,omitempty"`
	RegexPriority int      `json:"regex_priority,omitempty"`
	PreserveHost  bool     `json:"preserve_host,omitempty"`
	StripPath     bool     `json:"strip_path,omitempty"`
	Service       Service  `json:"service,omitempty"`
}

// AddRoute creates a new Kong Route object
func (c *KongAdminAPIClient) AddRoute(route Route) (r Route, err error) {
	b, err := json.Marshal(route)
	if err != nil {
		fmt.Println("Error marshaling Route: " + err.Error())
		return r, err
	}

	req := c.buildRequest("POST", "/routes", nil, bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	err = c.makeRequest(req, &r)
	return r, err
}

// ListRoutes lists the Route objects in Kong
func (c *KongAdminAPIClient) ListRoutes() (results ListRoutesResult, err error) {
	req := c.buildRequest("GET", "/routes", nil, nil)
	err = c.makeRequest(req, &results)
	return results, err
}
