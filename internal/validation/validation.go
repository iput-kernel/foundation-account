package validation

import (
	"fmt"
	"net/mail"
	"regexp"

	"github.com/google/uuid"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("英数字とアンダースコアのみの３文字以上100文字以下を許容します")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("有効なEmailではありません")
	}
	return nil
}

func ValidateEmailId(value string) error {
	// Parse the string to check if it's a valid UUID
	if _, err := uuid.Parse(value); err != nil {
		return fmt.Errorf("UUIDの形式が間違っています")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}
