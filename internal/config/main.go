package config

import (
	"example/go-echo-stuff/webserver/internal/utils"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
}

func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Host:      "0.0.0.0",
			Port:      80,
			Debug:     false,
			URL:       "http://0.0.0.0:80",
			ProxyMode: "direct",
			RootDir:   "/app/static",
		},
	}
}

func (c *Config) Dump() string {
	out, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("Writing yaml failed: %v", err)
	}

	return string(out)
}

func (c *Config) Load() {
	config_file := "config.yaml"

	// check if config.yaml exists
	if !utils.FileExists(config_file) {
		log.Printf("No config file found at %v, skipped", config_file)
		return
	}

	// read config.yaml into string
	yamlFile, err := os.ReadFile(config_file)
	if err != nil {
		log.Fatalf("Reading %v failed: %v", config_file, err)
	}

	// unmarshal yaml into configuration
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Parsing %v config.yaml failed: %v", config_file, err)
	}
}
