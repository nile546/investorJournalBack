package grabscurrency

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/nile546/diplom/internal/models"
	"golang.org/x/net/html/charset"
)

type GrabCBR struct {
}

type Valutes struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	Currency string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func (g *GrabCBR) GrabUsdEur() (*[]models.CurrencyRatio, error) {

	var resp *http.Response
	var req *http.Request
	var err error

	req, err = http.NewRequest(http.MethodGet, "http://www.cbr.ru/scripts/XML_daily.asp", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, err
	}

	defer resp.Body.Close()

	valutes := &Valutes{}

	byteValue, _ := ioutil.ReadAll(resp.Body)

	reader := bytes.NewReader(byteValue)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&valutes)
	if err != nil {
		return nil, err
	}

	currencyRatios := &[]models.CurrencyRatio{}

	for _, valute := range valutes.Valutes {

		if valute.Currency == "USD" || valute.Currency == "EUR" {
			currencyRatio := models.CurrencyRatio{
				FirstCurrency:  currencyConvert(valute.Currency),
				SecondCurrency: currencyConvert("RUB"),
				Ratio:          valute.Value,
			}
			*currencyRatios = append(*currencyRatios, currencyRatio)
		}

	}

	return currencyRatios, nil

}

func currencyConvert(currency string) models.Currencies {

	if currency == "USD" {
		return 1
	}

	if currency == "EUR" {
		return 2
	}

	if currency == "RUB" {
		return 3
	}

	return 0
}
