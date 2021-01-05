package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config ...
type Config struct {
	Address string
	Port    string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Address: "localhost",
		Port:    "4000",
	}
}

//LoadConfig ...
func (c *Config) LoadConfig(filePath string) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, c)
	if err != nil {
		return err
	}

	return nil

}
