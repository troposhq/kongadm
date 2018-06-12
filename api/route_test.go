package api

import (
	"fmt"
	"testing"
)

func TestAddRoute(t *testing.T) {
	service, _ := client.AddService(Service{
		URL: "http://localhost:8001",
	})
	r, err := client.AddRoute(Route{
		Paths:   []string{"/test"},
		Service: service,
	})

	if err != nil {
		fmt.Print(err.Error())
	}

	client.DeleteRoute(r.ID)
	client.DeleteService(service.ID)
}
