package client

import (
	"crypto/tls"
	"time"
)

// Config represents the necessary configuration for a HTTP client.
type config struct {
	// MaxIdleConns controls maximum of idle connections using keep-alive.
	MaxIdleConns int

	// Timeout specifies the time limit for a request including getting a response from server.
	Timeout time.Duration

	// KeepAlive specifies the interval between keep-alive probes.
	KeepAlive time.Duration

	// IdleConnTimeout is the maximum amount of time a keep-alive connection is kept open.
	IdleConnTimeout time.Duration

	// TLSHandshakeTimeout specifies the time limit to wait for a TLS handshake.
	TLSHandshakeTimeout time.Duration

	// ExpectContinueTimeout is the maximum amount of time to wait for a server first response.
	ExpectContinueTimeout time.Duration

	// TLSConfig specifies necessary TLS configuration.
	TLSConfig *tls.Config

	//EnableHTTP2 controls whether to use HTTP2
	EnableHTTP2 bool
}

// Default configurations used for the HTTP client
const (
	maxIdleConns          = 100
	timeout               = 30 * time.Second
	keepAlive             = 30 * time.Second
	idleConnTimeout       = 90 * time.Second
	tlsHandshakeTimeout   = 10 * time.Second
	expectContinueTimeout = 1 * time.Second
)

// setDefaults adds missing configuration past by the caller by using defaultConfig
func setDefaults() *config {
	return &config{
		MaxIdleConns:          maxIdleConns,
		Timeout:               timeout,
		KeepAlive:             keepAlive,
		IdleConnTimeout:       idleConnTimeout,
		TLSHandshakeTimeout:   tlsHandshakeTimeout,
		ExpectContinueTimeout: expectContinueTimeout,
	}
}
