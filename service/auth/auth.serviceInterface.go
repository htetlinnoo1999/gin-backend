package service

import (
	"xpm-auth/data/request"
)

type AuthService interface {
	Register(user request.RegisterRequest) (string, error)
	Login(credentials request.LoginRequest) (string, error)
}
