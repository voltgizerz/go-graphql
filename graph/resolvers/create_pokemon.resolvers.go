package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/go-graphql/graph/generated"
	"github.com/go-graphql/models"
)

func (r *mutationResolver) CreatePokemon(ctx context.Context, input models.CreatePokemonInput) (*models.CreatePokemonPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
