package controllers

import (
	"api-pokemon/internal/services"
	"api-pokemon/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func GetPokemon(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	pokemonData, err := services.RequisitionToPokemonData(name)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("error to making request for pokemon data")
	}

	pokemonDetails, err := services.RequisitionToPokemonDetails(name)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("error to making request for pokemon details")
	}

	var pokemonComplete models.PokemonComplete

	pokemonComplete.Ability = pokemonData
	pokemonComplete.Form = pokemonDetails

	return ctx.Status(fiber.StatusOK).JSON(pokemonComplete)
}
