package service

import (
	"github.com/go-graphql/models"
	"github.com/go-graphql/repository"
)

// Type -
type Type struct {
	TypeRepo repository.TypeRepositoryInterface
}

// TypeServiceInterface - .
type TypeServiceInterface interface {
	FetchAll(typeID int) ([]*models.Type, error)
	FetchAllByPokemonID(pokemonID int) ([]*models.Type, error)
}

// NewTypeService - .
func NewTypeService(typeRepo repository.TypeRepositoryInterface) TypeServiceInterface {
	return &Type{
		TypeRepo: typeRepo,
	}
}

// FetchAll -
func (t *Type) FetchAll(typeID int) ([]*models.Type, error) {
	types, err := t.TypeRepo.FindAll(typeID)
	if err != nil {
		return nil, err
	}
	return types, nil
}

// FetchAllByPokemonID -
func (t *Type) FetchAllByPokemonID(pokemonID int) ([]*models.Type, error) {
	types, err := t.TypeRepo.FindAllByPokemonID(pokemonID)
	if err != nil {
		return nil, err
	}
	return types, nil
}
