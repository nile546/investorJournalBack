package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config ...
type Config struct {
	Production       bool
	Protocol         string
	Host             string
	Port             string
	ConnectionString string
	DatabaseHost     string
	TokenKey         string
	MailerLogin      string
	MailerPass       string
	MailerSender     string
	MailerHost       string
	MailerPort       string
	LandingAddress   string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Production:       true,
		Protocol:         "http://",
		Host:             "localhost",
		Port:             "4000",
		ConnectionString: "111",
		DatabaseHost:     "localhost:5432",
		TokenKey:         "tokenkey",
		LandingAddress:   "localhost:4200",
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
