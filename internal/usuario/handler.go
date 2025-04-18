package usuario

import (
	"encoding/json"
	"net/http"

	"github.com/alerdn/go-api/internal/shared"
	"github.com/alerdn/go-api/pkg/response"
)

func CadastrarHandler(w http.ResponseWriter, r *http.Request) {
	var usuario Usuario

	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		response.JSON(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	criado, err := Cadastrar(usuario)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, err.Error())
	}

	response.JSON(w, http.StatusCreated, criado)
}

func ListarHandler(w http.ResponseWriter, r *http.Request) {
	usuarios, err := Listar()
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Erro ao listar usuários")
		return
	}
	response.JSON(w, http.StatusOK, usuarios)
}

func PerfilHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(shared.UsuarioIDKey).(int)

	usuario, err := BuscarPerfil(userID)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, "Erro ao buscar perfil")
		return
	}

	response.JSON(w, http.StatusOK, usuario)
}
