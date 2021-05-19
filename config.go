package main

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Metroretro *MetroretroConfig `yaml:"metroretro"`
	Sections   []string          `yaml:"sections"`
	Github     *GithubConfig     `yaml:"github"`
}

type MetroretroConfig struct {
	Session    string `yaml:"session"`
	SessionSig string `yaml:"session_sig"`
}

type GithubConfig struct {
	Username    string `yaml:"username"`
	AccessToken string `yaml:"access_token"`
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

	if cfg.Github != nil && (cfg.Github.Username == "" || cfg.Github.AccessToken == "") {
		return nil, errors.New("invalid github section in configuration")
	}

	return cfg, nil
}
