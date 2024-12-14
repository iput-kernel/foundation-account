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

var testDAO repository.DAO

func TestMain(m *testing.M) {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DSN())
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testDAO = repository.NewDAO(connPool)
	os.Exit(m.Run())
}
