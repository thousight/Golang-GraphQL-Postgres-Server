package resolvers

import (
	"golang-graphql-server/graph/generated"
	"golang-graphql-server/graph/services/postgres"
	"golang-graphql-server/graph/services/redis"

	redisLib "github.com/gomodule/redigo/redis"
	"github.com/jackc/pgx/v4"
)

// Resolver serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	dbClient    pgx.Conn
	redisClient redisLib.Conn
}

// InitializeResolver creates a new resolver context and initialize service connections and clients
func InitializeResolver() generated.Config {
	resolver := Resolver{}

	resolver.dbClient = *postgres.Connect()
	resolver.redisClient = redis.Connect()

	return generated.Config{
		Resolvers: &resolver,
	}
}
