package main

import (
	"os"
	"os/signal"

	"github.com/go-graphql/config"
	"github.com/go-graphql/database"
	"github.com/go-graphql/internal/app/repository"
	"github.com/go-graphql/internal/app/service"
	"github.com/go-graphql/pkg/env"
	"github.com/go-graphql/pkg/gqlgen"
	"github.com/go-graphql/pkg/logger"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-graphql/internal/app/graph/generated"
	"github.com/go-graphql/internal/app/graph/resolvers"
)

// Set up a goroutine that will run the signal handler.
// This goroutine will block until a signal is received.
func handleSignal(c chan os.Signal) {
	s := <-c
	logger.Log.Warn("got signal: ", s)

	os.Exit(0)
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Set up a goroutine that will run the signal handler.
	// This goroutine will block until a signal is received.
	go handleSignal(c)

	env.LoadENV()
	_ = config.NewConfig()

	// initialize database
	db, err := database.InitDB()
	if err != nil {
		logger.Log.Error(err)
	}
	// initialize repository
	pokemonRepo := repository.NewPokemonRepository(db)
	typeRepo := repository.NewTypeRepository(db)

	// initialize service
	authService := service.NewAuthService()

	pokemonService := service.NewPokemonService(pokemonRepo)
	typeService := service.NewTypeService(typeRepo)

	// Access needed on resolvers
	resolverData := resolvers.ResolverData{
		PokemonService: pokemonService,
		TypeService:    typeService,
		AuthService:    authService,
	}

	resolver := resolvers.NewResolver(resolverData)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))

	// initialize GQL
	gqlgen.InitGQL(srv)
}
