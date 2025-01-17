package repository

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/iput-kernel/foundation-account/internal/application/auth"
	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
	"github.com/iput-kernel/foundation-account/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) db.User {
	userID, err := uuid.NewUUID()
	require.NoError(t, err)
	hashedPassword, err := auth.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := db.CreateUserParams{
		ID:           userID,
		Name:         gofakeit.Name(),
		PasswordHash: hashedPassword,
		Email:        gofakeit.Email(),
		Role:         db.RoleStudent,
		Credit:       cfg.Cred.DefaultCredit,
	}

	user, err := testDAO.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.ID, user.ID)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Role, user.Role)
	require.NotZero(t, user.CreatedAt)

	return user
}

func deleteUser(t *testing.T, userID uuid.UUID) {
	err := testDAO.DeleteUser(context.Background(), userID)
	require.NoError(t, err)
}

func TestUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)

	deleteUser(t, user1.ID)
	deleteUser(t, user2.ID)
}
