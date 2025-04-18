package routes

import (
	"net/http"

	"github.com/alerdn/go-api/internal/auth"
	"github.com/alerdn/go-api/internal/usuario"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/api/v1/", func(r chi.Router) {
		r.Post("/login", auth.LoginHandler)
		r.Post("/register", usuario.CadastrarHandler)

		r.Group(func(r chi.Router) {
			r.Use(auth.JWTMiddleware)

			r.Get("/usuarios", usuario.ListarHandler)
			r.Get("/perfil", usuario.PerfilHandler)
		})
	})

	return r
}
