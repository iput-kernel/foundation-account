package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// bcrypt使ってハッシュ
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("パスワードのハッシュに失敗: %w", err)
	}
	return string(hashedPassword), nil
}

// bcryptでパスワードのチェック
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
