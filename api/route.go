package api

import "strings"

// ListRoutesResult is the structure returned from the /routes route in Kong
type ListRoutesResult struct {
	Data []Route `json:"data"`
}

// Route is an individual Route object
type Route struct {
	ID            string   `json:"id"`
	CreatedAt     int      `json:"created_at"`
	UpdatedAt     int      `json:"updated_at"`
	Protocols     []string `json:"protocols"`
	Methods       []string `json:"methods"`
	Hosts         []string `json:"hosts"`
	Paths         []string `json:"paths"`
	RegexPriority int      `json:"regex_priority"`
	PreserveHost  bool     `json:"preserve_host"`
}

// CreateRoute create a new Kong Route object
func (c *KongAdminAPIClient) CreateRoute(b string) (r Route, err error) {
	req := c.buildRequest("POST", "/routes", nil, strings.NewReader(b))
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
