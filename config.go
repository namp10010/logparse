// Package config provides a type that represent configuration
package main

import (
	"errors"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config defines configuration
type Config struct {
	Fields []string `yaml:"fields" mapstructure:"fields"`
	Skips  []string `yaml:"skips"  mapstructure:"skips"`
}

// LoadConfig loads configuration from file, env and flags and return compiled and validated config
func LoadConfig() (*Config, error) {
	v := viper.New()

	f := flag.CommandLine
	f.StringP("config", "c", "./config.yaml", "The configuration file to use to configure this application")
	flag.Parse()

	configFile, err := f.GetString("config")
	if err != nil {
		exit(err, 1)
	}
	if configFile == "" {
		return nil, errors.New("missing config")
	}

	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		exit(err, 1)
	}

	return &config, nil
}