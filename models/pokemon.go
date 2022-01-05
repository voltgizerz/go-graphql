package models

type Pokemon struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	BaseExperience int     `json:"baseExperience"`
	Types          []*Type `json:"types"`
}
