package modpacksch

import (
	"net/http"
)

var (
	ApiRoot = "https://api.modpacks.ch"
)

type Client struct {
	httpClient *http.Client

	Packs *PackService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{httpClient: httpClient}
	c.Packs = &PackService{client: c}
	return c
}
