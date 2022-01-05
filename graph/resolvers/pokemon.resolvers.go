package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/go-graphql/graph/generated"
	"github.com/go-graphql/models"
)

func (r *queryResolver) Pokemon(ctx context.Context, pokemonID *int) (*models.Pokemon, error) {
	return &models.Pokemon{ID: "1", Name: "Bulbasaur"}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
