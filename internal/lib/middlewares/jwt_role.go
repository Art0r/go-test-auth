package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/Art0r/go-test-auth/internal/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtRole(role string) gin.HandlerFunc {
	var jwtKey = []byte(os.Getenv("SECRET"))

	return func(c *gin.Context) {

		authorization := c.GetHeader("Authorization")
		tokenString := strings.Replace(authorization, "Bearer ", "", 1)

		claims := &types.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims.Role != role {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user does not have enought permission"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
