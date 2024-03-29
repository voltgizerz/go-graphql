package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"

	"github.com/go-graphql/internal/app/auth"
	"github.com/go-graphql/internal/app/models"
	"github.com/go-graphql/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/gqlerror"
)

// DeletePokemon is the resolver for the deletePokemon field.
func (r *mutationResolver) DeletePokemon(ctx context.Context, input models.DeletePokemonInput) (*models.DeletePokemonPayload, error) {
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
	}).Info("DeletePokemon mutationResolver")

	err = r.Resolver.PokemonService.Delete(ctx, input.ID)
	if err != nil {
		return nil, gqlerror.Errorf(err.Error())
	}

	return &models.DeletePokemonPayload{
		Success: true,
	}, nil
}
