package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Metroretro *MetroretroConfig `yaml:"metroretro"`
	Sections   []string          `yaml:"sections"`
}

type MetroretroConfig struct {
	Session    string `yaml:"session"`
	SessionSig string `yaml:"session_sig"`
}

func newConfig(filepath string) (*Config, error) {
	fp, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	if err = yaml.NewDecoder(fp).Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
