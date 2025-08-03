package main

import (
	"log"

	"github.com/Art0r/go-test-auth/internal/features/admin"
	"github.com/Art0r/go-test-auth/internal/features/auth"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"time"

	"github.com/gin-contrib/cors"
)

func setCorsConfig(r *gin.Engine) {

	config := cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTION"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Origin", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	setCorsConfig(r)
	auth.HandleAuthRoutes(r)
	admin.HandleAdminRoutes(r)

	r.Run()
}
