package usecase

import (
	"log"
	"os"
	"time"

	"github.com/iput-kernel/foundation-account/internal/model"
	"github.com/iput-kernel/foundation-account/internal/repository"
	"github.com/iput-kernel/foundation-account/internal/validator"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(req model.UserRequest) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(req model.UserRequest) (model.UserResponse, error) {
	log.Println("Request.password:", req.Password)
	if err := uu.uv.RegisterValidate(req); err != nil {
		log.Println("Request Validate Error", err)
		return model.UserResponse{}, err
	}

	userId := uuid.New()

	auth := &model.Auth{
		CredLevel:    2,
		PasswordHash: req.Password,
		UserID:       userId,
	}
	log.Println("Auth:", auth)

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}

	auth.PasswordHash = string(hash)
	newUser := model.User{
		ID:    userId,
		Email: req.Email,
		Name:  req.Name,
		Auth:  auth,
	}
	log.Println("New user:", newUser)
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Auth.PasswordHash), []byte(user.Auth.PasswordHash))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
