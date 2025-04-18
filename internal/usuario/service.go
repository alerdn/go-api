package usuario

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Cadastrar(usuario Usuario) (Usuario, error) {
	if err := Validar(usuario); err != nil {
		return Usuario{}, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)
	if err != nil {
		return Usuario{}, err
	}
	usuario.Senha = string(hashed)

	criado, err := Criar(usuario)
	if err != nil {
		return Usuario{}, errors.New("erro ao cadastrar: " + err.Error())
	}

	return criado, nil
}

func BuscarPerfil(userID int) (Usuario, error) {
	usuario, err := BuscarPorID(userID)
	if err != nil {
		return Usuario{}, errors.New("erro ao buscar perfil: " + err.Error())
	}

	return *usuario, nil
}
