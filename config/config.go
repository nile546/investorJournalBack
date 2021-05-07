package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config ...
type Config struct {
	Production       bool
	Address          string
	Port             string
	ConnectionString string
	DatabaseHost     string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Production:       true,
		Address:          "localhost",
		Port:             "4000",
		ConnectionString: "111",
		DatabaseHost:     "localhost:5432",
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
