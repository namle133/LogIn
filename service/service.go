package service

import (
	"context"
	"github.com/namle133/LogIn.git/LogIn_Project/domain"
)

type IUser interface {
	CreateUser(c context.Context, u *domain.User) error
	SignIn(c context.Context, creds *domain.User) (*domain.Claims, string, error)
	//LogOut(c context.Context) interface{}
}
