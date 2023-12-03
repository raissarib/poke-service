package services

import (
	"api-pokemon/pkg/models"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

const urlApiPokemonData = "http://localhost:3000/pokemon/"

func RequisitionToPokemonData(name string) (*models.PokemonData, error) {
	client := &http.Client{}

	var queryBuilder strings.Builder
	queryBuilder.WriteString(urlApiPokemonData + name)

	var pokemonData *models.PokemonData
	req, err := http.NewRequest("GET", queryBuilder.String(), nil)

	if err != nil {
		return nil, errors.New("error to create request for pokemon data")
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, errors.New("error to send request for pokemon data")
	}
	defer res.Body.Close()

	bodyRes, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, errors.New("error to read body of pokemon data")
	}

	if err = json.Unmarshal(bodyRes, &pokemonData); err != nil {
		return nil, errors.New("error to decode response of pokemon data")
	}

	return pokemonData, nil
}
