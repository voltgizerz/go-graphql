package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/go-graphql/graph/generated"
	"github.com/go-graphql/models"
	"github.com/go-graphql/services"
)

func (r *pokemonResolver) Types(ctx context.Context, obj *models.Pokemon) ([]*models.Type, error) {
	id, _ := strconv.Atoi(obj.ID)
	service := services.Type{PokemonID: id}
	return service.FetchAllByPokemonID(r.DB)
}

func (r *queryResolver) Pokemon(ctx context.Context, pokemonID int) (*models.Pokemon, error) {
	service := services.Pokemon{ID: pokemonID}
	return service.FetchOne(r.DB)
}

// Pokemon returns generated.PokemonResolver implementation.
func (r *Resolver) Pokemon() generated.PokemonResolver { return &pokemonResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type pokemonResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
