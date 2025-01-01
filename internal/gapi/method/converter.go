package method

import (
	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
	accountv1 "github.com/iput-kernel/foundation-account/internal/pb/account/model/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.VerifyEmail) *accountv1.User {
	return &accountv1.User{
		Username:  user.Name,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}
