package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.21 DO NOT EDIT.

import (
	"context"

	"github.com/go-graphql/auth"
	"github.com/go-graphql/logger"
	"github.com/go-graphql/models"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/gqlerror"
)

// UpdatePokemon is the resolver for the updatePokemon field.
func (r *mutationResolver) UpdatePokemon(ctx context.Context, input models.UpdatePokemonInput) (*models.UpdatePokemonPayload, error) {
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

	return nil, gqlerror.Errorf("Not Implemented!")
}
