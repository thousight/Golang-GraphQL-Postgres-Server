package main

import (
	"golang-graphql-server/graph/generated"
	"golang-graphql-server/graph/resolvers"
	"golang-graphql-server/graph/utils"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	utils.InitDotEnv()

	port := os.Getenv("PORT")

	server := handler.NewDefaultServer(generated.NewExecutableSchema(resolvers.InitializeResolver()))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
