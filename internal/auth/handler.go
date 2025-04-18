package auth

import (
	"encoding/json"
	"net/http"

	"github.com/alerdn/go-api/internal/usuario"
	"github.com/alerdn/go-api/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	u, err := usuario.BuscarPorEmail(req.Email)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Erro interno")
		return
	}
	if u == nil {
		response.JSON(w, http.StatusUnauthorized, "E-mail ou senha inválidos")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Senha), []byte(req.Senha)); err != nil {
		response.JSON(w, http.StatusUnauthorized, "E-mail ou senha inválidos")
		return
	}

	token, err := GerarToken(u.ID)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Erro ao gerar token")
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"token": token})
}
