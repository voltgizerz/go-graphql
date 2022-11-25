package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/go-graphql/graph/generated"
	"github.com/go-graphql/models"
	"github.com/vektah/gqlparser/gqlerror"
)

func (r *mutationResolver) CreatePokemon(ctx context.Context, input models.CreatePokemonInput) (*models.CreatePokemonPayload, error) {
	pokemon, err := r.Resolver.PokemonService.Create(input)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	return &models.CreatePokemonPayload{
		Success: true,
		Pokemon: pokemon,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
