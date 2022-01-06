package resolvers

import (
	db "github.com/go-graphql/database"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver -
type Resolver struct {
	DB *db.Database
}
