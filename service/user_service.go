package service

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/namle133/LogIn.git/LogIn_Project/domain"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	Db *gorm.DB
}

var jwtKey = []byte("my-secrect-key")

func (us *UserService) SignIn(c context.Context, creds *domain.User) (*domain.Claims, string, error) {
	e := us.Db.First(&creds, "username=? AND password = ? AND email=?", creds.Username, creds.Password, creds.Email).Error
	if e != nil {
		return nil, "", e
	}

	expirationTime := time.Now().Add(3 * time.Minute)
	claims := &domain.Claims{
		Username: creds.Username,
		Password: creds.Password,
		Email:    creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, "", e
	}
	return claims, tokenString, nil
}

func (us *UserService) CreateUser(c context.Context, u *domain.User) error {
	if u == nil {
		return fmt.Errorf("user does not empty")
	}
	if u.Username == "" {
		return fmt.Errorf("Error with username")
	}
	if u.Email == "" {
		return fmt.Errorf("Error with email")
	}
	if u.Password == "" {
		return fmt.Errorf("Error with password")
	}
	err := us.Db.Create(u).Error
	if err != nil {
		return err
	}
	return nil

}

//func (p *UserService) LogOut(c context.Context) string {
//}
