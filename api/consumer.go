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
func ListConsumers() (r ListConsumersResult, err error) {
	req := buildRequest("GET", APIURLBase+"/consumers", nil, nil)
	err = makeRequest(req, &r)
	return r, err
}

// GetConsumer fetches an API by name or id
func GetConsumer(usernameOrID string) (r API, err error) {
	req := buildRequest("GET", APIURLBase+"/consumers/"+usernameOrID, nil, nil)
	err = makeRequest(req, &r)
	return r, err
}

// CreateConsumer creates a new Consumer object in Kong
func CreateConsumer(b string) (r API, err error) {
	req := buildRequest("POST", APIURLBase+"/consumers", nil, strings.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	err = makeRequest(req, &r)
	return r, err
}

// DeleteConsumer will delete a Consumer object in Kong
func DeleteConsumer(nameOrID string) (e error) {
	req := buildRequest("DELETE", APIURLBase+"/consumers/"+nameOrID, nil, nil)
	e = makeRequest(req, nil)
	return e
}
