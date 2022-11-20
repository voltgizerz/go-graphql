package resolvers

import (
	"github.com/go-graphql/config"
	"github.com/go-graphql/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver -
type Resolver struct {
	DB             *config.Database
	PokemonService service.PokemonServiceInterface
	TypeService    service.TypeServiceInterface
}
