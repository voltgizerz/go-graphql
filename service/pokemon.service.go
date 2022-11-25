package service

import (
	"github.com/go-graphql/models"
	"github.com/go-graphql/repository"
)

// PokemonServiceInterface - .
type PokemonServiceInterface interface {
	FetchOne(id int) (*models.Pokemon, error)
	FetchAll(limit, offset int) ([]*models.Pokemon, error)
	Create(input models.CreatePokemonInput) (*models.Pokemon, error)
	Delete(id int) error
}

// PokemonService -
type PokemonService struct {
	PokemonRepo repository.PokemonRepositoryInterface
}

// NewPokemonService - .
func NewPokemonService(pokemonRepo repository.PokemonRepositoryInterface) PokemonServiceInterface {
	return &PokemonService{
		PokemonRepo: pokemonRepo,
	}
}

// FetchOne - fetch one pokemon data
func (p *PokemonService) FetchOne(id int) (*models.Pokemon, error) {
	pokemon, err := p.PokemonRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}

// FetchAll - fetch all pokemon data
func (p *PokemonService) FetchAll(limit, offset int) ([]*models.Pokemon, error) {
	pokemons, err := p.PokemonRepo.FetchAllPokemonData(limit, offset)
	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

// Create - create a new pokemon
func (p *PokemonService) Create(input models.CreatePokemonInput) (*models.Pokemon, error) {
	pokemon, err := p.PokemonRepo.Create(input)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}

// Delete - delete data pokemon
func (p *PokemonService) Delete(id int) error {
	_, err := p.PokemonRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
