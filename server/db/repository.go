package db

import (
	"database/sql"
	"fmt"

	"github.com/santosjordi/posgoexp/challenges/ctx-client-server/quote"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) SaveQuote(quote *quote.DolarQuote) error {
	query := `
		INSERT INTO quotes (code, codein, name, high, low, var_bid, pct_change, bid, ask, timestamp, create_date)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		quote.Code,
		quote.Codein,
		quote.Name,
		quote.High,
		quote.Low,
		quote.VarBid,
		quote.PctChange,
		quote.Bid,
		quote.Ask,
		quote.Timestamp,
		quote.CreateDate,
	)
	if err != nil {
		return fmt.Errorf("failed to execute prepared statement: %w", err)
	}

	return nil
}
