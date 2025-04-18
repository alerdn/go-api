package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConectarDB() {
	var err error

	dns := os.Getenv("DATABASE_URL")
	DB, err = sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("Erro ao abrir conexão no banco de dados: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Erro ao conectar no banco de dados: ", err)
	}

	log.Println("Banco de dados conectado com sucesso!")
}
