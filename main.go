package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/raissarib/poke-service.git/server/routes"
)

type Forms []struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	app := fiber.New()

	routes.Pokemon(app)

	// Inicia o servidor na porta 2000
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
