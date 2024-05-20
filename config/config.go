package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config holds the configuration for RabbitMQ
type Config struct {
	RabbitMQ RabbitMQConfig `yaml:"rabbitmq"`
}

// RabbitMQConfig holds the RabbitMQ specific configuration
type RabbitMQConfig struct {
	URL           string `yaml:"url"`
	ManagementURL string `yaml:"management_url"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
}

// LoadConfig loads the configuration from a YAML file
func LoadConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
