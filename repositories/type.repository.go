package repositories

import (
	db "github.com/go-graphql/database"
	"github.com/go-graphql/models"
)

// TypeRepository -
type TypeRepository struct {
	DB *db.Database
}

// ProvideTypeRepo -
func ProvideTypeRepo(DB *db.Database) TypeRepository {
	return TypeRepository{
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
