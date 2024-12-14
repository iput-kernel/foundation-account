package gapi

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) authFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, err
	}

	payload, err := server.tokenMaker.VerifyToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "無効なトークンです: %s", err)
	}

	// コンテキストにペイロードを追加して、後でアクセスできるようにする
	newCtx := context.WithValue(ctx, "payload", payload)
	return newCtx, nil
}
