package repository

import (
	"github.com/go-graphql/config"
	"github.com/go-graphql/models"
)

// TypeRepositoryInterface - .
type TypeRepositoryInterface interface {
	FindAll(typeID *int) ([]*models.Type, error)
	FindAllByPokemonID(pokemonID int) ([]*models.Type, error)
}

// TypeRepository -
type TypeRepository struct {
	DB *config.Database
}

// NewTypeRepository -
func NewTypeRepository(DB *config.Database) TypeRepositoryInterface {
	return &TypeRepository{
		DB: DB,
	}
}

func (p *TypeRepository) queryBuilder(baseQuery string, typeID *int) (string, []interface{}) {
	var vals []interface{}
	if typeID != nil {
		baseQuery += "WHERE id = ?"
		vals = append(vals, *typeID)
	}

	return baseQuery, vals
}

// FindAll -
func (p *TypeRepository) FindAll(typeID *int) ([]*models.Type, error) {
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
	rows, err := p.DB.Query("SELECT t.* from types t JOIN pokemon_types pt ON t.id=pt.type_id where pt.pokemon_id = ?", pokemonID)
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
