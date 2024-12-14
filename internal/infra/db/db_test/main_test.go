package db_test

import (
	"context"
	"log"
	"testing"

	"os"

	"github.com/iput-kernel/foundation-account/internal/config"
	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../../../..")
	if err != nil {
		log.Fatal("環境変数の読み込みに失敗")
	}
	connPool, err := pgxpool.New(context.Background(), config.DSN())
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = db.New(connPool)
	os.Exit(m.Run())
}
