package sys

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Connect string           `yaml:"connect"`
	Blocks  map[string]Block `yaml:"blocks"`
}

type Block struct {
	Name string `yaml:"name"`
}

func GetConfig(filename string) (*Config, error) {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading yaml file; %w", err)
	}
	var config = new(Config)
	if err := yaml.Unmarshal(yamlFile, config); err != nil {
		return nil, fmt.Errorf("error unmarhsalling yaml file; %w", err)
	}
	return config, nil
}
