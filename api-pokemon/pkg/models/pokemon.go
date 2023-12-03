package models

type PokemonComplete struct {
	Ability *PokemonData `json:"ability"`
	Form    []*Forms     `json:"form"`
}
