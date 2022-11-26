package service

import (
	"context"

	"github.com/go-graphql/models"
	"github.com/go-graphql/repository"
)

// TypeServiceInterface - .
type TypeServiceInterface interface {
	FetchOne(ctx context.Context, typeID int) ([]*models.Type, error)
	FetchAll(ctx context.Context, pokemonID int) ([]*models.Type, error)
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

// FetchOne -
func (t *Type) FetchOne(ctx context.Context, typeID int) ([]*models.Type, error) {
	types, err := t.TypeRepo.FindTypeByID(typeID)
	if err != nil {
		return nil, err
	}

	return types, nil
}

// FetchAll -
func (t *Type) FetchAll(ctx context.Context, pokemonID int) ([]*models.Type, error) {
	types, err := t.TypeRepo.FindAllByPokemonID(pokemonID)
	if err != nil {
		return nil, err
	}

	return types, nil
}
