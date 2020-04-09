// Package modpacksch provides a client for using the
// modpacks.ch API.
package modpacksch

import (
	"net/http"
)

var (
	ApiRoot = "https://api.modpacks.ch"
)

// A client manages communication with the modpacks.ch API.
type Client struct {
	httpClient *http.Client

	Packs *PackService
}

type service struct {
	client *Client
}

// NewClient returns a new modpacks.ch API client. If a nil
// httpClient is provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{httpClient: httpClient}
	c.Packs = &PackService{client: c}
	return c
}
