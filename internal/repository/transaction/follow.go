package repository

import (
	"context"
	"database/sql"
	"time"
)

// FollowTxParams contains the input parameters of the follow transaction
type FollowTxParams struct {
	FollowingUserID string
	FollowedUserID  string
}

// FollowTxResult is the result of the follow transaction
type FollowTxResult struct {
	Follow Follow
}

// Follow represents a follow relationship
type Follow struct {
	FollowingUserID string
	FollowedUserID  string
	CreatedAt       sql.NullTime
}

// FollowTx performs a follow transaction
func (store *SQLStore) FollowTx(ctx context.Context, arg FollowTxParams) (FollowTxResult, error) {
	var result FollowTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Insert a new follow record
		err = q.CreateFollow(ctx, arg.FollowingUserID, arg.FollowedUserID)
		if err != nil {
			return err
		}

		// Optionally, you can retrieve the follow record to return
		result.Follow = Follow{
			FollowingUserID: arg.FollowingUserID,
			FollowedUserID:  arg.FollowedUserID,
			CreatedAt:       sql.NullTime{Time: time.Now(), Valid: true},
		}

		return nil
	})

	return result, err
}
