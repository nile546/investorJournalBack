package instruments

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nile546/diplom/internal/models"
)

type Cryptoinstrument struct {
}

func (c *Cryptoinstrument) GrabCrypto(cryptoUrl string) (*[]models.Crypto, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", cryptoUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "75416f37-656b-4dd1-8cf5-9e5a382d3e88")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	type data struct {
		Data []models.Crypto `json:"data"`
	}

	crypt := &data{}

	respBody, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(respBody, crypt)

	fmt.Println(crypt)
	return &crypt.Data, nil

}
