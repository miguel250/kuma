package client

import (
	"crypto/tls"
	"time"
)

type Option func(*config)

// WithMaxIdleConns controls maximum of idle connections using keep-alive.
func WithMaxIdleConns(max int) Option {
	return func(c *config) {
		c.MaxIdleConns = max
	}
}

// WithTimeout specifies the time limit for a request including getting a response from server.
func WithTimeout(d time.Duration) Option {
	return func(c *config) {
		c.Timeout = d
	}
}

// WithKeepAlive specifies the interval between keep-alive probes.
func WithKeepAlive(d time.Duration) Option {
	return func(c *config) {
		c.KeepAlive = d
	}
}

// WithIdleConnTimeout is the maximum amount of time a keep-alive connection is kept open.
func WithIdleConnTimeout(d time.Duration) Option {
	return func(c *config) {
		c.IdleConnTimeout = d
	}
}

// WithTLSHandshakeTimeout specifies the time limit to wait for a TLS handshake.
func WithTLSHandshakeTimeout(d time.Duration) Option {
	return func(c *config) {
		c.TLSHandshakeTimeout = d
	}
}

// WithExpectContinueTimeout is the maximum amount of time to wait for a server first response.
func WithExpectContinueTimeout(d time.Duration) Option {
	return func(c *config) {
		c.ExpectContinueTimeout = d
	}
}

// WithTLSConfig specifies necessary TLS configuration.
func WithTLSConfig(tlsConf *tls.Config) Option {
	return func(c *config) {
		c.TLSConfig = tlsConf
	}
}

// WithEnableHTTP2 controls whether to use HTTP2
func WithEnableHTTP2(b bool) Option {
	return func(c *config) {
		c.EnableHTTP2 = b
	}
}
