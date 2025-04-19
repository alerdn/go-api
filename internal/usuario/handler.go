package usuario

import (
	"net/http"

	"github.com/alerdn/go-api/internal/shared"
	"github.com/gin-gonic/gin"
)

func CadastrarHandler(c *gin.Context) {
	var usuario Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	criado, err := Cadastrar(usuario)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, criado)
}

func ListarHandler(c *gin.Context) {
	usuarios, err := Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao listar usuários"})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

func PerfilHandler(c *gin.Context) {
	userID := c.GetInt(shared.UsuarioIDKey)

	usuario, err := BuscarPerfil(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar perfil"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}
