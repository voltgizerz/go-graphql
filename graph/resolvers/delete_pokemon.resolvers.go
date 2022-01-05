package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/go-graphql/models"
	"github.com/go-graphql/services"
)

func (r *mutationResolver) DeletePokemon(ctx context.Context, input models.DeletePokemonInput) (*models.DeletePokemonPayload, error) {
	service := services.Pokemon{ID: input.ID}
	err := service.Delete(r.DB)
	if err != nil {
		return nil, err
	}
	return &models.DeletePokemonPayload{Success: true}, nil
}
