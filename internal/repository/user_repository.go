package repository

import (
	db "github.com/iput-kernel/foundation-account/internal/db/sqlc"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *db.CreateUserParams) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *db.CreateUserParams) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
