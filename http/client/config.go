package client

import (
	"crypto/tls"
	"time"
)

// Config represents the necessary configuration for a HTTP client.
type Config struct {
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

// defaultConfig is the default configuration used for the HTTP client
var defaultConfig = &Config{
	MaxIdleConns:          100,
	Timeout:               30 * time.Second,
	KeepAlive:             30 * time.Second,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	EnableHTTP2:           true,
}

// setDefaults adds missing configuration past by the caller by using defaultConfig
func setDefaults(config *Config) {

	if config.MaxIdleConns == 0 {
		config.MaxIdleConns = defaultConfig.MaxIdleConns
	}

	if config.Timeout == 0 {
		config.Timeout = defaultConfig.Timeout
	}

	if config.KeepAlive == 0 {
		config.KeepAlive = defaultConfig.KeepAlive
	}

	if config.IdleConnTimeout == 0 {
		config.IdleConnTimeout = defaultConfig.IdleConnTimeout
	}

	if config.TLSHandshakeTimeout == 0 {
		config.TLSHandshakeTimeout = defaultConfig.TLSHandshakeTimeout
	}

	if config.ExpectContinueTimeout == 0 {
		config.ExpectContinueTimeout = defaultConfig.ExpectContinueTimeout
	}
}
