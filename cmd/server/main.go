package main

import (
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

	r.Run(":" + port)
}
