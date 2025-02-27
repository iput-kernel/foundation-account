package gapi

import (
	"context"

	"github.com/google/uuid"
	"github.com/iput-kernel/foundation-account/internal/domain"
	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
	accountv1 "github.com/iput-kernel/foundation-account/internal/pb/account/auth/v1"
	modelv1 "github.com/iput-kernel/foundation-account/internal/pb/account/model/v1"
	"github.com/iput-kernel/foundation-account/internal/validation"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) VerifyEmail(ctx context.Context, req *accountv1.VerifyEmailRequest) (*accountv1.VerifyEmailResponse, error) {
	violations := validateVerifyEmailRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	verifyEmailId, err := uuid.Parse(req.GetVerifyEmailId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "IDが無効な形式です")
	}
	verifyResult, err := server.store.GetVerifyEmail(ctx, verifyEmailId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "内部エラーが発生しました。")
	}
	if verifyResult.SecretCode != req.GetSecretCode() {
		return nil, status.Errorf(codes.InvalidArgument, "無効な認証コードです")
	}

	role := domain.DetectRole(verifyResult.Email)

	arg := db.CreateUserParams{
		ID:           verifyEmailId,
		Name:         verifyResult.Name,
		Email:        verifyResult.Email,
		PasswordHash: verifyResult.PasswordHash,
		Role:         *role,
		Credit:       int64(server.config.Cred.DefaultCredit),
		Level:        1,
	}

	createUserResult, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "本登録に失敗しました。")
	}

	rsp := &accountv1.VerifyEmailResponse{
		User: &modelv1.User{
			Email:     createUserResult.Email,
			Username:  createUserResult.Name,
			CreatedAt: timestamppb.New(createUserResult.CreatedAt),
		},
	}
	return rsp, nil
}

func validateVerifyEmailRequest(req *accountv1.VerifyEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validation.ValidateEmailId(req.GetVerifyEmailId()); err != nil {
		violations = append(violations, fieldViolation("email_id", err))
	}

	if err := validation.ValidateSecretCode(req.GetSecretCode()); err != nil {
		violations = append(violations, fieldViolation("secret_code", err))
	}

	return violations
}
