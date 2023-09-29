package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	DefaultLibraryPath string      `yaml:"default-library-path"`
	MatchGroups        MatchGroups `yaml:"match"`
	source             string      `yaml:"-"`
}

func NewConfig(source string) *Config {
	return &Config{source: source}
}

func (c *Config) Source() string {
	return c.source
}

func (c *Config) readFile(path string) (bytesRead []byte, err error) {
	bytesRead, err = os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	return bytesRead, nil
}

func (c *Config) readYaml() error {
	yamlBytesRead, err := c.readFile(c.source)
	if err != nil {
		return fmt.Errorf("read yaml: %w", err)
	}

	err = yaml.Unmarshal(yamlBytesRead, c)
	if err != nil {
		return fmt.Errorf("ummarshal yaml: %w", err)
	}

	return nil
}

func (c *Config) Read() error {
	err := c.readYaml()
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	return nil
}

func (c *Config) writeFile(path string, data []byte) error {
	err := os.WriteFile(path, data, os.FileMode(0660))
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	return nil
}

// Create creates config file at path with current configuration
func (c *Config) Create() error {
	marshaledYaml, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("marshal yaml: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(c.source), os.FileMode(0774))
	if err != nil {
		return fmt.Errorf("make config dir: %w", err)
	}

	err = c.writeFile(c.source, marshaledYaml)
	if err != nil {
		return fmt.Errorf("write to file: %w", err)
	}

	return nil
}
