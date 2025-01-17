package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
)

type TxTransferParam struct {
	FromUser db.User
	ToUser   db.User
	Amount   int64
}

type TxTransferResult struct {
	Transfer db.Transfer
}

func (store *SQLStore) TxTransfer(ctx context.Context, arg TxTransferParam) (TxTransferResult, error) {
	var result TxTransferResult

	err := store.execTx(ctx, func(q *db.Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, db.CreateTransferParams{
			FromUserID: arg.FromUser.ID,
			ToUserID:   arg.ToUser.ID,
			Amount:     arg.Amount,
		})
		if err != nil {
			return err
		}

		// デッドロックを防ぐために若い順に処理する
		if arg.FromUser.CreatedAt.Before(arg.ToUser.CreatedAt) {
			q.CreateStatement(ctx, db.CreateStatementParams{
				ID:     uuid.New(),
				UserID: arg.FromUser.ID,
				Amount: -arg.Amount,
				Reason: fmt.Sprintf("クレジット送信: %s", arg.ToUser.Name),
			})

			q.AddUserCredit(ctx, db.AddUserCreditParams{
				ID:     arg.FromUser.ID,
				Amount: arg.Amount,
			})

			q.CreateStatement(ctx, db.CreateStatementParams{
				ID:     uuid.New(),
				UserID: arg.ToUser.ID,
				Amount: arg.Amount,
				Reason: fmt.Sprintf("クレジット受信: %s", arg.FromUser.Name),
			})

			q.AddUserCredit(ctx, db.AddUserCreditParams{
				ID:     arg.ToUser.ID,
				Amount: arg.Amount,
			})
		} else {
			q.CreateStatement(ctx, db.CreateStatementParams{
				ID:     uuid.New(),
				UserID: arg.FromUser.ID,
				Amount: arg.Amount,
				Reason: fmt.Sprintf("クレジット受信: %s", arg.ToUser.Name),
			})

			q.AddUserCredit(ctx, db.AddUserCreditParams{
				ID:     arg.ToUser.ID,
				Amount: arg.Amount,
			})

			q.CreateStatement(ctx, db.CreateStatementParams{
				ID:     uuid.New(),
				UserID: arg.FromUser.ID,
				Amount: arg.Amount,
				Reason: fmt.Sprintf("クレジット送信: %s", arg.ToUser.Name),
			})

			q.AddUserCredit(ctx, db.AddUserCreditParams{
				ID:     arg.FromUser.ID,
				Amount: -arg.Amount,
			})
		}
		return nil
	})
	return result, err

}
