package main

import (
	"github.com/gin-gonic/gin"
	"github.com/namle133/LogIn.git/LogIn_Project/database"
	"github.com/namle133/LogIn.git/LogIn_Project/http/decode"
	"github.com/namle133/LogIn.git/LogIn_Project/http/encode"
	"github.com/namle133/LogIn.git/LogIn_Project/service"
	"net/http"
)

func main() {
	r := gin.Default()
	p := &service.Product{Db: database.ConnectDatabase()}
	var i service.IProduct = p
	r.POST("/signin", func(c *gin.Context) {
		creds := decode.SignInRequest(c)
		claims, err := i.SignIn(c, creds)
		if err != nil {
			c.String(http.StatusBadRequest, "Cannot sign in")
			return
		}
		encode.SignInResponse(c, claims)
	})
	r.POST("/signup", func(c *gin.Context) {
		creds := decode.SignUpRequest(c)
		err := i.SignUp(c, creds)
		if err != nil {
			c.String(http.StatusBadRequest, "Cannot sign up")
			return
		}
		encode.SignUpResponse(c)
	})
	r.POST("/logout", func(c *gin.Context) {
		err := i.LogOut(c)
		if err != nil {
			c.String(http.StatusBadRequest, "Cannot log out")
			return
		}
		encode.LogoutResponse(c)
	})
	r.Run(":8000")
}
