package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func ConectarDB() {
	var err error

	driverName := os.Getenv("DB_DRIVER")
	dns := os.Getenv("DATABASE_URL")
	DB, err = sql.Open(driverName, dns)
	if err != nil {
		log.Fatal("Erro ao abrir conexão no banco de dados: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Erro ao conectar no banco de dados: ", err)
	}

	log.Println("Banco de dados conectado com sucesso!")

	criarTabelaUsuarios()
	criarTabelaProdutos()
}

func criarTabelaUsuarios() {
	sql := `
	CREATE TABLE IF NOT EXISTS usuarios (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nome TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		senha TEXT NOT NULL
	);`

	if _, err := DB.Exec(sql); err != nil {
		log.Fatal("Erro ao criar tabela usuarios:", err)
	}
}

func criarTabelaProdutos() {
	sql := `
	CREATE TABLE IF NOT EXISTS produtos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nome TEXT NOT NULL,
		preco DOUBLE(15,2) NOT NULL
	);`

	if _, err := DB.Exec(sql); err != nil {
		log.Fatal("Erro ao criar tabela produtos:", err)
	}
}
