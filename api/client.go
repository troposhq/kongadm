package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
)

func buildRequest(method, path string, query url.Values, b io.Reader) *http.Request {
	if query == nil {
		query = url.Values{}
	}

	u := &url.URL{
		Scheme:   "http",
		Path:     path,
		RawQuery: query.Encode(),
	}

	r, err := http.NewRequest(method, u.String(), b)

	if err != nil {
		log.Fatal(err)
	}

	return r
}

func (c *KongAdminAPIClient) makeRequest(r *http.Request, d interface{}) error {
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
