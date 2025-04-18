package usuario

type Usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Senha string `json:"senha,omitempty" validate:"required,min=6"`
}
