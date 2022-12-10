package main

import (
	"github.com/go-graphql/config"
	"github.com/go-graphql/repository"
	"github.com/go-graphql/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-graphql/graph/generated"
	"github.com/go-graphql/graph/resolvers"
)

func main() {
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
