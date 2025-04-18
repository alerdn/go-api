package config

import (
	"log"

	"github.com/joho/godotenv"
)

func CarregarEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Erro ao carregar o arquivo .env")
	}
}
