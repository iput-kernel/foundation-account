// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddUserCredit(ctx context.Context, arg AddUserCreditParams) (User, error)
	CreateFollow(ctx context.Context, arg CreateFollowParams) (Follow, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateStatement(ctx context.Context, arg CreateStatementParams) (Statement, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetStatements(ctx context.Context, id uuid.UUID) (Statement, error)
	GetTransfer(ctx context.Context, id uuid.UUID) (Transfer, error)
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByName(ctx context.Context, name string) (User, error)
	GetVerifyEmail(ctx context.Context, id uuid.UUID) (VerifyEmail, error)
	ListFollows(ctx context.Context, arg ListFollowsParams) ([]Follow, error)
	ListStatements(ctx context.Context, arg ListStatementsParams) ([]Statement, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	Verify(ctx context.Context, arg VerifyParams) (VerifyEmail, error)
}

var _ Querier = (*Queries)(nil)
