package instruments

import (
	"encoding/csv"
	"errors"
	"io"

	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/nile546/diplom/internal/models"
)

type Stockinstrument struct {
	mp map[string]models.StockInstrument
}

func (r *Stockinstrument) GrabAll(spburl string, mskurl string) (*[]models.StockInstrument, error) {
	r.mp = make(map[string]models.StockInstrument)
	err := r.spbgrab(spburl)
	if err != nil {
		return nil, err
	}

	err = r.mskgrab(mskurl)
	if err != nil {
		return nil, err
	}

	stocks := &[]models.StockInstrument{}

	for _, stockInstrument := range r.mp {
		*stocks = append(*stocks, stockInstrument)
	}

	return stocks, nil
}

func (r *Stockinstrument) spbgrab(u string) error {
	var resp *http.Response
	var req *http.Request
	var err error

	req, err = http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}

	type payload struct {
		headers  map[string]string
		formData map[string]string
	}

	client := &http.Client{}

	resp, err = client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("SPBExchange not respond")
	}

	pl := new(payload)
	pl.headers = make(map[string]string)
	pl.formData = make(map[string]string)

	pl.headers["Cookie"] = resp.Cookies()[0].Name + "=" + resp.Cookies()[0].Value

	// Parse and store
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
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
		return err //nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil
	}

	var ID int64 = 0

	cs := csv.NewReader(resp.Body)
	cs.FieldsPerRecord = -1
	cs.LazyQuotes = true
	cs.Comma = ';'
	for {
		record, err := cs.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf("Read error: Stock(SPB) with id = %d : %+v", ID, err)
			break
		}
		if ID == 0 {
			ID++
			continue
		}
		title, err := convertWin1251toUTF8(record[2])
		if err != nil {
			log.Errorf("Convert title error: Stock(SPB) with id = %d : %+v", ID, err)
			ID++
			continue
		}

		ticker, err := convertWin1251toUTF8(record[6])
		if err != nil {
			log.Errorf("Convert ticker error: Stock(SPB) with id = %d : %+v", ID, err)
			ID++
			continue
		}

		isin, err := convertWin1251toUTF8(record[7])
		if err != nil {
			log.Errorf("Convert ticker error: Stock(SPB) with id = %d : %+v", ID, err)
			ID++
			continue
		}

		if strings.Contains(ticker, "обыкновенн") {
			ticker, err = convertWin1251toUTF8(record[7])
			if err != nil {
				log.Errorf("Convert ticker error: Stock(SPB) with id = %d : %+v", ID, err)
				ID++
				continue
			}
			isin, err = convertWin1251toUTF8(record[8])
			if err != nil {
				log.Errorf("Convert ticker error: Stock(SPB) with id = %d : %+v", ID, err)
				ID++
				continue
			}
		}

		if strings.Contains(ticker, "расписки") {
			ticker, err = convertWin1251toUTF8(record[9])
			if err != nil {
				log.Errorf("Convert ticker error: Stock(SPB) with id = %d : %+v", ID, err)
				ID++
				continue
			}
			isin, err = convertWin1251toUTF8(record[10])
			if err != nil {
				log.Errorf("Convert ticker error: Stock(SPB) with id = %d : %+v", ID, err)
				ID++
				continue
			}
		}

		tp, err := convertWin1251toUTF8(record[4])
		if err != nil {
			log.Errorf("Convert type error: Stock(SPB) with id = %d : %+v", ID, err)
			ID++
			continue
		}

		if tp == "" {
			tp, err = convertWin1251toUTF8(record[5])
			if err != nil {
				log.Errorf("Convert type error: Stock(SPB) with id = %d : %+v", ID, err)
				ID++
				continue
			}
		}

		r.mp[isin] = models.StockInstrument{
			ID:     ID,
			Title:  title,
			Ticker: &ticker,
			Type:   &tp,
			Isin:   &isin,
		}

		ID++
	}

	return nil
}

func (r *Stockinstrument) mskgrab(u string) error {
	var resp *http.Response
	var req *http.Request
	var err error

	req, err = http.NewRequest(http.MethodGet, u, nil)
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

	var ID int64 = 0

	cs := csv.NewReader(resp.Body)
	cs.FieldsPerRecord = -1
	cs.LazyQuotes = true
	cs.Comma = ';'
	for {
		record, err := cs.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf("Read error: Stock(MSK) with id = %d : %+v", ID, err)
			break
		}
		if ID == 0 {
			ID++
			continue
		}

		ticker, err := convertWin1251toUTF8(record[7])
		if ticker == "" {
			continue
		}
		if err != nil {
			log.Errorf("Convert ticker error: Stock(MSK) with id = %d : %+v", ID, err)
			ID++
			continue
		}

		isin, err := convertWin1251toUTF8(record[8])
		if err != nil {
			log.Errorf("Convert ticker error: Stock(MSK) with id = %d : %+v", ID, err)
			ID++
			continue
		}

		title, err := convertWin1251toUTF8(record[11])
		if err != nil {
			log.Errorf("Convert title error: Stock(MSK) with id = %d : %+v", ID, err)
			ID++
			continue
		}

		tp, err := convertWin1251toUTF8(record[4])
		if err != nil {
			log.Errorf("Convert type error: Stock(MSK) with id = %d : %+v", ID, err)
			ID++
			continue
		}
		r.mp[isin] = models.StockInstrument{
			ID:     ID,
			Title:  title,
			Ticker: &ticker,
			Type:   &tp,
			Isin:   &isin,
		}

		ID++
	}

	return err
}
