package validator

import (
	"github.com/iput-kernel/foundation-account/internal/domain"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IRegisterRequestValidator interface {
	RegisterValidate(user domain.RegisterRequest) error
	LoginValidate(user domain.LoginRequest) error
}

type IUserValidator interface {
	UserValidate(user domain.RegisterRequest) error
}

type userValidator struct{}
type RegisterValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func NewRegisterValidator() IRegisterValidator {
	return &RegisterValidator{}
}

func (uv *RegisterValidator) RegisterValidate(user domain.RegisterRequest) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 100).Error("limited max 100 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Username,
			validation.Required.Error("name is required"),
			validation.RuneLength(1, 20).Error("limited max 20 char"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("limited min 6 max 30 char"),
		),
	)
}

func (uv *userValidator) LoginValidate(user domain.LoginRequest) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 100).Error("limited max 100 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("limited min 6 max 30 char"),
		),
	)
}
