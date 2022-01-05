// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type CreatePokemonInput struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"baseExperience"`
}

type CreatePokemonPayload struct {
	Success bool     `json:"success"`
	Pokemon *Pokemon `json:"pokemon"`
}

type DeletePokemonInput struct {
	ID int `json:"id"`
}

type DeletePokemonPayload struct {
	Success bool `json:"success"`
}

type Pokemon struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"baseExperience"`
}

type PokemonType struct {
	ID        string `json:"id"`
	PokemonID int    `json:"pokemonID"`
	TypeID    int    `json:"typeID"`
	Slot      int    `json:"slot"`
}

type Type struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	DamageTypeID int    `json:"damageTypeID"`
}

type UpdatePokemonInput struct {
	ID             int     `json:"id"`
	Name           *string `json:"name"`
	Height         *int    `json:"height"`
	Weight         *int    `json:"weight"`
	BaseExperience *int    `json:"baseExperience"`
}

type UpdatePokemonPayload struct {
	Success bool     `json:"success"`
	Pokemon *Pokemon `json:"pokemon"`
}
