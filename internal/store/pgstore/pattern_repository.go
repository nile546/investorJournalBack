package pgstore

import (
	"database/sql"

	"github.com/nile546/diplom/internal/models"
)

type PatternRepository struct {
	db *sql.DB
}

func (p *PatternRepository) CreatePattern(pattern *models.Pattern) error {

	q := "INSERT INTO patterns (name, description, user_id, instrument_type, icon) VALUES ($1, $2, $3, $4, $5)"

	res, err := p.db.Exec(q, pattern.Name, pattern.Description, pattern.UserID, pattern.InstrumentType, pattern.Icon)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (p *PatternRepository) UpdatePattern(pattern *models.Pattern) error {

	q := "UPDATE patterns SET (name, description, icon) = ($1, $2, $3) WHERE id = $4"

	res, err := p.db.Exec(q, pattern.Name, pattern.Description, pattern.Icon, pattern.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (p *PatternRepository) GetAllPattern(userId int64) (*[]models.Pattern, error) {

	q := `SELECT * FROM patterns where user_id IS NULL OR user_id=$1`

	res, err := p.db.Query(q, userId)
	if err != nil {
		return nil, err
	}

	ptrns := &[]models.Pattern{}

	for res.Next() {
		ptrn := models.Pattern{}
		err = res.Scan(&ptrn.ID, &ptrn.Name, &ptrn.Description, &ptrn.Icon, &ptrn.UserID, &ptrn.CreatedAt)
		if err != nil {
			return nil, err
		}
		*ptrns = append(*ptrns, ptrn)
	}

	return ptrns, nil
}

func (p *PatternRepository) DeletePattern(id int64) error {

	q := "DELETE FROM patterns WHERE id=$1"

	res, err := p.db.Exec(q, id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil

}
