package repository

import (
	"github.com/go-graphql/internal/app/models"
	"github.com/go-graphql/database"
)

// TypeRepositoryInterface - .
type TypeRepositoryInterface interface {
	FindTypeByID(typeID int) ([]*models.Type, error)
	FindAllByPokemonID(id int) ([]*models.Type, error)
}

// TypeRepository -
type TypeRepository struct {
	DB *database.Database
}

// NewTypeRepository -
func NewTypeRepository(DB *database.Database) TypeRepositoryInterface {
	return &TypeRepository{
		DB: DB,
	}
}

func (p *TypeRepository) queryBuilder(baseQuery string, typeID int) (string, []interface{}) {
	var vals []interface{}
	if typeID != 0 {
		baseQuery += "WHERE id = ?"
		vals = append(vals, typeID)
	}

	return baseQuery, vals
}

// FindTypeByID -
func (p *TypeRepository) FindTypeByID(typeID int) ([]*models.Type, error) {
	query, vals := p.queryBuilder("SELECT * from types ", typeID)
	rows, err := p.DB.Query(query, vals...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []*models.Type
	for rows.Next() {
		var t models.Type
		if err := rows.Scan(&t.ID, &t.Name, &t.DamageTypeID); err != nil {
			return nil, err
		}
		types = append(types, &t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return types, nil
}

// FindAllByPokemonID -
func (p *TypeRepository) FindAllByPokemonID(pokemonID int) ([]*models.Type, error) {
	findAllDataByPokemonID := "SELECT t.* from types t JOIN pokemon_types pt ON t.id=pt.type_id where pt.pokemon_id = ?"
	rows, err := p.DB.Query(findAllDataByPokemonID, pokemonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []*models.Type
	for rows.Next() {
		var t models.Type
		if err := rows.Scan(&t.ID, &t.Name, &t.DamageTypeID); err != nil {
			return nil, err
		}
		types = append(types, &t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return types, nil
}
