package gapi

import (
	"crypto/ed25519"
	"fmt"

	"github.com/iput-kernel/foundation-account/internal/application/auth"
	"github.com/iput-kernel/foundation-account/internal/config"
	"github.com/iput-kernel/foundation-account/internal/infra/db/repository"
	"github.com/iput-kernel/foundation-account/internal/infra/worker"
	"github.com/iput-kernel/foundation-account/internal/pb"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedAccountServiceServer
	publicKey       ed25519.PublicKey
	config          config.Config
	store           repository.DAO
	tokenMaker      auth.Maker
	taskDistributor worker.TaskDistributor
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
		publicKey:       publicKey,
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
