package client

import (
	"context"
	"io"
	"net/http"
)

// A request represents a HTTP request. It is use create a http Request with context.
type request struct {
	// ctx is context to be use while create a http.request.
	ctx context.Context

	// method is the http method to be send to server.
	method string

	// url specifies which server to connect to and send request.
	url string

	// contentType specifies the content-type of request body.
	contentType string

	// body specifies content to be send to server from client.
	body io.Reader
}

// makeRequest make a request to server and returns an instance of http.Response.
func (r request) makeRequest(c *Client) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(r.ctx, r.method, r.url, r.body)
	if err != nil {
		return nil, err
	}

	if r.contentType != "" {
		req.Header.Set("Content-Type", r.contentType)
	}

	return c.client.Do(req)
}
