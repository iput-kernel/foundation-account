package auth

import (
	"crypto/ed25519"
	"testing"
	"time"

	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
	"github.com/iput-kernel/foundation-account/internal/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	require.NoError(t, err)
	maker, err := NewPasetoMaker(publicKey, privateKey)
	require.NoError(t, err)

	email := util.RandomEmail()
	role := db.RoleStudent
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(email, role, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, email, payload.Username)
	require.Equal(t, string(role), payload.Role)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	require.NoError(t, err)
	maker, err := NewPasetoMaker(publicKey, privateKey)
	require.NoError(t, err)

	token, payload, err := maker.CreateToken(util.RandomOwner(), db.RoleStudent, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
