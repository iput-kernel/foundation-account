package method

import (
	"context"

	"github.com/google/uuid"
	"github.com/iput-kernel/foundation-account/internal/application/auth"
	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
	accountv1 "github.com/iput-kernel/foundation-account/internal/pb/account/auth/v1"
	modelv1 "github.com/iput-kernel/foundation-account/internal/pb/account/model/v1"
	"github.com/iput-kernel/foundation-account/internal/validation"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Method) Login(ctx context.Context, req *accountv1.LoginRequest) (*accountv1.LoginResponse, error) {
	violations := validateLoginRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	authResult, err := server.Store.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "内部エラーです")
	}

	err = auth.CheckPassword(authResult.PasswordHash, req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "IDまたはパスワードが正しくありません")
	}

	accessToken, accessPayload, err := server.TokenMaker.CreateToken(
		authResult.ID.String(),
		authResult.Role,
		server.Config.Token.AccessDuration,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "トークンの作成に失敗しました")
	}

	refreshToken, refreshPayload, err := server.TokenMaker.CreateToken(
		authResult.ID.String(),
		authResult.Role,
		server.Config.Token.RefreshDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "トークンの作成に失敗しました")
	}

	mtdt := server.ExtractMetadata(ctx)
	session, err := server.Store.CreateSession(ctx, db.CreateSessionParams{
		ID:           uuid.New(),
		UserID:       authResult.ID,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "セッションの作成に失敗しました")
	}

	rsp := &accountv1.LoginResponse{
		User: &modelv1.User{
			Username:  authResult.Name,
			Email:     authResult.Email,
			CreatedAt: timestamppb.New(authResult.CreatedAt),
		},
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}

	return rsp, nil
}

func validateLoginRequest(req *accountv1.LoginRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validation.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}
	if err := validation.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	return violations
}
