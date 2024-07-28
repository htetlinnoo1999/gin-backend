package repository

import model "xpm-auth/models"

type AuthRepository interface {
	Insert(user model.User) (string, error)
	FindByEmail(email string) (user model.User, err error)
}
