package resolvers

import "github.com/go-graphql/config"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver -
type Resolver struct {
	DB *config.Database
}

// NewResolver - .
func NewResolver(db *config.Database) *Resolver {
	return &Resolver{
		DB: db,
	}
}
