package main

import (
	"api-pokemon/internal/handles"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	handles.Pokemon(app)

	log.Fatal(app.Listen(":8080"))
}
