package service

import (
	"github.com/go-graphql/config"
	"github.com/go-graphql/models"
	"github.com/go-graphql/repository"
)

// PokemonService -
type PokemonService struct {
	PokemonRepo repository.PokemonRepository
	ID          int
	Limit       *int
	Offset      *int
}

// PokemonInterface - .
type PokemonInterface interface {
	FetchOne(DB *config.Database) (*models.Pokemon, error)
	FetchAll(DB *config.Database) ([]*models.Pokemon, error)
	Create(DB *config.Database, input models.CreatePokemonInput) (*models.Pokemon, error)
	Delete(DB *config.Database) error
}

// NewTypeService - .
func NewPokemonService(pokemonRepo repository.PokemonRepository) PokemonInterface {
	return &PokemonService{
		PokemonRepo: pokemonRepo,
	}
}

// FetchOne - fetch one pokemon data
func (poke *PokemonService) FetchOne(DB *config.Database) (*models.Pokemon, error) {
	repoPokemon := repository.NewPokemonRepository(DB)
	pokemon, err := repoPokemon.FindByID(poke.ID)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

// FetchAll - fetch all pokemon data
func (poke *PokemonService) FetchAll(DB *config.Database) ([]*models.Pokemon, error) {
	repoPokemon := repository.NewPokemonRepository(DB)
	pokemons, err := repoPokemon.FindAll(poke.Limit, poke.Offset)
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}

// Create - create a new pokemon
func (poke *PokemonService) Create(DB *config.Database, input models.CreatePokemonInput) (*models.Pokemon, error) {
	repoPokemon := repository.NewPokemonRepository(DB)
	pokemon, err := repoPokemon.Create(input)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

// Delete - delete data pokemon
func (poke *PokemonService) Delete(DB *config.Database) error {
	repoPokemon := repository.NewPokemonRepository(DB)
	_, err := repoPokemon.Delete(poke.ID)
	if err != nil {
		return err
	}
	return nil
}
