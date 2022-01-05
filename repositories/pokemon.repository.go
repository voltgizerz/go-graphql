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
		var p models.Pokemon
		if err := rows.Scan(&p.ID, &p.Name, &p.Height, &p.Weight, &p.BaseExperience); err != nil {
			return nil, err
		}
		pokemons = append(pokemons, &p)
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

// Create  -
func (p *PokemonRepository) Create(pokemonID int) error {
	_, err := p.DB.Exec("INSERT INTO from pokemons where id = ?", pokemonID)
	if err != nil {
		return err
	}
	return nil
}

// Update  -
func (p *PokemonRepository) Update(pokemonID int) error {
	_, err := p.DB.Exec("UPDATE from pokemons where id = ?", pokemonID)
	if err != nil {
		return err
	}
	return nil
}

// Delete  -
func (p *PokemonRepository) Delete(pokemonID int) error {
	_, err := p.DB.Exec("DELETE from pokemons where id = ?", pokemonID)
	if err != nil {
		return err
	}
	return nil
}
