package main

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	HealthCheckInterval string `json:"healthCheckInterval"`
	RootDir             string `json:"-"`
	Port                string `json:"-"`
}

func (c *Config) loadConfig(file string) error {
	path := c.RootDir + file
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) loadEnvs() error {
	root, ok := os.LookupEnv("ROOT_DIR")
	if !ok {
		return errors.New("Failed to find 'ROOT_DIR' env!")
	}
	c.RootDir = root
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return errors.New("Failed to find 'PORT' env!")
	}
	c.Port = port
	return nil
}
