package pgstore

import "database/sql"

type TinkoffTokenRepository struct {
	db *sql.DB
}

func (t *TinkoffTokenRepository) InsertTinkoffToken(token string, userID int64) error {

	q := "INSERT INTO tinkoff_tokens(token, user_id) VALUES ($1, $2)"

	res, err := t.db.Exec(q, token, userID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (t *TinkoffTokenRepository) GetTinkoffToken(userID int64) (string, error) {

	q := "SELECT token FROM tinkoff_tokens WHERE id=$1"

	var token string

	err := t.db.QueryRow(q, userID).Scan(&token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (t *TinkoffTokenRepository) UpdateTinkoffToken(token string, userID int64) error {

	q := "UPDATE tinkoff_tokens SET token = $1 WHERE user_id=$2"

	res, err := t.db.Exec(q, token, userID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
