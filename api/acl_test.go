package api

import (
	"fmt"
	"testing"
)

func TestAssociateConsumers(t *testing.T) {
	client := New("localhost:8001", nil)
	consumer, _ := client.CreateConsumer(Consumer{Username: "alois"})
	res, e := client.ACLAssociateConsumer(consumer.ID, "admins")
	if e != nil {
		fmt.Println(e)
		t.Error("Error associating consumer to acl")
	}

	if res.Group != "admins" {
		t.Error("Error associating consumer to acl")
	}

	client.DeleteConsumer(consumer.ID)
}
