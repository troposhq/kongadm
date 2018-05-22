package api

import (
	"fmt"
	"testing"
)

func TestAssociateConsumers(t *testing.T) {
	client := New("localhost:8001")
	_, e := client.AssociateConsumer("d334c28c-11ab-4c7e-8464-50ef660e82fd", "admins")
	fmt.Println(e)
}
