package services

import (
	"api-pokemon/pkg/models"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

const urlApiPokemonDetails = "http://localhost:5091/Pokemon/"

func RequisitionToPokemonDetails(name string) ([]*models.Forms, error) {
	client := &http.Client{}

	var queryBuilder strings.Builder
	queryBuilder.WriteString(urlApiPokemonDetails + name)

	var forms []*models.Forms
	req, err := http.NewRequest("GET", queryBuilder.String(), nil)

	if err != nil {
		return nil, errors.New("error to create request for pokemon details")
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, errors.New("error to send request for pokemon details")
	}
	defer res.Body.Close()

	bodyRes, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, errors.New("error to read body of pokemon details")
	}

	if err = json.Unmarshal(bodyRes, &forms); err != nil {
		return nil, errors.New("error to decode response of pokemon details")
	}

	return forms, nil
}
