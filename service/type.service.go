package service

import (
	"github.com/go-graphql/models"
	"github.com/go-graphql/repository"
)

// TypeServiceInterface - .
type TypeServiceInterface interface {
	FetchOne(typeID int) ([]*models.Type, error)
	FetchAll(pokemonID int) ([]*models.Type, error)
}

// Type -
type Type struct {
	TypeRepo repository.TypeRepositoryInterface
}

// NewTypeService - .
func NewTypeService(typeRepo repository.TypeRepositoryInterface) TypeServiceInterface {
	return &Type{
		TypeRepo: typeRepo,
	}
}

// FetchAll -
func (t *Type) FetchOne(typeID int) ([]*models.Type, error) {
	types, err := t.TypeRepo.FindTypeByID(typeID)
	if err != nil {
		return nil, err
	}

	return types, nil
}

// FetchAllByPokemonID -
func (t *Type) FetchAll(pokemonID int) ([]*models.Type, error) {
	types, err := t.TypeRepo.FindAllByPokemonID(pokemonID)
	if err != nil {
		return nil, err
	}

	return types, nil
}
