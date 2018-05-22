package api

import (
	"encoding/json"
	"fmt"
	"strings"
)

type AssociateConsumerRequest struct {
	Group string `json:"group"`
}

func (c *KongAdminAPIClient) AssociateConsumers(consumer string, b AssociateConsumerRequest) (r *string, err error) {
	body, _ := json.Marshal(b)
	fmt.Println(string(body))
	req := buildRequest("POST", c.APIURLBase+"/consumers/"+consumer+"/acls", nil, strings.NewReader(string(body)))
	req.Header.Add("Content-Type", "application/json")
	err = c.makeRequest(req, &r)
	return r, err
}
