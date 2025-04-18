package usuario

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func Validar(usuario Usuario) error {
	return validate.Struct(usuario)
}
