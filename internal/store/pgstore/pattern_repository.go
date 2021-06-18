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

func (p *PatternRepository) GetAllPattern(tp *models.TableParams, id int64) error {

	q := `SELECT id, name, description, icon, instrument_type, user_id, created_at
	FROM patterns
	WHERE user_id IS NULL OR user_id=$1
	LIMIT $2 
	OFFSET $3;
	`

	rows, err := p.db.Query(
		q,
		id,
		tp.Pagination.ItemsPerPage,
		tp.Pagination.PageNumber*tp.Pagination.ItemsPerPage,
	)
	if err != nil {
		return err
	}

	source := []models.Pattern{}

	count := 0

	for rows.Next() {
		count++
		var p models.Pattern
		err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.Icon, &p.InstrumentType, &p.UserID, &p.CreatedAt)
		if err != nil {
			return err
		}

		source = append(source, p)
	}

	if count == 0 {
		return nil
	}

	tp.Source = source

	return nil
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
