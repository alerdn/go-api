package routes

import (
	"github.com/alerdn/go-api/internal/auth"
	"github.com/alerdn/go-api/internal/usuario"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/login", auth.LoginHandler)
		api.POST("/register", usuario.CadastrarHandler)

		protected := api.Group("/")
		protected.Use(auth.JWTMiddleware())
		{
			protected.GET("/usuarios", usuario.ListarHandler)
			protected.GET("/perfil", usuario.PerfilHandler)
		}
	}

	return r
}
