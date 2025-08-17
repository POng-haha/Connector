package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server       ServerConfig           `yaml:"server"`
	Logger       LoggerConfig           `yaml:"logger"`
	APIKeys      []APIKeyIncoming       `yaml:"apiKeys"`
	Destinations map[string]Destination `yaml:"destinations"`
	Routes       map[string]Route       `yaml:"routes"`
}
type ServerConfig struct {
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}
type LoggerConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}
type APIKeyIncoming struct {
	Key         string   `yaml:"key"`
	ClientName  string   `yaml:"clientName"`
	Status      string   `yaml:"status"`
	Permissions []string `yaml:"permissions"`
}
type Destination struct {
	Type   string `yaml:"type"`
	Host   string `yaml:"host"`
	APIKey string `yaml:"apiKey"`
}
type Route struct {
	System  string `yaml:"System"`
	Service string `yaml:"Service"`
	Format  string `yaml:"Format"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
