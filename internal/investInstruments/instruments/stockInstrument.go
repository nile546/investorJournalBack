package instruments

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type StockInstrument struct {
}

type payload struct {
	headers  map[string]string
	formData map[string]string
	//instruments []*model.Instrument
	client *http.Client
}

func (r *StockInstrument) GrabPage() error {
	u := "https://spbexchange.ru/ru/listing/securities/list/"
	var resp *http.Response
	var req *http.Request
	var err error

	req, err = http.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return err
	}

	pl := &payload{}
	pl = new(payload)

	pl.client = &http.Client{}

	pl.headers = make(map[string]string)
	pl.headers["Host"] = "spbexchange.ru"
	pl.headers["Content-type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	pl.headers["Origin"] = "https://spbexchange.ru"
	pl.headers["Connection"] = "keep-alive"
	pl.headers["Pragma"] = "no-cache"
	pl.headers["Cache-Control"] = "no-cache"
	pl.headers["X-Requested-With"] = "XMLHttpRequest"
	pl.headers["X-MicrosoftAjax"] = "Delta=true"
	pl.headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36"
	pl.headers["Accept"] = "*/*"
	pl.headers["Sec-Fetch-Site"] = "same-origin"
	pl.headers["Sec-Fetch-Mode"] = "cors"
	pl.headers["Sec-Fetch-Dest"] = "empty"
	pl.headers["Referer"] = "https://spbexchange.ru/ru/listing/securities/list/"
	pl.headers["Accept-Language"] = "ru,ru-RU;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6"

	pl.formData = make(map[string]string)
	pl.formData["__EVENTARGUMENT"] = ""
	pl.formData["ctl00$ctl00$ctl02"] = "ctl00$ctl00$BXContentOuter$BXContent$up|ctl00$ctl00$BXContentOuter$BXContent$pager$ctl02$ctl00"
	pl.formData["__EVENTTARGET"] = "ctl00$ctl00$BXContentOuter$BXContent$pager$ctl02$ctl00"
	pl.formData["bitrix_include_areas"] = "N"
	pl.formData["__LASTFOCUS"] = ""
	pl.formData["ctl00$ctl00$BXContentOuter$BXContent$rbl"] = "1"
	pl.formData["ctl00$ctl00$BXContentOuter$BXContent$tbInstrument"] = ""
	pl.formData["ctl00$ctl00$BXContentOuter$BXContent$ddlRecom"] = ""
	pl.formData["ctl00$ctl00$BXContentOuter$BXContent$ddlSector"] = ""
	pl.formData["ctl00$ctl00$BXContentOuter$BXContent$ddlSubSektor"] = ""
	pl.formData["ctl00$ctl00$tbName"] = ""
	pl.formData["ctl00$ctl00$tbEmail"] = ""
	pl.formData["ctl00$ctl00$tbPhone"] = ""
	pl.formData["ctl00$ctl00$TextBox1"] = ""
	pl.formData["ctl00$ctl00$TextBox2"] = ""
	pl.formData["__VIEWSTATEGENERATOR"] = "65598315"
	pl.formData["__ASYNCPOST"] = "true"

	data := url.Values{}
	if pl.formData != nil {
		for k, v := range pl.formData {
			data.Set(k, v)
		}
	}

	req, err = http.NewRequest(http.MethodPost, u, strings.NewReader(data.Encode()))

	if pl.headers != nil {
		for k, v := range pl.headers {
			req.Header.Add(k, v)
		}
	}

	resp, err = pl.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("NOOK")
	}

	return nil
}
