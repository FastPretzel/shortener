package queries

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

func (q *Queries) Get(ctx context.Context, shortLink string) (string, error) {
	selectLinkQuery := `SELECT orig_link FROM links WHERE short_link = $1`
	var origLink string
	err := q.pool.QueryRow(ctx, selectLinkQuery, shortLink).Scan(&origLink)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("rows expected but not returned: %w", err)
		}
		return "", fmt.Errorf("failed to select row from table: %w", err)
	}
	return origLink, nil
}

func (q *Queries) Create(ctx context.Context, origLink, shortLink string) error {
	insertLinkQuery := `INSERT INTO links (short_link, orig_link) VALUES ($1, $2)`
	_, err := q.pool.Exec(ctx, insertLinkQuery, shortLink, origLink)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return fmt.Errorf("row already exists")
		}
		return fmt.Errorf("failed to insert row into table: %w", err)
	}
	return nil
}
