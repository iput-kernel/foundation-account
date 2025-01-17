package repository

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/iput-kernel/foundation-account/internal/config"
	"github.com/iput-kernel/foundation-account/internal/infra/db/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

// テストでのグローバル変数
var testDAO repository.DAO
var cfg config.Config

func TestMain(m *testing.M) {
	var err error
	cfg, err = config.LoadConfig("../../../../..")
	log.Printf("config:%+v ", cfg)
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), cfg.DSN())
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testDAO = repository.NewDAO(connPool)
	os.Exit(m.Run())
}
