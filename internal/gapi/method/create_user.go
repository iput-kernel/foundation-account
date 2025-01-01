package method

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/iput-kernel/foundation-account/internal/application/auth"
	"github.com/iput-kernel/foundation-account/internal/domain"
	"github.com/iput-kernel/foundation-account/internal/infra/db/repository"
	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
	"github.com/iput-kernel/foundation-account/internal/infra/worker"
	accountv1 "github.com/iput-kernel/foundation-account/internal/pb/account/auth/v1"
	"github.com/iput-kernel/foundation-account/internal/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Method) CreateUser(ctx context.Context, req *accountv1.CreateUserRequest) (*accountv1.CreateUserResponse, error) {
	role := domain.DetectRole(req.GetEmail())
	if role == nil {
		return nil, status.Errorf(codes.InvalidArgument, "サービス対象外のメールアドレスです。")
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "不明なエラーが発生しました。")
	}
	hashedPassword, err := auth.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "不明なエラーが発生しました。")
	}
	arg := repository.TxCreateUserParam{
		CreateVerifyEmailParams: db.CreateVerifyEmailParams{
			ID:           id,
			Name:         req.GetUsername(),
			Email:        req.GetEmail(),
			PasswordHash: hashedPassword,
			SecretCode:   util.RandomString(32),
		},
		AfterCreate: func(user db.VerifyEmail) error {
			taskPayload := &worker.PayloadSendVerifyEmail{
				ID: id,
			}
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.QueueCritical),
			}
			return server.TaskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
		},
	}
	txResult, err := server.Store.TxCreateUser(ctx, arg)
	if err != nil {
		if repository.ErrorCode(err) == repository.UniqueViolation {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "ユーザーの作成に失敗: %s", err)
	}

	rsp := &accountv1.CreateUserResponse{
		User: convertUser(txResult.User),
	}

	return rsp, nil
}
