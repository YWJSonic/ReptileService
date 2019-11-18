package handlehttp

import "net/http"

type httpClient struct {
	Client *http.Client
}

var clientConnect *httpClient
