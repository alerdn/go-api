package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alerdn/go-api/config"
	"github.com/alerdn/go-api/routes"
)

func main() {
	config.CarregarEnv()
	config.ConectarDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := routes.SetupRoutes()

	fmt.Printf("Servidor rodando em http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
