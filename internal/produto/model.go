package produto

type Produto struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome" validate:"required"`
	Preco float64 `json:"preco" validate:"required"`
}
