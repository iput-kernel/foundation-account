package auth

import (
	"time"

	db "github.com/iput-kernel/foundation-account/internal/infra/db/sqlc"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/ed25519"
)

// Pasetoトークン作成機
type PasetoMaker struct {
	paseto     *paseto.V2
	publicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey
}

// NewPasetoMaker creates a new PasetoMaker with asymmetric keys
func NewPasetoMaker(publicKey ed25519.PublicKey, privateKey ed25519.PrivateKey) (Maker, error) {
	maker := &PasetoMaker{
		paseto:     paseto.NewV2(),
		publicKey:  publicKey,
		privateKey: privateKey,
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(email string, role db.Role, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(email, string(db.RoleStudent), duration)
	if err != nil {
		return "", payload, err
	}

	token, err := maker.paseto.Sign(maker.privateKey, payload, nil)
	return token, payload, err
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Verify(token, maker.publicKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
