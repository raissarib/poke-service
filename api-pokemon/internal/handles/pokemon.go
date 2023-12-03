package handles

import (
	"api-pokemon/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

type Forms []struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func Pokemon(app *fiber.App) {
	user := app.Group("/pokemon")

	user.Get("/:name", controllers.GetPokemon)
}
