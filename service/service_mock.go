package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/namle133/LogIn.git/LogIn_Project/domain"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Product struct {
	Db *gorm.DB
}

var jwtKey = []byte("my-secrect-key")

func (p *Product) SignIn(c *gin.Context, creds *domain.User) (*domain.Claims, error) {

	e := p.Db.First(&creds, "username=? AND password = ? AND email=?", creds.Username, creds.Password, creds.Email).Error
	if e != nil {
		c.String(http.StatusBadRequest, "%v", e)
		return nil, e
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
		c.String(http.StatusInternalServerError, "%v", err)
		return nil, e
	}

	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	return claims, nil
}

func (p *Product) SignUp(c *gin.Context, creds *domain.User) error {
	_, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.String(http.StatusUnauthorized, "%v", err)
			return err
		}
		c.String(http.StatusBadRequest, "%v", err)
		return err
	}
	creds1 := &domain.User{Username: creds.Username, Password: creds.Password, Email: creds.Email}
	p.Db.Create(&creds1)
	return nil

}

func (p *Product) LogOut(c *gin.Context) error {
	ck, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.String(http.StatusUnauthorized, "%v", err)
			return err
		}
		c.String(http.StatusBadRequest, "%v", err)
		return err
	}
	c.SetCookie("token", ck, -1, "/", "localhost", false, true)
	return nil

}
