package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/Art0r/go-test-auth/internal/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func loginService(c *gin.Context) (t *string, s int) {
	var jwtKey = []byte(os.Getenv("SECRET"))
	var credentials = &types.Login{}

	if err := c.ShouldBindJSON(credentials); err != nil {
		return nil, http.StatusBadRequest
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &types.Claims{
		Username: credentials.Username,
		Role:     "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, http.StatusInternalServerError
	}

	return &tokenString, http.StatusOK
}
