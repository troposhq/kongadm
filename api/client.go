package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

func (c KongAdminAPIClient) buildRequest(method, path string, query url.Values, b io.Reader) *http.Request {
	if query == nil {
		query = url.Values{}
	}

	u, err := url.Parse(c.APIURL)
	if err != nil {
		fmt.Println("Error parsing Admin API URL")
		fmt.Println(err.Error())
		return nil
	}

	u.Path = filepath.Join(u.Path, path)
	u.RawQuery = query.Encode()

	r, err := http.NewRequest(method, u.String(), b)
	if c.AuthStrategy != nil {
		c.AuthStrategy.configure(r)
	}

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return r
}

func (c KongAdminAPIClient) makeRequest(r *http.Request, d interface{}) error {
	res, err := c.client.Do(r)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK &&
		res.StatusCode != http.StatusCreated &&
		res.StatusCode != http.StatusAccepted &&
		res.StatusCode != http.StatusNoContent {
		return errors.New(res.Status)
	}

	if d != nil {
		return json.NewDecoder(res.Body).Decode(d)
	}

	return nil
}

type KongAdminAPIClient struct {
	APIURL       string
	client       *http.Client
	AuthStrategy KongAdminAPIAuthStrategy
}

// New creates a new instance of the KongAdminAPIClient
func New(baseURL string, authStrategy KongAdminAPIAuthStrategy) *KongAdminAPIClient {
	// default to http if not included in URL
	if !strings.HasPrefix(baseURL, "http") {
		baseURL = "http://" + baseURL
	}
	return &KongAdminAPIClient{
		APIURL:       baseURL,
		client:       &http.Client{},
		AuthStrategy: authStrategy,
	}
}
