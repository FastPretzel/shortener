package repo

import (
	"shortener/internal/repo/queries"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	repoPG struct {
		*queries.Queries
		pool *pgxpool.Pool
	}

	repoIM struct {
		mem map[string]string
		//origShortMap map[string]string
		mu *sync.Mutex
	}
)

func NewPG(p *pgxpool.Pool) Repository {
	return &repoPG{
		Queries: queries.New(p),
		pool:    p,
	}
}

func NewIM() Repository {
	return &repoIM{
		mem: make(map[string]string),
		mu:  &sync.Mutex{},
	}
}
