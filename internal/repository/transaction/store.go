package repository

import (
	"context"

	db "github.com/iput-kernel/foundation-account/internal/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	db.Querier
	FollowTx(ctx context.Context, arg FollowTxParams) (FollowTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*db.Queries
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  db.New(connPool),
	}
}
