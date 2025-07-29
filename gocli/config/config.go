package config

import (
	"encoding/json"
	"gocli/types"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	DefaultConfigDir  = ".shitposteragent"
	DefaultConfigFile = "shitposteragent.json"
)

// GetConfigPath returns the config path in the user's home directory
func GetConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, DefaultConfigDir, DefaultConfigFile)
}

// LoadConfig loads config from the given path, or default if empty
func LoadConfig(path string) (*types.Config, error) {
	if path == "" {
		path = GetConfigPath()
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg types.Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// SaveConfig saves config to the default path
func SaveConfig(cfg *types.Config) error {
	path := GetConfigPath()
	os.MkdirAll(filepath.Dir(path), 0700)
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0600)
}
