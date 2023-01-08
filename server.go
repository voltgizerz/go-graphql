package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/go-graphql/config"
	"github.com/go-graphql/logger"
	"github.com/go-graphql/repository"
	"github.com/go-graphql/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-graphql/graph/generated"
	"github.com/go-graphql/graph/resolvers"
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
	signal.Notify(c, os.Interrupt, syscall.SIGTSTP)

	// Set up a goroutine that will run the signal handler.
	// This goroutine will block until a signal is received.
	go handleSignal(c)

	config.LoadENV()

	// initialize database
	db := config.InitDB()

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
	config.InitGQL(srv)
}
