package models

type Pokemon struct {
	Number    int      `json:"number"`
	Name      string   `json:"name"`
	Gen       int      `json:"gen"`
	Species   string   `json:"species"`
	Types     []string `json:"types"`
	Abilities []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Hidden      bool   `json:"hidden"`
	} `json:"abilities"`
	Height string `json:"height"`
	Weight string `json:"weight"`
}
