package auth

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/iput-kernel/foundation-account/internal/config"
	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
	"github.com/o1egl/paseto"
)

// Pasetoトークン作成機
type PasetoMaker struct {
	paseto *paseto.V2
	config config.Config
}

// NewPasetoMaker creates a new PasetoMaker with asymmetric keys
func NewPasetoMaker(config config.Config) (Maker, error) {
	if len(config.SecretKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("無効なキーサイズ: %d 文字である必要があります", chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{
		paseto: paseto.NewV2(),
		config: config,
	}
	return maker, nil
}

func (maker *PasetoMaker) CreateToken(email string, role db.Role, credit int64, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(maker.config, email, string(db.RoleStudent), credit, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := maker.paseto.Encrypt([]byte(maker.config.SecretKey), payload, nil)
	return token, payload, err
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, []byte(maker.config.SecretKey), payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
