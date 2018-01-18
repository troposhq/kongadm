package api

// ListAPIResult is the structure returned from the /apis route in Kong
type ListAPIResult struct {
	Total int   `json:"total"`
	Data  []API `json:"data"`
}

// API is an individual API object
type API struct {
	CreatedAt              int      `json:"created_at"`
	StripURI               bool     `json:"strip_uri"`
	ID                     string   `json:"id"`
	Name                   string   `json:"name"`
	UpstreamURL            string   `json:"upstream_url"`
	HTTPIfTerminated       bool     `json:"http_if_terminated"`
	HTTPSOnly              bool     `json:"https_only"`
	Retries                int      `json:"retries"`
	UpstreamSendTimeout    int      `json:"upstream_send_timeout"`
	UpstreamConnectTimeout int      `json:"upstream_connect_timeout"`
	UpstreamReadTimeout    int      `json:"upstream_read_timeout"`
	Methods                []string `json:"methods"`
	PreserveHost           bool     `json:"preserve_host"`
}

// ListAPI lists the API objects in Kong
func ListAPI() (results ListAPIResult, err error) {
	req := buildRequest("GET", apiURLBase+"apis", nil, nil)
	err = makeRequest(req, &results)
	return results, err
}

// GetAPI fetches an API by name or id
func GetAPI(nameOrID string) (r API, err error) {
	req := buildRequest("GET", apiURLBase+"apis/"+nameOrID, nil, nil)
	err = makeRequest(req, &r)
	return r, err
}
