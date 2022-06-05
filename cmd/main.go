package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/namle133/LogIn.git/LogIn_Project/database"
	"github.com/namle133/LogIn.git/LogIn_Project/http/decode"
	"github.com/namle133/LogIn.git/LogIn_Project/http/encode"
	"github.com/namle133/LogIn.git/LogIn_Project/service"
	"net/http"
)

func main() {
	r := gin.Default()
	us := &service.UserService{Db: database.ConnectDatabase()}
	var i service.IUser = us
	err := us.UserAdmin()
	if err != nil {
		fmt.Sprintf("can't create useradmin with err: %v", err)
		return
	}

	r.POST("/signin", func(c *gin.Context) {
		u := decode.InputUser(c)
		claims, err := i.SignIn(c, u)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %v", err))
			return
		}
		encode.SignInResponse(c, claims)
	})

	r.POST("/createuser", func(c *gin.Context) {
		err := i.CheckRowToken(c)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %v", err))
			return
		}
		u := decode.InputUser(c)
		er := i.CreateUser(c, u)
		if er != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %v", err))
			return
		}
		encode.CreateUserResponse(c)
	})

	r.DELETE("/logout", func(c *gin.Context) {
		err := i.LogOut(c)
		if err != nil {
			c.String(http.StatusBadRequest, "LogOut Failed")
			return
		}
		encode.LogoutResponse(c)
	})
	r.Run(":8000")
}
