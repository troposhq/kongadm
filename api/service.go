package api

import (
	"bytes"
	"encoding/json"
)

// ListServicesResult is the structure returned from the /apis route in Kong
type ListServicesResult struct {
	Data []Service `json:"data"`
}

// Service is an individual API object
type Service struct {
	ID             string `json:"id,omitempty"`
	CreatedAt      int    `json:"created_at,omitempty"`
	UpdatedAt      int    `json:"updated_at,omitempty"`
	ConnectTimeout int    `json:"connect_timeout,omitempty"`
	Protocol       string `json:"protocol,omitempty"`
	Host           string `json:"host,omitempty"`
	Port           int    `json:"port,omitempty"`
	Path           string `json:"path,omitempty"`
	Name           string `json:"name,omitempty"`
	Retries        int    `json:"retries,omitempty"`
	ReadTimeout    int    `json:"read_timeout,omitempty"`
	WriteTimeout   int    `json:"write_timeout,omitempty"`
	URL            string `json:"url,omitempty"`
}

// AddService creates a new Service object in Kong
func (c KongAdminAPIClient) AddService(s Service) (r Service, err error) {
	b, err := json.Marshal(s)
	req := c.buildRequest("POST", "/services", nil, bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	err = c.makeRequest(req, &r)
	return r, err
}

// ListServices lists Service objects in Kong
func (c KongAdminAPIClient) ListServices() (results ListServicesResult, err error) {
	req := c.buildRequest("GET", "/services", nil, nil)
	err = c.makeRequest(req, &results)
	return results, err
}

// GetService fetches a Service by name or id
func (c KongAdminAPIClient) GetService(s string) (r Service, err error) {
	req := c.buildRequest("GET", "/services/"+s, nil, nil)
	err = c.makeRequest(req, &r)
	return r, err
}

// UpdateService will update a Service object in Kong
func (c KongAdminAPIClient) UpdateService(s Service) (r Service, err error) {
	nameOrID := s.Name
	if nameOrID == "" {
		nameOrID = s.ID
	}

	b, err := json.Marshal(s)
	req := c.buildRequest("PATCH", "/services/"+nameOrID, nil, bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	err = c.makeRequest(req, &r)
	return r, err
}

// DeleteService will delete a Service object in Kong
func (c KongAdminAPIClient) DeleteService(s string) error {
	req := c.buildRequest("DELETE", "/services/"+s, nil, nil)
	err := c.makeRequest(req, nil)
	return err
}
