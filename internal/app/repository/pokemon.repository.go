package repository

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/go-graphql/internal/app/models"
		"github.com/go-graphql/database"
)

// PokemonRepositoryInterface - .
type PokemonRepositoryInterface interface {
	FetchAllPokemonData(limit int, offset int) ([]*models.Pokemon, error)
	FindByID(id int) (*models.Pokemon, error)
	Create(input models.CreatePokemonInput) (*models.Pokemon, error)
	Update(id int) (int64, error)
	Delete(id int) (int64, error)
}

// PokemonRepository -
type PokemonRepository struct {
	DB *database.Database
}

// NewPokemonRepository -
func NewPokemonRepository(DB *database.Database) PokemonRepositoryInterface {
	return &PokemonRepository{
		DB: DB,
	}
}

func (p *PokemonRepository) queryBuilder(baseQuery string, limit int, offset int) (string, []interface{}) {
	var vals []interface{}
	if limit != 0 {
		baseQuery += "LIMIT ? "
		vals = append(vals, limit)
	}

	if offset != 0 {
		baseQuery += "OFFSET ? "
		vals = append(vals, offset)
	}

	return baseQuery, vals
}

// FetchAllPokemonData -
func (p *PokemonRepository) FetchAllPokemonData(limit int, offset int) ([]*models.Pokemon, error) {
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
func (p *PokemonRepository) FindByID(id int) (*models.Pokemon, error) {
	var pokemon models.Pokemon
	row := p.DB.QueryRow("SELECT * from pokemons where id = ?", id)
	if err := row.Scan(&pokemon.ID, &pokemon.Name, &pokemon.Height, &pokemon.Weight, &pokemon.BaseExperience); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return &pokemon, nil
}

// Create  -
func (p *PokemonRepository) Create(input models.CreatePokemonInput) (*models.Pokemon, error) {
	queryInsert := "INSERT INTO pokemons(name, height, weight, base_experience) VALUES (?,?,?,?)"
	row, err := p.DB.Exec(queryInsert, input.Name, input.Height, input.Weight, input.BaseExperience)
	if err != nil {
		return nil, err
	}

	lastID, err := row.LastInsertId()
	if err != nil {
		return nil, err
	}

	pokemon := &models.Pokemon{
		ID:             strconv.Itoa(int(lastID)),
		Name:           input.Name,
		Height:         input.Height,
		Weight:         input.Weight,
		BaseExperience: input.BaseExperience,
		Types:          []*models.Type{},
	}

	return pokemon, nil
}

// Update  -
func (p *PokemonRepository) Update(id int) (int64, error) {
	res, err := p.DB.Exec("UPDATE from pokemons where id = ?", id)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// Delete  -
func (p *PokemonRepository) Delete(id int) (int64, error) {
	res, err := p.DB.Exec("DELETE from pokemons where id = ?", id)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}
