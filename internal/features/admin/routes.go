package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func protected(c *gin.Context) {
	username := c.MustGet("username").(string)
	role := c.MustGet("role").(string)
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected route!", "user": username, "role": role})
}
