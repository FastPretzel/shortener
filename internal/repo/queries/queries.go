package queries

import "github.com/jackc/pgx/v5/pgxpool"

type Queries struct {
	pool *pgxpool.Pool
}

func New(p *pgxpool.Pool) *Queries {
	return &Queries{pool: p}
}
