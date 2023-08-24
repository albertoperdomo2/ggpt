package ggpt

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

type Config struct {
	GptRole   string            `yaml:"gpt_role"`
	ApiURL    string            `yaml:"api_url"`
	ApiKey    string            `yaml:"api_key"`
	MaxTokens int               `yaml:"max_tokens"`
	Model     string            `yaml:"model"`
	Roles     map[string]string `yaml:"roles"`
}

func LoadConfig() (*Config, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(currentUser.HomeDir, ".ggpt_config.yaml")
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config Config
	yamlDecoder := yaml.NewDecoder(configFile)
	err = yamlDecoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
