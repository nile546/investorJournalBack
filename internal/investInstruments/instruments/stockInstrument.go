package instruments

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/nile546/diplom/internal/models"
)

type StockInstrument struct {
}

func (r *StockInstrument) GrabPage() (*[]models.Stock, error) {
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
	pl.formData = make(map[string]string)

	pl.headers["Cookie"] = resp.Cookies()[0].Name + "=" + resp.Cookies()[0].Value

	fmt.Println(resp.Body)

	// Parse and store
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("input").Each(func(index int, input *goquery.Selection) {
		switch name, _ := input.Attr("name"); name {

		case "__VIEWSTATE":
			pl.formData[name], _ = input.Attr("value")
		case "bxValidationToken":
			pl.formData[name], _ = input.Attr("value")
		case "__EVENTVALIDATION":
			pl.formData[name], _ = input.Attr("value")
		}
	})

	pl.headers["Host"] = "spbexchange.ru"
	pl.headers["Content-type"] = "application/x-www-form-urlencoded"
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

	pl.formData["bitrix_include_areas"] = "N"
	pl.formData["__EVENTTARGET"] = "ctl00$BXContent$list$LinkButton1"
	pl.formData["__EVENTARGUMENT"] = ""
	pl.formData["__LASTFOCUS"] = ""
	pl.formData["ctl00$searchform1$searchform1$searchform1$query"] = "Поиск..."
	pl.formData["ctl00$BXContent$list$tbSearch"] = ""
	pl.formData["ctl00$BXContent$list$ddlCBView"] = ""
	pl.formData["ctl00$BXContent$list$ddlCBCat"] = ""
	pl.formData["__VIEWSTATEGENERATOR"] = "8882E091"

	data := url.Values{}

	for k, v := range pl.formData {
		data.Set(k, v)
	}
	req2, err := http.NewRequest(http.MethodPost, u, strings.NewReader(data.Encode()))

	for k, v := range pl.headers {
		req2.Header.Add(k, v)
	}

	resp, err = client.Do(req2)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("SPBExchange not respond")
	}

	stocks := &[]models.Stock{}
	var ID int64 = 0

	cs := csv.NewReader(resp.Body)
	cs.FieldsPerRecord = -1
	cs.LazyQuotes = true
	cs.Comma = ';'
	for {
		record, e := cs.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		if ID != 0 {
			stock := models.Stock{
				ID:     ID,
				Title:  record[2],
				Ticket: record[6],
			}
			*stocks = append(*stocks, stock)
		}
		ID++
	}

	return stocks, nil
}
