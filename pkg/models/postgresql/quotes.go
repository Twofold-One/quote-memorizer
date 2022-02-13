package postgresql

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Twofold-One/quote-memorizer/pkg/models"
)

type QuoteModel struct {
	DB *sql.DB
}

func (m *QuoteModel) Insert(author, quote string) (int, error) {

	stmt := `insert into quotes (author, quote, created)
	values ($1, $2, current_timestamp)`

	_, err := m.DB.Exec(stmt, author, quote)
	if err != nil {
		return 0, err
	}

	var id int
	row := m.DB.QueryRow(`select currval(
		pg_get_serial_sequence('quotes', 'id'));`)
		if err := row.Scan(&id); err != nil {
			if err == sql.ErrNoRows {
				return id, fmt.Errorf("id: %d: no such id", id)
			}
			return id, fmt.Errorf("id %d: %v", id, err)
		}
	return int(id), nil
}

func (m *QuoteModel) Get(id int) (*models.Quote, error) {

	stmt := `select id, author, quote, created from quotes
	where id = $1`

	row := m.DB.QueryRow(stmt, id)
	q := &models.Quote{}
	err := row.Scan(&q.ID, &q.Author, &q.Quote, &q.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return q, nil
}

func (m *QuoteModel) Latest() ([]*models.Quote, error) {
	stmt := `select id, author, quote, created from quotes
	order by created desc limit 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []*models.Quote

	for rows.Next() {
		q := &models.Quote{}
		err = rows.Scan(&q.ID, &q.Author, &q.Quote, &q.Created)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, q)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return quotes, nil
}