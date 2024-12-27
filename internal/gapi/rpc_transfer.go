package gapi

import (
	"context"

	"github.com/iput-kernel/foundation-account/internal/infra/db/repository"
	accountv1 "github.com/iput-kernel/foundation-account/internal/pb/account/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) Transfer(ctx context.Context, req *accountv1.TransferRequest) (*accountv1.TransferResponse, error) {
	fromUser, err := server.store.GetUserByName(ctx, req.GetFromUserName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "内部エラーが発生しました")
	}

	toUser, err := server.store.GetUserByName(ctx, req.GetToUserName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "内部エラーが発生しました")
	}

	arg := repository.TxTransferParam{
		FromUser: fromUser,
		ToUser:   toUser,
		Amount:   req.GetAmount(),
	}

	result, err := server.store.TxTransfer(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "内部エラーが発生しました")
	}

	return &accountv1.TransferResponse{
		FromUserName: fromUser.Name,
		ToUserName:   toUser.Name,
		Amount:       result.Transfer.Amount,
	}, nil
}
