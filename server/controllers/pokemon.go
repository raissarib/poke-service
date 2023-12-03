package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/raissarib/poke-service.git/models"
	"github.com/raissarib/poke-service.git/utils"
)

func GetPokemon(c *fiber.Ctx) error {
	name := c.Params("name")

	apiURL := fmt.Sprintf("https://ex.traction.one/pokedex/pokemon/%s", name)

	body, err := utils.Get(apiURL)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON("Erro ao pegar as informações do pokemon")
	}

	var pokemons []models.Pokemon
	err = json.Unmarshal(body, &pokemons)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Erro ao pegar as informações do pokemon")
	}

	// Retorna as informações da característica como resposta

	var pokemon models.Pokemon
	if len(pokemons) > 0 {
		pokemon = pokemons[0]
	}

	return c.Status(http.StatusOK).JSON(pokemon)
}
