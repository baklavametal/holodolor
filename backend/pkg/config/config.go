package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config struct to hold application configuration
type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

// NewConfig creates a new Config struct by reading environment variables
func NewConfig() (*Config, error) {
	var conf Config
	err := envconfig.Process("HOLODOLOR", &conf)
	return &conf, err
}
