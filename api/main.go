package api

import "net/http"

type KongAdminAPIClient struct {
	APIURLBase string
	client     *http.Client
}

func New(baseUrl string) *KongAdminAPIClient {
	return &KongAdminAPIClient{
		APIURLBase: baseUrl,
		client:     &http.Client{},
	}
}
