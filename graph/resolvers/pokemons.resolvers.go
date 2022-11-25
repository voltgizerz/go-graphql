package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/go-graphql/models"
	"github.com/go-graphql/utils"
)

func (r *queryResolver) Pokemons(ctx context.Context, limit *int, offset *int) ([]*models.Pokemon, error) {
	return r.PokemonService.FetchAll(utils.GetSafeInt(limit), utils.GetSafeInt(offset))
}
