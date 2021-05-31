package instruments

import (
	"encoding/csv"
	"errors"
	"net/http"

	"github.com/nile546/diplom/internal/models"
)

type Bankinstruments struct {
}

func (d *Bankinstruments) GrabBanks(banksUrl string) (*[]models.Bank, error) {
	var resp *http.Response
	var req *http.Request
	var err error

	req, err = http.NewRequest(http.MethodGet, banksUrl, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Banki.ru not respond")
	}

	defer resp.Body.Close()

	banks := &[]models.Bank{}
	var ID int64 = 0

	cs := csv.NewReader(resp.Body)
	cs.FieldsPerRecord = -1
	cs.LazyQuotes = true
	cs.Comma = ';'
	for {
		record, err := cs.Read()
		if err != nil {
			//TODO: ADD TO LOGER
			break
		}
		if ID < 3 {
			ID++
			continue
		}
		title, err := convert(record[2])
		if err != nil {
			//TODO: ADD TO LOGER
			ID++
			continue
		}

		bank := models.Bank{
			ID:    ID,
			Title: title,
		}
		*banks = append(*banks, bank)
		ID++
	}

	return banks, nil
}
