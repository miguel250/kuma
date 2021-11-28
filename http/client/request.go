package client

import (
	"context"
	"io"
	"net/http"
)

// makeRequest returns a http.request with a context.
func makeRequest(ctx context.Context, url, method string, body io.Reader, headers ...Header) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	for _, header := range headers {
		req.Header.Add(header.Key, header.Value)
	}

	return req, nil
}
