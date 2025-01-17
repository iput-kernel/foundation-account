package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/iput-kernel/foundation-account/internal/config"
	"github.com/iput-kernel/foundation-account/internal/domain"
)

// トークンエラー　Invalidは無効な値
// Expiredは有効だったけど無効にされた
var (
	ErrInvalidToken = errors.New("トークンが無効")
	ErrExpiredToken = errors.New("トークンが期限切れ")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CredLevel int32     `json:"cred_level"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(config config.Config, email string, role string, credit int64, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	credLevel := domain.GetCredLevel(credit, config.Cred)

	payload := &Payload{
		ID:        tokenID,
		Username:  email,
		Role:      role,
		CredLevel: credLevel,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
