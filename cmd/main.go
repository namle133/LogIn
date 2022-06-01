package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/namle133/LogIn.git/LogIn_Project/database"
	"github.com/namle133/LogIn.git/LogIn_Project/http/decode"
	"github.com/namle133/LogIn.git/LogIn_Project/http/encode"
	"github.com/namle133/LogIn.git/LogIn_Project/service"
	"github.com/namle133/LogIn.git/LogIn_Project/transport"
	"net/http"
)

func main() {
	r := gin.Default()
	us := &service.UserService{Db: database.ConnectDatabase()}
	var i service.IUser = us
	r.POST("/signin", func(c *gin.Context) {
		u := decode.InputUser(c)
		claims, tkStr, err := i.SignIn(c, u)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %v", err))
			return
		}
		transport.SetCookieUser(tkStr, c)
		encode.SignInResponse(c, claims)
	})
	r.POST("/createuser", func(c *gin.Context) {
		ck, err := transport.GetCookieUser(c)
		if err != nil || ck == "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %v", err))
			return
		}
		u := decode.InputUser(c)
		er := i.CreateUser(c, u)
		if er != nil {
			c.String(http.StatusBadRequest, "Cannot create user")
			return
		}
		encode.CreateUserResponse(c)
	})
	//r.POST("/logout", func(c context.Context) {
	//	err := i.LogOut(c)
	//	if err != nil {
	//		c.String(http.StatusBadRequest, "Cannot log out")
	//		return
	//	}
	//	encode.LogoutResponse(c)
	//})
	r.Run(":8000")
}
