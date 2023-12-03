package models

type Forms struct {
	BaseStats struct {
		Hp      int `json:"hp"`
		Attack  int `json:"attack"`
		Defense int `json:"defense"`
		SpAtk   int `json:"spAtk"`
		SpDef   int `json:"spDef"`
		Speed   int `json:"speed"`
	} `json:"baseStats"`
	Training struct {
		EvYield        string `json:"evYield"`
		CatchRate      string `json:"catchRate"`
		BaseFriendship string `json:"baseFriendship"`
		BaseExp        string `json:"baseExp"`
		GrowthRate     string `json:"growthRate"`
	} `json:"training"`
}
