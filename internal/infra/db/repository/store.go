package repository

import (
	"context"

	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type DAO interface {
	db.Querier
	TxFollow(ctx context.Context, arg TxFollowParams) (TxFollowResult, error)
	TxCreateUser(ctx context.Context, arg TxCreateUserParam) (TxCreateUserResult, error)
}

// トランザクションに利用するツールとクエリを一通り格納
type SQLStore struct {
	connPool *pgxpool.Pool
	*db.Queries
}

// NewStore creates a new store
func NewDAO(connPool *pgxpool.Pool) DAO {
	return &SQLStore{
		connPool: connPool,
		Queries:  db.New(connPool),
	}
}
