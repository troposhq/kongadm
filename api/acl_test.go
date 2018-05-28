package api

import (
	"fmt"
	"testing"
)

func TestAssociateConsumers(t *testing.T) {
	client := New("localhost:8001")
	res, e := client.ACLAssociateConsumer("a9232e57-bda4-44d8-8c79-b0050f8ec217", "admins")
	if e != nil {
		fmt.Println(e)
		t.Error("Error associating consumer to acl")
	}

	if res.Group != "admins" {
		t.Error("Error associating consumer to acl")
	}
}
