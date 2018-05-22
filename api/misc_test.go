package api

import (
	"fmt"
	"testing"
)

func TestAssociateConsumers(t *testing.T) {
	client := New("localhost:8001")
	req := AssociateConsumerRequest{
		Group: "admins",
	}
	_, e := client.AssociateConsumers("d334c28c-11ab-4c7e-8464-50ef660e82fd", req)
	fmt.Println(e)
}
