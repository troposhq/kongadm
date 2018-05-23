package api

import (
	"bytes"
	"encoding/json"
)

// ListConsumersResult is the API response from Kong for fetching Consumers
type ListConsumersResult struct {
	Data []Consumer `json:"data"`
}

// Consumer is a Kong Consumer object
type Consumer struct {
	ID        string `json:"id,omitempty"`
	CustomID  string `json:"custom_id,omitempty"`
	Username  string `json:"username,omitempty"`
	CreatedAt int    `json:"created_at,omitempty"`
}

// ListConsumers fetches Consumers from the Kong API
func (c *KongAdminAPIClient) ListConsumers() (r ListConsumersResult, err error) {
	req := c.buildRequest("GET", "/consumers", nil, nil)
	err = c.makeRequest(req, &r)
	return r, err
}

// GetConsumer fetches an API by name or id
func (c *KongAdminAPIClient) GetConsumer(usernameOrID string) (r Consumer, err error) {
	req := c.buildRequest("GET", "/consumers/"+usernameOrID, nil, nil)
	err = c.makeRequest(req, &r)
	return r, err
}

// CreateConsumer creates a new Consumer object in Kong
func (c *KongAdminAPIClient) CreateConsumer(i Consumer) (r Consumer, err error) {
	b, err := json.Marshal(i)
	req := c.buildRequest("POST", "/consumers", nil, bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	err = c.makeRequest(req, &r)
	return r, err
}

// DeleteConsumer will delete a Consumer object in Kong
func (c *KongAdminAPIClient) DeleteConsumer(nameOrID string) (e error) {
	req := c.buildRequest("DELETE", "/consumers/"+nameOrID, nil, nil)
	e = c.makeRequest(req, nil)
	return e
}
