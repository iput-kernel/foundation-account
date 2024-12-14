// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: statement.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO statements (
  id,
  user_id,
  amount
) VALUES (
  $1, $2, $3
) RETURNING id, user_id, amount, reason, created_at
`

type CreateEntryParams struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Amount int64     `json:"amount"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Statement, error) {
	row := q.db.QueryRow(ctx, createEntry, arg.ID, arg.UserID, arg.Amount)
	var i Statement
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.Reason,
		&i.CreatedAt,
	)
	return i, err
}

const getStatements = `-- name: GetStatements :one
SELECT id, user_id, amount, reason, created_at FROM statements
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetStatements(ctx context.Context, id uuid.UUID) (Statement, error) {
	row := q.db.QueryRow(ctx, getStatements, id)
	var i Statement
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.Reason,
		&i.CreatedAt,
	)
	return i, err
}

const listStatements = `-- name: ListStatements :many
SELECT id, user_id, amount, reason, created_at FROM statements
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListStatementsParams struct {
	UserID uuid.UUID `json:"user_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) ListStatements(ctx context.Context, arg ListStatementsParams) ([]Statement, error) {
	rows, err := q.db.Query(ctx, listStatements, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Statement{}
	for rows.Next() {
		var i Statement
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			&i.Reason,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
