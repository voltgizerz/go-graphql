package resolvers

import (
	"github.com/go-graphql/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// ResolverData -
type ResolverData struct {
	PokemonService service.PokemonServiceInterface
	TypeService    service.TypeServiceInterface
	// Other interfaces needed....
}

// NewResolver -
func NewResolver(data ResolverData) *Resolver {
	return &Resolver{
		PokemonService: data.PokemonService,
		TypeService:    data.TypeService,
	}
}

// Resolver -
type Resolver struct {
	PokemonService service.PokemonServiceInterface
	TypeService    service.TypeServiceInterface
}
