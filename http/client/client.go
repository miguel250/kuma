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
	config *Config

	// client specifies an instance of http.Client.
	client *http.Client
}

// Get issues a GET request to the URL.
func (c *Client) Get(ctx context.Context, url string) (resp *http.Response, err error) {
	r := request{
		ctx:         ctx,
		method:      http.MethodGet,
		url:         url,
		contentType: "",
		body:        nil,
	}
	return r.makeRequest(c)
}

// Post issues a POST request to the URL and sends a body to server.
func (c *Client) Post(ctx context.Context, url, contentType string, body io.Reader) (resp *http.Response, err error) {
	r := request{
		ctx:         ctx,
		method:      http.MethodPost,
		url:         url,
		contentType: contentType,
		body:        body,
	}
	return r.makeRequest(c)
}

// New returns a new instances of Client. It sets the necessary defaults if configuration is nil.
// It will also set any missing fields with defaults variables.
func New(config *Config) *Client {

	if config == nil {
		config = defaultConfig
	}

	setDefaults(config)

	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   config.Timeout,
			KeepAlive: config.KeepAlive,
			DualStack: true,
		}).Dial,
		ForceAttemptHTTP2:   config.EnableHTTP2,
		MaxIdleConns:        config.MaxIdleConns,
		IdleConnTimeout:     config.IdleConnTimeout,
		TLSHandshakeTimeout: config.TLSHandshakeTimeout,
	}

	if config.TLSConfig != nil {
		transport.TLSClientConfig = config.TLSConfig
	}

	return &Client{
		config: config,
		client: &http.Client{
			Transport: transport,
		},
	}
}
