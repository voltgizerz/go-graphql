package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/go-graphql/models"
	"github.com/go-graphql/utils"
	"github.com/vektah/gqlparser/gqlerror"
)

func (r *queryResolver) Types(ctx context.Context, typeID *int) ([]*models.Type, error) {
	res, err := r.Resolver.TypeService.FetchAll(utils.GetSafeInt(typeID))
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	return res, nil
}
