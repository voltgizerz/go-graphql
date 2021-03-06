package services

import (
	db "github.com/go-graphql/database"
	"github.com/go-graphql/models"
	"github.com/go-graphql/repositories"
)

// Pokemon -
type Pokemon struct {
	ID     int
	Limit  *int
	Offset *int
}

// FetchOne - fetch one pokemon data
func (poke *Pokemon) FetchOne(DB *db.Database) (*models.Pokemon, error) {
	repoPokemon := repositories.ProvidePokemonRepo(DB)
	pokemon, err := repoPokemon.FindByID(poke.ID)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

// FetchAll - fetch all pokemon data
func (poke *Pokemon) FetchAll(DB *db.Database) ([]*models.Pokemon, error) {
	repoPokemon := repositories.ProvidePokemonRepo(DB)
	pokemons, err := repoPokemon.FindAll(poke.Limit, poke.Offset)
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}

// Create - create a new pokemon
func (poke *Pokemon) Create(DB *db.Database, input models.CreatePokemonInput) (*models.Pokemon, error) {
	repoPokemon := repositories.ProvidePokemonRepo(DB)
	pokemon, err := repoPokemon.Create(input)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

// Delete - delete data pokemon
func (poke *Pokemon) Delete(DB *db.Database) error {
	repoPokemon := repositories.ProvidePokemonRepo(DB)
	err := repoPokemon.Delete(poke.ID)
	if err != nil {
		return err
	}
	return nil
}
