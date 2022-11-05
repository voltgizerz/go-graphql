package services

import (
	"github.com/go-graphql/config"
	"github.com/go-graphql/models"
	"github.com/go-graphql/repository"
)

// Pokemon -
type Pokemon struct {
	ID     int
	Limit  *int
	Offset *int
}

// FetchOne - fetch one pokemon data
func (poke *Pokemon) FetchOne(DB *config.Database) (*models.Pokemon, error) {
	repoPokemon := repository.ProvidePokemonRepo(DB)
	pokemon, err := repoPokemon.FindByID(poke.ID)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

// FetchAll - fetch all pokemon data
func (poke *Pokemon) FetchAll(DB *config.Database) ([]*models.Pokemon, error) {
	repoPokemon := repository.ProvidePokemonRepo(DB)
	pokemons, err := repoPokemon.FindAll(poke.Limit, poke.Offset)
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}

// Create - create a new pokemon
func (poke *Pokemon) Create(DB *config.Database, input models.CreatePokemonInput) (*models.Pokemon, error) {
	repoPokemon := repository.ProvidePokemonRepo(DB)
	pokemon, err := repoPokemon.Create(input)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

// Delete - delete data pokemon
func (poke *Pokemon) Delete(DB *config.Database) error {
	repoPokemon := repository.ProvidePokemonRepo(DB)
	err := repoPokemon.Delete(poke.ID)
	if err != nil {
		return err
	}
	return nil
}
