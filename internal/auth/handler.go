package auth

import (
	"net/http"

	"github.com/alerdn/go-api/internal/usuario"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func LoginHandler(c *gin.Context) {
	var req loginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	u, err := usuario.BuscarPorEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro interno"})
		return
	}
	if u == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "E-mail ou senha inválidos"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Senha), []byte(req.Senha)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "E-mail ou senha inválidos"})
		return
	}

	token, err := GerarToken(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
