package decode

import (
	"github.com/gin-gonic/gin"
	"github.com/namle133/LogIn.git/LogIn_Project/domain"
	"net/http"
)

func SignInRequest(c *gin.Context) *domain.User {
	var creds *domain.User
	err := c.BindJSON(&creds)
	if err != nil {
		c.String(http.StatusBadRequest, "%v", err)
		return nil
	}
	return creds
}

func SignUpRequest(c *gin.Context) *domain.User {
	var creds *domain.User
	er := c.BindJSON(&creds)

	if er != nil {
		c.String(http.StatusBadRequest, "%v", er)
		return nil
	}
	return creds
}
