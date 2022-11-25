package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/go-graphql/models"
	"github.com/go-graphql/utils"
	"github.com/vektah/gqlparser/gqlerror"
)

// Types is the resolver for the types field.
func (r *queryResolver) Types(ctx context.Context, typeID *int) ([]*models.Type, error) {
	res, err := r.Resolver.TypeService.FetchAll(ctx, utils.GetSafeInt(typeID))
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	return res, nil
}
