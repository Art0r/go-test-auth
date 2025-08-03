package auth

import "github.com/gin-gonic/gin"

func HandleAuthRoutes(r *gin.Engine) {
	router := r.Group("/auth")

	router.POST("/login", login)
}
