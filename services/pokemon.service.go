package services

import (
	db "github.com/go-graphql/database"
	"github.com/go-graphql/models"
	"github.com/go-graphql/repositories"
)

type Pokemon struct {
	ID     int
	Limit  *int
	Offset *int
}

func (poke *Pokemon) FetchOne(DB *db.Database) (*models.Pokemon, error) {
	repoPokemon := repositories.ProvidePokemonRepo(DB)
	pokemon, err := repoPokemon.FindByID(poke.ID)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

func (poke *Pokemon) FetchAll(DB *db.Database) ([]*models.Pokemon, error) {
	repoPokemon := repositories.ProvidePokemonRepo(DB)
	pokemons, err := repoPokemon.FindAll(poke.Limit, poke.Offset)
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}

func (poke *Pokemon) Create(DB *db.Database, input models.CreatePokemonInput) (*models.Pokemon, error) {
	repoPokemon := repositories.ProvidePokemonRepo(DB)
	pokemon, err := repoPokemon.Create(input)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

func (poke *Pokemon) Delete(DB *db.Database) error {
	repoPokemon := repositories.ProvidePokemonRepo(DB)
	err := repoPokemon.Delete(poke.ID)
	if err != nil {
		return err
	}
	return nil
}
