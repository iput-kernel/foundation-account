package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/iput-kernel/foundation-account/internal/application/auth"
	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
	"github.com/iput-kernel/foundation-account/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) db.User {
	hashedPassword, err := auth.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	userID, err := uuid.NewUUID()
	require.NoError(t, err)
	arg := db.CreateUserParams{
		ID:           userID,
		Name:         util.RandomOwner(),
		PasswordHash: hashedPassword,
		Email:        util.RandomEmail(),
		Role:         db.RoleStudent,
		Credit:       100,
		Level:        1,
	}

	user, err := testDAO.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
