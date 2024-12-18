// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "users" (
  id,
  name,
  email,
  password_hash,
  role,
  credit,
  level
)
VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING id, name, email, password_hash, role, credit, level, updated_at, created_at
`

type CreateUserParams struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Role         Role      `json:"role"`
	Credit       int64     `json:"credit"`
	Level        int32     `json:"level"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.PasswordHash,
		arg.Role,
		arg.Credit,
		arg.Level,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.Credit,
		&i.Level,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, name, email, password_hash, role, credit, level, updated_at, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.Credit,
		&i.Level,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, password_hash, role, credit, level, updated_at, created_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.Credit,
		&i.Level,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT id, name, email, password_hash, role, credit, level, updated_at, created_at FROM users
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetUserByName(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByName, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.Credit,
		&i.Level,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
