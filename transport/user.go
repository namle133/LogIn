package transport

import (
	"github.com/gin-gonic/gin"
)

func SetCookieUser(tknStr string, c *gin.Context) {
	c.SetCookie("token", tknStr, 3600, "/", "localhost", false, true)
}

func GetCookieUser(c *gin.Context) (string, error) {
	ck, err := c.Cookie("token")
	if err != nil {
		return "", err
	}
	return ck, nil
}
