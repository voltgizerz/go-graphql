package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.21 DO NOT EDIT.

import (
	"context"

	"github.com/go-graphql/auth"
	"github.com/go-graphql/logger"
	"github.com/go-graphql/models"
	"github.com/go-graphql/utils"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/gqlerror"
)

// Types is the resolver for the types field.
func (r *queryResolver) Types(ctx context.Context, typeID *int) ([]*models.Type, error) {
	user, err := auth.ForContext(ctx)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("User Authorization")
		return nil, gqlerror.Errorf(err.Error())
	}

	logger.Log.WithFields(logrus.Fields{
		"user_id":  user.UserID,
		"is_admin": user.IsAdmin,
	}).Info("Types queryResolver")

	res, err := r.Resolver.TypeService.FetchAll(ctx, utils.GetSafeInt(typeID))
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	return res, nil
}
