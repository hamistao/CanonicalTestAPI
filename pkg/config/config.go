package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// represents the server configuration
type Config struct {
	ServerPort int    `yaml:"server_port"`
	DSN        string `yaml:"dsn"`
}

// loads configurations from .yml file
func Load(file string) (*Config, error) {

	var cfg Config
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(bytes, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
