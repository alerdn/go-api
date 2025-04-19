package produto

import "github.com/alerdn/go-api/config"

func Criar(produto Produto) (Produto, error) {

	config.DB.Prepare("INSERT INTO produtos (nome, preco) VALUES (?, ?)")
	

	return produto, nil
}