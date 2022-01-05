package services

import (
	db "github.com/go-graphql/database"
	"github.com/go-graphql/models"
	"github.com/go-graphql/repositories"
)

type Type struct {
	ID *int
	PokemonID int
}

func (t *Type) FetchAll(DB *db.Database) ([]*models.Type, error) {
	repoType := repositories.ProvideTypeRepo(DB)
	types, err := repoType.FindAll(t.ID)
	if err != nil {
		return nil, err
	}
	return types, nil
}

func (t *Type) FetchAllByPokemonID(DB *db.Database) ([]*models.Type, error) {
	repoType := repositories.ProvideTypeRepo(DB)
	types, err := repoType.FindAllByPokemonID(t.PokemonID)
	if err != nil {
		return nil, err
	}
	return types, nil
}
