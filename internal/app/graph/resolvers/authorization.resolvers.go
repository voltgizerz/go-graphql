package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"

	"github.com/go-graphql/internal/app/graph/generated"
	"github.com/go-graphql/internal/app/models"
	"github.com/go-graphql/logger"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/gqlerror"
)

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, input models.LoginInput) (*models.LoginPayload, error) {
	token, err := r.AuthService.Login(ctx, input)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Login")
		return nil, gqlerror.Errorf(err.Error())
	}

	res := &models.LoginPayload{
		Success: true,
		Message: "Login Success",
		Token:   token,
	}

	return res, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }