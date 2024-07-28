package service

import (
	"errors"
	"time"
	"xpm-auth/data/request"
	"xpm-auth/helper"
	model "xpm-auth/models"
	repository "xpm-auth/repositories/auth"

	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	Validate       *validator.Validate
}

// Login implements AuthService.
func (u *AuthServiceImpl) Login(credentials request.LoginRequest) (string, error) {
	validateErr := u.Validate.Struct(credentials)
	helper.Error(validateErr)
	user, userErr := u.AuthRepository.FindByEmail(credentials.Email)
	if userErr != nil {
		return "", userErr
	}
	isCorrectPassword := helper.CheckPasswordHash(credentials.Password, user.Password)
	if !isCorrectPassword {
		return "", errors.New("invalid credentials")
	}
	token, jwtErr := helper.GenerateJWT(user)
	return token, jwtErr
}

// Register implements AuthService.
func (u *AuthServiceImpl) Register(user request.RegisterRequest) (string, error) {
	err := u.Validate.Struct(user)
	helper.Error(err)

	hashedPassword, hashError := helper.HashPassword(user.Password)
	if hashError != nil {
		helper.Error(hashError)
	}
	userModel := model.User{
		Name:      &user.Name,
		Email:     user.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, customError := u.AuthRepository.Insert(userModel)
	return result, customError
}

func NewAuthServiceImpl(authRepository repository.AuthRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		Validate:       validate,
	}
}
