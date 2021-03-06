// Package modpacksch provides a client for using the modpacks.ch
// API.
package modpacksch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL   = "https://api.modpacks.ch/"
	defaultUserAgent = "go-modpacksch"

	StatusSuccess = "success"
	StatusError = "error"
)

// A client manages communication with the modpacks.ch API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests. Defaults to the production
	// modpacks.ch API, but can be set to a domain endpoint
	// to use with other instances. BaseURL should always be
	// set with a trailing slash.
	BaseURL *url.URL

	// User Agent used when communicating with the modpacks.ch API.
	UserAgent string

	// Services used for accessing different parts of the modpacks.ch
	// API.
	Packs *PackService
	Tags *TagService
	CurseForge *CurseForgeService
	Mods *ModService
}

type service struct {
	client *Client
}

// NewClient returns a new modpacks.ch API client. If a nil client is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client: httpClient,
		BaseURL: baseURL,
		UserAgent: defaultUserAgent,
	}
	c.Packs = &PackService{client: c}
	c.Tags = &TagService{client: c}
	c.CurseForge = &CurseForgeService{client: c}
	c.Mods = &ModService{client: c}
	return c
}

// NewRequest creates an API request. A relative URL can be provided
// in urlStr, in which case it is resolved to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
// If specified, the value pointed to by the body is JSON encoded and
// included as the request body.
func (c *Client) NewRequest(method string, urlStr string, body interface{}) (*http.Request, error) {
	// Resolve absolute URL
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// Encode body as JSON
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// Create the request
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// Set request headers
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response
// is JSON decoded and stored in the value pointed to by v.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}

func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{
		Response: r,
	}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			return err
		}
	}
	// Repopulate body so that it can proceed to be read again, should
	// there be no error.
	r.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	// If we have no error information, assume success
	if errorResponse.Status == "" && errorResponse.Message == "" {
		return nil
	}

	// If we're a "success", then proceed
	if errorResponse.Status == StatusSuccess {
		return nil
	}

	return errorResponse
}

type ErrorResponse struct {
	Response *http.Response

	Status string `json:"status"`
	Message string `json:"message"`
}

var _ error = (*ErrorResponse)(nil)

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %v - %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Status, r.Message)
}
