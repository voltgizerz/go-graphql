package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/go-graphql/graph/generated"
	"github.com/go-graphql/models"
	"github.com/go-graphql/service"
)

func (r *mutationResolver) CreatePokemon(ctx context.Context, input models.CreatePokemonInput) (*models.CreatePokemonPayload, error) {
	srv := service.PokemonService{}
	pokemon, err := srv.Create(r.DB, input)
	if err != nil {
		return nil, err
	}
	return &models.CreatePokemonPayload{Success: true, Pokemon: pokemon}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
