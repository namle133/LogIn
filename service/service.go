package service

import (
	"context"
	"github.com/namle133/LogIn.git/LogIn_Project/domain"
)

type IUser interface {
	CreateUser(c context.Context, u *domain.UserInit) error
	SignIn(c context.Context, u *domain.UserInit) (*domain.Claims, error)
	UserAdmin() error
	CheckRowToken(c context.Context) error
	LogOut(c context.Context) error
}
