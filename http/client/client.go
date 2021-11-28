package client

import (
	"context"
	"io"
	"net"
	"net/http"
)

// A Client represents HTTP client to send request to server.
type Client struct {
	// config specifies configuration for Client.
	config *config

	// client specifies an instance of http.Client.
	client *http.Client
}

// Header represents a HTTP header
type Header struct {
	Key   string
	Value string
}

func WithHeader(key, value string) Header {
	return Header{key, value}
}

// Get issues a GET request to the URL.
func (c *Client) Get(ctx context.Context, url string, headers ...Header) (resp *http.Response, err error) {
	req, err := makeRequest(ctx, url, http.MethodGet, nil, headers...)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

// Post issues a POST request to the URL and sends a body to server.
func (c *Client) Post(ctx context.Context, url string, body io.Reader, headers ...Header) (resp *http.Response, err error) {
	req, err := makeRequest(ctx, url, http.MethodPost, body, headers...)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	for _, header := range c.config.DefaultHeaders {
		req.Header.Add(header.Key, header.Value)
	}

	return c.client.Do(req)
}

// New returns a new instances of Client. It sets the necessary defaults if configuration is nil.
// It will also set any missing fields with defaults variables.
func New(opts ...Option) *Client {
	c := setDefaults()

	for _, opt := range opts {
		opt(c)
	}

	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   c.Timeout,
			KeepAlive: c.KeepAlive,
			DualStack: true,
		}).Dial,
		ForceAttemptHTTP2:   c.EnableHTTP2,
		MaxIdleConns:        c.MaxIdleConns,
		IdleConnTimeout:     c.IdleConnTimeout,
		TLSHandshakeTimeout: c.TLSHandshakeTimeout,
	}

	if c.TLSConfig != nil {
		transport.TLSClientConfig = c.TLSConfig
	}

	return &Client{
		config: c,
		client: &http.Client{
			Transport: transport,
		},
	}
}
