package api

import (
	"bytes"
	"encoding/json"
)

// {"group":"group1","created_at":1527531744000,"id":"b82d1271-52b9-41d7-b1f0-e68cd7fc4c1b","consumer_id":"7b8b2059-3707-4fb6-9241-fdadd800a770"}

type ACLAssociateConsumerRequest struct {
	Group string `json:"group"`
}

type ACLAssociateConsumerResponse struct {
	ID         string `json:"id"`
	Group      string `json:"group"`
	ConsumerID string `json:"consumer_id"`
	CreatedAt  int64  `json:"created_at"`
}

// AssociateConsumer will add a consumer to an ACL group
func (c KongAdminAPIClient) ACLAssociateConsumer(consumer string, group string) (r ACLAssociateConsumerResponse, err error) {
	i, _ := json.Marshal(ACLAssociateConsumerRequest{Group: group})
	req := c.buildRequest("POST", "/consumers/"+consumer+"/acls", nil, bytes.NewReader(i))
	req.Header.Add("Content-Type", "application/json")
	err = c.makeRequest(req, &r)
	return r, err
}
