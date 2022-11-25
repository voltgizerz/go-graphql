package main

import (
	"context"
	"net/http"
	"os"

	"github.com/go-graphql/config"
	"github.com/go-graphql/logger"
	"github.com/go-graphql/repository"
	"github.com/go-graphql/service"
	"github.com/vektah/gqlparser/gqlerror"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-graphql/auth"
	"github.com/go-graphql/graph/generated"
	"github.com/go-graphql/graph/resolvers"
)

const defaultPort = "8080"

func main() {
	config.LoadENV()

	// initialize database
	db := config.InitDB()

	// initialize repository
	pokemonRepo := repository.NewPokemonRepository(db)
	typeRepo := repository.NewTypeRepository(db)

	// initialize service
	pokemonService := service.NewPokemonService(pokemonRepo)
	typeService := service.NewTypeService(typeRepo)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{
			PokemonService: pokemonService,
			TypeService:    typeService,
		},
	}))

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		logger.Log.Error(err)
		return gqlerror.Errorf("Internal server error!")
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", auth.Middleware(srv))

	logger.Log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	logger.Log.Fatal(http.ListenAndServe(":"+port, nil))
}
