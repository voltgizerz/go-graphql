package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/go-graphql/graph/generated"
	"github.com/go-graphql/models"
	"github.com/vektah/gqlparser/gqlerror"
)

// Types is the resolver for the types field.
func (r *pokemonResolver) Types(ctx context.Context, obj *models.Pokemon) ([]*models.Type, error) {
	id, err := strconv.Atoi(obj.ID)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	res, err := r.Resolver.TypeService.FetchAll(ctx, id)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	return res, nil
}

// Pokemon is the resolver for the pokemon field.
func (r *queryResolver) Pokemon(ctx context.Context, pokemonID int) (*models.Pokemon, error) {
	return r.Resolver.PokemonService.FetchOne(ctx, pokemonID)
}

// Pokemon returns generated.PokemonResolver implementation.
func (r *Resolver) Pokemon() generated.PokemonResolver { return &pokemonResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type pokemonResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
