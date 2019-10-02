package main

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

// Config ...
type Config struct {
	Apps []App `yaml:"apps"`
}

// App ...
type App struct {
	Name string `yaml:"name"`
	Up   map[string]interface{}
	Down map[string]interface{}
}

// ParseConfig ...
func ParseConfig(b []byte) (*Config, error) {
	config := &Config{}

	err := yaml.Unmarshal(b, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// AppNotFound ...
type AppNotFound struct {
	Name string
}

func (a AppNotFound) Error() string {
	return fmt.Sprintf("Application %s not found", a.Name)
}

// ConfigFor ...
func ConfigFor(config *Config, appName string, direction string) (map[string]interface{}, error) {
	for _, app := range config.Apps {
		if app.Name == appName {
			switch direction {
			case "UP":
				return app.Up, nil
			case "DOWN":
				return app.Down, nil
			}
		}
	}

	return nil, &AppNotFound{Name: appName}
}
