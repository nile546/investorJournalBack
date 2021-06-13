package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config ...
type Config struct {
	Production               bool
	Protocol                 string
	Host                     string
	Port                     string
	ConnectionString         string
	DatabaseHost             string
	TokenKey                 string
	MailerLogin              string
	MailerPass               string
	MailerSender             string
	MailerHost               string
	MailerPort               string
	LandingAddress           string
	SpbexchangeAddress       string
	MskexchangeAddress       string
	BankiUrl                 string
	CryptoUrl                string
	CryptoKey                string
	HoursUpdateInstruments   int
	MinutesUpdateInstruments int
	SecondsUpdateInstruments int
	LogLevel                 string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Production:               true,
		Protocol:                 "http://",
		Host:                     "localhost",
		Port:                     "4000",
		ConnectionString:         "111",
		DatabaseHost:             "localhost:5432",
		TokenKey:                 "tokenkey",
		LandingAddress:           "localhost:4200",
		SpbexchangeAddress:       "https://spbexchange.ru/ru/listing/securities/list/",
		MskexchangeAddress:       "https://www.moex.com/ru/listing/securities-list-csv.aspx?type=2",
		BankiUrl:                 "https://www.banki.ru/banks/ratings/export.php",
		CryptoUrl:                "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest",
		CryptoKey:                "75416f37-656b-4dd1-8cf5-9e5a382d3e88",
		HoursUpdateInstruments:   3,
		MinutesUpdateInstruments: 0,
		SecondsUpdateInstruments: 0,
		LogLevel:                 "debug",
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
