package config

import (
	"io"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Listen string `yaml:"listen"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func ReadConfig(in io.Reader) (*Config, error) {
	var config *Config
	err := yaml.NewDecoder(in).Decode(&config)
	return config, err
}