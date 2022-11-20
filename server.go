package main

import (
	"net/http"
	"os"

	"github.com/go-graphql/config"
	"github.com/go-graphql/logger"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := resolvers.NewResolver(db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", auth.Middleware(srv))

	logger.Log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	logger.Log.Fatal(http.ListenAndServe(":"+port, nil))
}
