package modpacksch

import (
	"encoding/json"
	"net/http"
)

func (c *Client) newRequest(method string, path string, body interface{}) (*http.Request, error) {
	absoluteUrl := ApiRoot + path

	// todo: implement sending a body
	request, err := http.NewRequest(method, absoluteUrl, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "go-modpacksch/0.1.0")

	return request, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
