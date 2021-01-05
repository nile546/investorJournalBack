package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config ...
type Config struct {
	Address  string
	Port     string
	FilePath string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Address:  "localhost",
		Port:     "4000",
		FilePath: "./config/config.json",
	}
}

//LoadConfig ...
func (c *Config) LoadConfig() error {

	jsonFile, err := os.Open(c.FilePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	//var conf Config

	err = json.Unmarshal(byteValue, c)
	if err != nil {
		return err
	}

	return nil

}
