package repository

import (
	"errors"
	model "xpm-auth/models"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	Db *gorm.DB
}

// FindByEmail implements UserRepository.
func (u *AuthRepositoryImpl) FindByEmail(email string) (users model.User, err error) {
	var user model.User
	result := u.Db.First(&user, "email = ?", email)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("user does not exist")
	}
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

// Insert implements UserRepository.
func (u *AuthRepositoryImpl) Insert(user model.User) (string, error) {
	err := u.Db.Create(&user).Error

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && (pgErr.Code == "23505") {
			return "", errors.New("user already existed")
		}
	}
	return "User created successfully.", nil
}

func NewAuthRepositoryImpl(Db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{Db: Db}
}
