package api

import (
	"fmt"
)

// ListAPI lists the API objects in Kong
func ListAPI() (sr string, err error) {
	fmt.Println("Listing API Objects...")

	return "results", nil
}
