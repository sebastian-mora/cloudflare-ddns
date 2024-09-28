package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type DNSRecord struct {
	Name    string `yaml:"name"`
	Content string `yaml:"content"`
	Proxied bool   `yaml:"proxied"`
	TTL     int    `yaml:"ttl"`
	Type    string `yaml:"type"`
}

type Zone struct {
	ZoneID  string               `yaml:"zoneId"`
	Name    string               `yaml:"name"`
	Records map[string]DNSRecord `yaml:"records"` // Use a map for records
}

type Config struct {
	Zones             []Zone `yaml:"zones"` // List of zones
	ExternalIPService string `yaml:"externalIpService"`
}

// LoadConfig loads the DNS records from a YAML file.
func LoadConfig(filename string) (*Config, error) {
	// Open the YAML file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Initialize the config structure
	var config Config

	// Unmarshal YAML into the config struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %v", err)
	}

	return &config, nil
}
