package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ServerPort          int    `yaml:"server_port"`
	BooksFilePath       string `yaml:"books_file_path"`
	CollectionsFilePath string `yaml:"collections_file_path"`
}

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
