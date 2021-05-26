package instruments

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/nile546/diplom/internal/models"
)

type StockInstrument struct {
}

func (r *StockInstrument) GrabPage() ([]*models.Stock, error) {
	u := "https://spbexchange.ru/ru/listing/securities/list/"
	var resp *http.Response
	var req *http.Request
	var err error

	req, err = http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	type payload struct {
		headers  map[string]string
		formData map[string]string
	}

	client := &http.Client{}

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("SPBExchange not respond")
	}

	pl := new(payload)
	pl.headers = make(map[string]string)

	pl.headers["Cookie"] = resp.Cookies()[0].Name + "=" + resp.Cookies()[0].Value

	fmt.Println(resp.Body)

	// Parse and store
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	t, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(t)

	doc.Find("input").Each(func(index int, input *goquery.Selection) {
		switch name, _ := input.Attr("name"); name {

		case "__VIEWSTATE":
			pl.formData[name], _ = input.Attr("value")
		case "bxValidationToken":
			pl.formData[name], _ = input.Attr("value")
		}
	})

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
	pl.formData["__EVENTTARGET"] = "ctl00$BXContent$list$LinkButton1"
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

	for k, v := range pl.formData {
		data.Set(k, v)
	}

	req, err = http.NewRequest(http.MethodPost, u, strings.NewReader(data.Encode()))

	for k, v := range pl.headers {
		req.Header.Add(k, v)
	}

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("SPBExchange not respond")
	}

	fmt.Println("zzzzzzzzzzzzzzzz", resp.Body)

	return nil, nil
}
