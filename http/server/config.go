package server

import "time"

type Config struct {
	Port              int
	Addr              string
	EnableTLS         bool
	TLS               *ConfigTLS
	ShutdownTimeout   time.Duration
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
}

type ConfigTLS struct {
	Port int
}

var defaultConfig = &Config{
	Port:              0,
	Addr:              "localhost",
	ShutdownTimeout:   5 * time.Second,
	ReadHeaderTimeout: 5 * time.Second,
	ReadTimeout:       15 * time.Second,
}

func setDefault(config *Config) {
	if config.Port == 0 {
		config.Port = defaultConfig.Port
	}

	if config.Addr == "" {
		config.Addr = defaultConfig.Addr
	}

	if config.ShutdownTimeout == 0 {
		config.ShutdownTimeout = defaultConfig.ShutdownTimeout
	}

	if config.ReadHeaderTimeout == 0 {
		config.ReadHeaderTimeout = defaultConfig.ReadTimeout
	}

	if config.ReadTimeout == 0 {
		config.ReadTimeout = defaultConfig.ReadTimeout
	}
}
