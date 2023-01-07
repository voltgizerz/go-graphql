package config

import (
	"context"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-graphql/auth"
	"github.com/go-graphql/logger"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/gqlerror"
)

// defaultPort - running port
const defaultPort = "8080"

// InitGQL - init gqlgen
func InitGQL(srv *handler.Server) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(auth.Middleware())

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		user, _ := auth.ForContext(ctx)

		logger.Log.WithFields(logrus.Fields{
			"user_id": user.UserID,
			"error":   err,
		}).Error("Panic Recovered!")

		return gqlerror.Errorf("Internal server error!")
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	logger.Log.Infof("connect to http://localhost:%s/ for GraphQL playground", port)

	err := http.ListenAndServe("127.0.0.1:"+port, router)
	if err != nil {
		logger.Log.Fatal(err)
	}
}
