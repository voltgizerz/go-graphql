package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/go-graphql/models"
	"github.com/vektah/gqlparser/gqlerror"
)

// DeletePokemon is the resolver for the deletePokemon field.
func (r *mutationResolver) DeletePokemon(ctx context.Context, input models.DeletePokemonInput) (*models.DeletePokemonPayload, error) {
	err := r.Resolver.PokemonService.Delete(ctx, input.ID)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	return &models.DeletePokemonPayload{
		Success: true,
	}, nil
}
