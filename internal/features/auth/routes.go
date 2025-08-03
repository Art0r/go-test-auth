package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(c *gin.Context) {

	tokenString, status := loginService(c)

	switch status {
	case http.StatusOK:
		c.JSON(status, gin.H{"token": tokenString})
		return

	case http.StatusBadRequest:
		c.JSON(status, gin.H{"error": "Corpo da requisição inválido"})
		return

	case http.StatusInternalServerError:
		c.JSON(status, gin.H{"error": "Ocorreu um erro ao codificar o json"})
		return
	}
}
