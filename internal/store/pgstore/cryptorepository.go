package pgstore

import (
	"database/sql"
	"fmt"

	"github.com/nile546/diplom/internal/models"
)

type CryptoRepository struct {
	db *sql.DB
}

func (s *CryptoRepository) InsertCrypto(cryptos *[]models.Crypto) (err error) {

	q := `INSERT INTO crypto (title, ticker) VALUES ($1, $2)`

	var res sql.Result

	for _, crypto := range *cryptos {
		if res, err = s.db.Exec(q, crypto.Title, crypto.Ticker); err != nil {
			//TODO: ADD TO LOGER
			fmt.Println(err)
			continue
		}
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
