// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: follow.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createFollow = `-- name: CreateFollow :one
INSERT INTO follows (
  following_user_id,
  followed_user_id
) VALUES (
  $1, $2
) RETURNING following_user_id, followed_user_id, created_at
`

type CreateFollowParams struct {
	FollowingUserID uuid.UUID `json:"following_user_id"`
	FollowedUserID  uuid.UUID `json:"followed_user_id"`
}

func (q *Queries) CreateFollow(ctx context.Context, arg CreateFollowParams) (Follow, error) {
	row := q.db.QueryRow(ctx, createFollow, arg.FollowingUserID, arg.FollowedUserID)
	var i Follow
	err := row.Scan(&i.FollowingUserID, &i.FollowedUserID, &i.CreatedAt)
	return i, err
}

const listFollows = `-- name: ListFollows :many
SELECT following_user_id, followed_user_id, created_at FROM follows
WHERE 
    following_user_id = $1 OR
    followed_user_id = $2
ORDER BY created_at
LIMIT $3
OFFSET $4
`

type ListFollowsParams struct {
	FollowingUserID uuid.UUID `json:"following_user_id"`
	FollowedUserID  uuid.UUID `json:"followed_user_id"`
	Limit           int32     `json:"limit"`
	Offset          int32     `json:"offset"`
}

func (q *Queries) ListFollows(ctx context.Context, arg ListFollowsParams) ([]Follow, error) {
	rows, err := q.db.Query(ctx, listFollows,
		arg.FollowingUserID,
		arg.FollowedUserID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Follow{}
	for rows.Next() {
		var i Follow
		if err := rows.Scan(&i.FollowingUserID, &i.FollowedUserID, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
