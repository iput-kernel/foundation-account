package repository

import (
	"context"

	"github.com/google/uuid"
	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
)

type TxFollowParams struct {
	FollowingUserID uuid.UUID `json:"following_user_id"`
	FollowedUserID  uuid.UUID `json:"followed_user_id"`
}

type TxFollowResult struct {
	Follow db.Follow `json:"follow"`
}

// FollowTx performs a follow transaction
func (store *SQLStore) TxFollow(ctx context.Context, arg TxFollowParams) (TxFollowResult, error) {
	var result TxFollowResult

	err := store.execTx(ctx, func(q *db.Queries) error {
		var err error

		result.Follow, err = q.CreateFollow(ctx, db.CreateFollowParams{
			FollowingUserID: arg.FollowingUserID,
			FollowedUserID:  arg.FollowedUserID,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
