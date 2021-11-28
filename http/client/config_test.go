package client

import (
	"crypto/tls"
	"testing"
	"time"
)

func TestSetDefaults(t *testing.T) {
	c := setDefaults()

	if c.MaxIdleConns != maxIdleConns {
		t.Errorf("Default MaxIdleConns don't match got ('%d'), want ('%d')", c.MaxIdleConns, maxIdleConns)
	}

	if c.Timeout != timeout {
		t.Errorf("Default Timeout don't match got ('%d'), want ('%d')", c.Timeout, timeout)
	}
}

func TestWithMaxIdleConns(t *testing.T) {
	var (
		c     config
		value = 5
	)

	opt := WithMaxIdleConns(value)
	opt(&c)

	if c.MaxIdleConns != value {
		t.Errorf("MaxIdleConns don't match got ('%d'), want ('%d')", c.MaxIdleConns, maxIdleConns)
	}
}

func TestWithDuration(t *testing.T) {
	c := &config{}

	for _, test := range []struct {
		name   string
		d      time.Duration
		fn     func(time.Duration) Option
		result *time.Duration
	}{
		{"WithTimeout", 1, WithTimeout, &c.Timeout},
		{"WithKeepAlive", 2, WithKeepAlive, &c.KeepAlive},
		{"WithIdleConnTimeout", 3, WithIdleConnTimeout, &c.IdleConnTimeout},
		{"WithTLSHandshakeTimeout", 4, WithTLSHandshakeTimeout, &c.TLSHandshakeTimeout},
		{"WithExpectContinueTimeout", 5, WithExpectContinueTimeout, &c.ExpectContinueTimeout},
	} {
		t.Run(test.name, func(t *testing.T) {
			opt := test.fn(test.d)
			opt(c)

			if test.d != *test.result {
				t.Errorf("unexpect %s result got ('%v'), want ('%v')", test.name, *test.result, test.d)
			}
		})
	}
}

func TestWithTLSConfig(t *testing.T) {
	var (
		c     config
		value = tls.Config{
			ServerName: "testing-server",
		}
	)

	opt := WithTLSConfig(&value)
	opt(&c)

	if c.TLSConfig.ServerName != value.ServerName {
		t.Errorf("TLS ServerName don't match got ('%s'), want ('%s')", c.TLSConfig.ServerName, value.ServerName)
	}
}

func TestWithEnableHTTP2(t *testing.T) {
	var (
		c     config
		value = true
	)

	opt := WithEnableHTTP2(value)
	opt(&c)

	if c.EnableHTTP2 != value {
		t.Errorf("failed to enable HTTP 2 want ('%v'), got ('%v')", value, c.EnableHTTP2)
	}
}
