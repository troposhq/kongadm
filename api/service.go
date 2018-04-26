package api

import "strings"

// ListServicesResult is the structure returned from the /apis route in Kong
type ListServicesResult struct {
	Total int       `json:"total"`
	Data  []Service `json:"data"`
}

// Service is an individual API object
type Service struct {
	ID             string `json:"id"`
	CreatedAt      int    `json:"created_at"`
	UpdatedAt      int    `json:"updated_at"`
	ConnectTimeout int    `json:"connect_timeout"`
	Protocol       string `json:"protocol"`
	Host           string `json:"host"`
	Port           int    `json:"port"`
	Path           string `json:"path"`
	Name           string `json:"name"`
	Retries        int    `json:"retries"`
	ReadTimeout    int    `json:"read_timeout"`
	WriteTimeout   int    `json:"write_timeout"`
}

// ListServices lists the Service objects in Kong
func ListServices() (results ListServicesResult, err error) {
	req := buildRequest("GET", APIURLBase+"/services", nil, nil)
	err = makeRequest(req, &results)
	return results, err
}

// GetService fetches an Service by name or id
func GetService(nameOrID string) (r Service, err error) {
	req := buildRequest("GET", APIURLBase+"/services/"+nameOrID, nil, nil)
	err = makeRequest(req, &r)
	return r, err
}

// CreateService creates a new Service object in Kong
func CreateService(b string) (r Service, err error) {
	req := buildRequest("POST", APIURLBase+"/services", nil, strings.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	err = makeRequest(req, &r)
	return r, err
}

// DeleteService will delete an Service object in Kong
func DeleteService(nameOrID string) (e error) {
	req := buildRequest("DELETE", APIURLBase+"/services/"+nameOrID, nil, nil)
	e = makeRequest(req, nil)
	return e
}
