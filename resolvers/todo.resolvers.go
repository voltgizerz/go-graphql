package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"log"

	"github.com/apollo-graphql/graph/generated"
	"github.com/apollo-graphql/graph/models"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	log.Println("ASdssd")
	aaa:= &models.Todo{ID: "1"}
	a := []*models.Todo{aaa,aaa,aaa}
	return a, nil
}

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	log.Println(obj)
	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
