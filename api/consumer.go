package api

import "strings"

// ListConsumersResult is the API response from Kong for fetching Consumers
type ListConsumersResult struct {
	Total int        `json:"total"`
	Data  []Consumer `json:"data"`
}

// Consumer is a Kong Consumer object
type Consumer struct {
	ID        string `json:"id"`
	CustomID  string `json:"custom_id"`
	CreatedAt int    `json:"created_at"`
}

// ListConsumers fetches Consumers from the Kong API
func (c *KongAdminAPIClient) ListConsumers() (r ListConsumersResult, err error) {
	req := buildRequest("GET", c.APIURLBase+"/consumers", nil, nil)
	err = c.makeRequest(req, &r)
	return r, err
}

// GetConsumer fetches an API by name or id
func (c *KongAdminAPIClient) GetConsumer(usernameOrID string) (r Consumer, err error) {
	req := buildRequest("GET", c.APIURLBase+"/consumers/"+usernameOrID, nil, nil)
	err = c.makeRequest(req, &r)
	return r, err
}

// CreateConsumer creates a new Consumer object in Kong
func (c *KongAdminAPIClient) CreateConsumer(b string) (r Consumer, err error) {
	req := buildRequest("POST", c.APIURLBase+"/consumers", nil, strings.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	err = c.makeRequest(req, &r)
	return r, err
}

// DeleteConsumer will delete a Consumer object in Kong
func (c *KongAdminAPIClient) DeleteConsumer(nameOrID string) (e error) {
	req := buildRequest("DELETE", c.APIURLBase+"/consumers/"+nameOrID, nil, nil)
	e = c.makeRequest(req, nil)
	return e
}
