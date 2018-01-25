package api

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
