package service

import (
	"github.com/gin-gonic/gin"
	"github.com/namle133/LogIn.git/LogIn_Project/domain"
)

type IProduct interface {
	SignUp(c *gin.Context, creds *domain.User) error
	SignIn(c *gin.Context, creds *domain.User) (*domain.Claims, error)
	LogOut(c *gin.Context) error
}
