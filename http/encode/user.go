package encode

import (
	"github.com/gin-gonic/gin"
	"github.com/namle133/LogIn.git/LogIn_Project/domain"
	"net/http"
)

func SignInResponse(c *gin.Context, claims *domain.Claims) {
	c.String(http.StatusOK, "Welcome to  %v", claims.Username)
}

func SignUpResponse(c *gin.Context) {
	c.String(http.StatusOK, "%s", "SignUp Successfully!")
}

func LogoutResponse(c *gin.Context) {
	c.String(http.StatusOK, "%s", "Old cookie deleted. Logged out!")
}
