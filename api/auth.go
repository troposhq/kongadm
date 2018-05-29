package api

import "net/http"

type KongAdminAPIAuthStrategy interface {
	configure(req *http.Request)
}

type BasicAuthStrategy struct {
	Username string
	Password string
}

func (s BasicAuthStrategy) configure(req *http.Request) {
	req.SetBasicAuth(s.Username, s.Password)
}

type TokenStrategy struct {
	Token string
}

func (s TokenStrategy) configure(req *http.Request) {
	req.Header.Add("Authorization", "Bearer "+s.Token)
}
