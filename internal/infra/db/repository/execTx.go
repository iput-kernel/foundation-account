package repository

import (
	"context"
	"fmt"

	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
)

// トランザクションの実行
func (store *SQLStore) execTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("トランザクションエラー: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
