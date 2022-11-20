package service

import (
	"github.com/go-graphql/config"
	"github.com/go-graphql/models"
	"github.com/go-graphql/repository"
)

// Type -
type Type struct {
	TypeRepo  repository.TypeRepositoryInterface
	ID        *int
	PokemonID int
}

// TypeServiceInterface - .
type TypeServiceInterface interface {
	FetchAll(DB *config.Database) ([]*models.Type, error)
	FetchAllByPokemonID(DB *config.Database) ([]*models.Type, error)
}

// NewTypeService - .
func NewTypeService(typeRepo repository.TypeRepositoryInterface) TypeServiceInterface {
	return &Type{
		TypeRepo: typeRepo,
	}
}

// FetchAll -
func (t *Type) FetchAll(DB *config.Database) ([]*models.Type, error) {
	repoType := repository.NewTypeRepository(DB)
	types, err := repoType.FindAll(t.ID)
	if err != nil {
		return nil, err
	}
	return types, nil
}

// FetchAllByPokemonID -
func (t *Type) FetchAllByPokemonID(DB *config.Database) ([]*models.Type, error) {
	repoType := repository.NewTypeRepository(DB)
	types, err := repoType.FindAllByPokemonID(t.PokemonID)
	if err != nil {
		return nil, err
	}
	return types, nil
}
