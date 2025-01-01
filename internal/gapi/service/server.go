package service

import (
	"crypto/ed25519"
	"fmt"

	"github.com/iput-kernel/foundation-account/internal/application/auth"
	"github.com/iput-kernel/foundation-account/internal/config"
	"github.com/iput-kernel/foundation-account/internal/infra/db/repository"
	"github.com/iput-kernel/foundation-account/internal/infra/worker"
	accountv1 "github.com/iput-kernel/foundation-account/internal/pb/account/service/v1"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	accountv1.UnimplementedAccountServiceServer
	PublicKey       ed25519.PublicKey
	Config          config.Config
	Store           repository.DAO
	TokenMaker      auth.Maker
	TaskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config config.Config, store repository.DAO, taskDistributor worker.TaskDistributor) (*Server, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, fmt.Errorf("公開鍵の生成に失敗: %w", err)
	}
	tokenMaker, err := auth.NewPasetoMaker(publicKey, privateKey)
	if err != nil {
		return nil, fmt.Errorf("トークン生成機の作成に失敗: %w", err)
	}

	server := &Server{
		PublicKey:       publicKey,
		Config:          config,
		Store:           store,
		TokenMaker:      tokenMaker,
		TaskDistributor: taskDistributor,
	}

	return server, nil
}
