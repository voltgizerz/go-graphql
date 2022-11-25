package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/go-graphql/models"
)

func (r *mutationResolver) DeletePokemon(ctx context.Context, input models.DeletePokemonInput) (*models.DeletePokemonPayload, error) {
	err := r.Resolver.PokemonService.Delete(input.ID)
	if err != nil {
		return nil, err
	}

	return &models.DeletePokemonPayload{
		Success: true,
	}, nil
}
