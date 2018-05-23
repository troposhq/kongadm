package api

import "strings"

// AssociateConsumer will add a consumer to an ACL group
func (c *KongAdminAPIClient) AssociateConsumer(consumer string, group string) (r *string, err error) {
	req := c.buildRequest("POST", "/consumers/"+consumer+"/acls", nil, strings.NewReader(group))
	req.Header.Add("Content-Type", "application/json")
	err = c.makeRequest(req, &r)
	return nil, nil
}
