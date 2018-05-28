package api

import (
	"fmt"
	"testing"
)

var client = New("localhost:8001")
var service Service

func TestCreateService(t *testing.T) {
	s, err := client.AddService(Service{
		URL: "http://localhost:8001",
	})

	if err != nil {
		fmt.Print(err.Error())
	}

	service = s
}

func TestListServices(t *testing.T) {
	services, _ := client.ListServices()
	if len(services.Data) != 1 {
		t.Error("Expected services to be of length 1")
	}
}

func TestGetService(t *testing.T) {
	s, _ := client.GetService(service.ID)
	if s.ID != service.ID {
		t.Error("Expected to get the existing service")
	}
}

func TestUpdateService(t *testing.T) {
	service.Path = "/fooba"
	s, err := client.UpdateService(service)
	if err != nil {
		t.Error(err.Error())
	}
	if s.Path != service.Path {
		t.Error("Expected to update Service")
	}
}

func TestDeleteService(t *testing.T) {
	err := client.DeleteService(service.ID)
	if err != nil {
		fmt.Println(err.Error())
	}
}
