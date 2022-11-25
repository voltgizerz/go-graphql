package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/go-graphql/models"
	"github.com/go-graphql/utils"
)

func (r *queryResolver) Types(ctx context.Context, typeID *int) ([]*models.Type, error) {
	return r.Resolver.TypeService.FetchAll(utils.GetSafeInt(typeID))
}
