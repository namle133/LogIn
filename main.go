package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Product struct {
	Db *gorm.DB
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

var jwtKey = []byte("my-secrect-key")

type IProduct interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	LogOut(c *gin.Context)
}

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=Namle311 dbname=book port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Migrator().CurrentDatabase()
	return db
}

func (p *Product) SignIn(c *gin.Context) {
	var creds User
	err := c.BindJSON(&creds)
	if err != nil {
		c.String(http.StatusBadRequest, "%v", err)
		return
	}
	e := p.Db.First(&creds, "username=? AND password = ? AND email=?", creds.Username, creds.Password, creds.Email).Error
	if e != nil {
		c.String(http.StatusBadRequest, "%v", err)
		return
	}

	expirationTime := time.Now().Add(3 * time.Minute)
	claims := Claims{
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
		return
	}

	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	c.String(http.StatusOK, "Welcome to  %v", claims.Username)
}

func (p *Product) SignUp(c *gin.Context) {
	_, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.String(http.StatusUnauthorized, "%v", err)
			return
		}
		c.String(http.StatusBadRequest, "%v", err)
		return
	}
	var creds User
	er := c.BindJSON(&creds)
	creds1 := &User{Username: creds.Username, Password: creds.Password, Email: creds.Email}
	if er != nil {
		c.String(http.StatusBadRequest, "%v", er)
		return
	}
	p.Db.Create(&creds1)

	c.String(http.StatusOK, "%s", "SignUp Successfully!")
}

func (p *Product) LogOut(c *gin.Context) {
	ck, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.String(http.StatusUnauthorized, "%v", err)
			return
		}
		c.String(http.StatusBadRequest, "%v", err)
		return
	}
	c.SetCookie("token", ck, -1, "/", "localhost", false, true)
	c.String(http.StatusOK, "%s", "Old cookie deleted. Logged out!")
}

func main() {
	r := gin.Default()
	p := &Product{Db: ConnectDatabase()}
	var i IProduct = p
	r.POST("/signin", func(c *gin.Context) {
		i.SignIn(c)
	})
	r.POST("/signup", func(c *gin.Context) {
		i.SignUp(c)
	})
	r.POST("/logout", func(c *gin.Context) {
		i.LogOut(c)
	})
	r.Run(":8000")
}
