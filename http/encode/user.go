package encode

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/namle133/LogIn.git/LogIn_Project/domain"
	"net/http"
)

func SignInResponse(c *gin.Context, claims *domain.Claims) {
	c.String(http.StatusOK, fmt.Sprintf("Welcome to %v", claims.Username))
}

func CreateUserResponse(c *gin.Context) {
	c.String(http.StatusOK, "SignUp Success")
}

func LogoutResponse(c *gin.Context) {
	c.String(http.StatusOK, "Old cookie deleted. Logged out!")
}
