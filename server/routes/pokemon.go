package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raissarib/poke-service.git/server/controllers"
)

func Pokemon(app *fiber.App) {
	group := app.Group("pokemon/")

	group.Get(":id", controllers.GetPokemon)

}
