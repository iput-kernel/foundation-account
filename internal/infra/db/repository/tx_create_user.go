package repository

import (
	"context"

	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
)

type TxCreateUserParam struct {
	db.CreateVerifyEmailParams
	AfterCreate func(user db.VerifyEmail) error
}

type TxCreateUserResult struct {
	User db.VerifyEmail
}

func (store *SQLStore) TxCreateUser(ctx context.Context, arg TxCreateUserParam) (TxCreateUserResult, error) {
	var result TxCreateUserResult

	err := store.execTx(ctx, func(q *db.Queries) error {
		var err error

		result.User, err = q.CreateVerifyEmail(ctx, arg.CreateVerifyEmailParams)
		if err != nil {
			return err
		}

		return arg.AfterCreate(result.User)
	})
	return result, err
}
