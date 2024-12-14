package gapi

import (
	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
	"github.com/iput-kernel/foundation-account/internal/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.VerifyEmail) *pb.User {
	return &pb.User{
		Username:  user.Name,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}
