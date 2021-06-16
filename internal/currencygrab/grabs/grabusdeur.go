package grabscurrency

import (
	"net/http"
)

type GrabCBR struct {
}

func (g *GrabCBR) GrabUsdEur() error {

	var resp *http.Response
	var req *http.Request
	var err error

	req, err = http.NewRequest(http.MethodGet, "http://www.cbr.ru/scripts/XML_daily.asp", nil)
	if err != nil {
		return err
	}

	client := &http.Client{}

	resp, err = client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return err
	}

	defer resp.Body.Close()

	return nil

}
