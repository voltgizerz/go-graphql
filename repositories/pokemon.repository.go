package repositories

import (
	"database/sql"
	"errors"

	db "github.com/go-graphql/database"
	"github.com/go-graphql/models"
)

// PokemonRepository -
type PokemonRepository struct {
	DB *db.Database
}

// ProvidePokemonRepo -
func ProvidePokemonRepo(DB *db.Database) PokemonRepository {
	return PokemonRepository{
		DB: DB,
	}
}

func (p *PokemonRepository) queryBuilder(baseQuery string, limit *int, offset *int) (string, []interface{}) {
	var vals []interface{}
	if limit != nil {
		baseQuery += "LIMIT ? "
		vals = append(vals, *limit)
	}

	if offset != nil {
		baseQuery += "OFFSET ? "
		vals = append(vals, *offset)
	}
	return baseQuery, vals
}

// FindAll -
func (p *PokemonRepository) FindAll(limit *int, offset *int) ([]*models.Pokemon, error) {
	query, vals := p.queryBuilder("SELECT * from pokemons ", limit, offset)
	rows, err := p.DB.Query(query, vals...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pokemons []*models.Pokemon
	for rows.Next() {
		var poke models.Pokemon
		if err := rows.Scan(&poke.ID, &poke.Name, &poke.Height, &poke.Weight, &poke.BaseExperience); err != nil {
			return nil, err
		}
		pokemons = append(pokemons, &poke)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return pokemons, nil
}

// FindByID -
func (p *PokemonRepository) FindByID(pokemonID int) (*models.Pokemon, error) {
	var pokemon models.Pokemon
	row := p.DB.QueryRow("SELECT * from pokemons where id = ?", pokemonID)
	if err := row.Scan(&pokemon.ID, &pokemon.Name, &pokemon.Height, &pokemon.Weight, &pokemon.BaseExperience); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return &pokemon, nil
}
