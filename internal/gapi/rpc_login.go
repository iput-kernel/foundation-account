package gapi

import (
	"context"

	"github.com/google/uuid"
	"github.com/iput-kernel/foundation-account/internal/application/auth"
	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
	"github.com/iput-kernel/foundation-account/internal/pb"
	"github.com/iput-kernel/foundation-account/internal/validation"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	violations := validateLoginRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	authResult, err := server.store.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "内部エラーです")
	}

	err = auth.CheckPassword(authResult.PasswordHash, req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "IDまたはパスワードが正しくありません")
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		authResult.ID.String(),
		authResult.Role,
		server.config.Token.AccessDuration,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "トークンの作成に失敗しました")
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		authResult.ID.String(),
		authResult.Role,
		server.config.Token.RefreshDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "トークンの作成に失敗しました")
	}

	mtdt := server.extractMetadata(ctx)
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
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

	rsp := &pb.LoginResponse{
		User: &pb.User{
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

func validateLoginRequest(req *pb.LoginRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validation.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}
	if err := validation.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	return violations
}
