package admin

import (
	"github.com/Art0r/go-test-auth/internal/lib/middlewares"
	"github.com/gin-gonic/gin"
)

func HandleAdminRoutes(r *gin.Engine) {
	protectedRouter := r.Group("/admin", middlewares.JwtRole("admin"))

	protectedRouter.GET("/", protected)
}
