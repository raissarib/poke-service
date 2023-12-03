package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Forms []struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	app := fiber.New()

	app.Get(":name", func(c *fiber.Ctx) error {
		name := c.Params("name")

		// Monta a URL da PokéAPI com base no ID fornecido
		apiURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)

		// Faz uma requisição GET para a PokéAPI
		resposta, err := http.Get(apiURL)
		if err != nil {
			return c.Status(500).JSON("Erro ao fazer a requisição")
		}
		defer resposta.Body.Close()

		body, err := io.ReadAll(resposta.Body)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON("Erro ao fazer requisição")
		}

		var forms Forms
		err = json.Unmarshal(body, &forms)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON("Erro ao decodificar a resposta")
		}

		// Retorna as informações da característica como resposta
		return c.Status(http.StatusOK).JSON(forms)
	})

	// Inicia o servidor na porta 2000
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
