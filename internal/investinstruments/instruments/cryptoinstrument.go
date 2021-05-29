package instruments

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Cryptoinstrument struct {
}

func (c *Cryptoinstrument) GrabCrypto() error {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		return err
	}

	//q := url.Values{}
	//	q.Add("start", "1")
	//	q.Add("limit", "5000")
	//	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "75416f37-656b-4dd1-8cf5-9e5a382d3e88")
	//	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))
	return nil

}
