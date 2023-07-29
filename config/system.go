package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DefaultProject string `yaml:"default_project"`
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	// Check mandatory configs
	if config.Port == "" {
		return nil, fmt.Errorf("host is mandatory")
	}

	return &config, nil
}
