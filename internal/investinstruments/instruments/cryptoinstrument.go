package instruments

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nile546/diplom/internal/models"
)

type Cryptoinstrument struct {
}

type extoCrypto struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func (c *Cryptoinstrument) GrabAll(cryptoUrl string, cryptoKey string) (*[]models.CryptoInstrument, error) {
	cryptos, err := grabCrypto(cryptoUrl, cryptoKey)
	if err != nil {
		return nil, err
	}
	return cryptos, nil
}

func grabCrypto(cryptoUrl string, cryptoKey string) (*[]models.CryptoInstrument, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", cryptoUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", cryptoKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	type data struct {
		Data []extoCrypto `json:"data"`
	}

	extoCryptoList := &data{}

	respBody, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(respBody, extoCryptoList)

	resultCrypto := &[]models.CryptoInstrument{}

	for _, extoCrypto := range extoCryptoList.Data {
		crypto, err := convertExtoCryptoToCrypto(&extoCrypto)
		if err != nil {
			//TODO: ADD TO LOGER
			continue
		}

		*resultCrypto = append(*resultCrypto, *crypto)
	}

	return resultCrypto, nil
}

func convertExtoCryptoToCrypto(c *extoCrypto) (*models.CryptoInstrument, error) {
	if c == nil {
		return nil, nil
	}

	return &models.CryptoInstrument{
		Title:  c.Name,
		Ticker: c.Symbol,
	}, nil

}
