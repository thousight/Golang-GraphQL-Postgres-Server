package postgres

import (
	"golang-graphql-server/graph/utils"

	"fmt"
	"os"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
)

// Connect connects server to DB and returns a connection client
func Connect() squirrel.StatementBuilderType {
	config, err := pgx.ParseURI(utils.GetEnv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to database URI: %v\n", err)
		os.Exit(1)
	}

	db := stdlib.OpenDB(config)

	fmt.Println("Database connected")

	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db)
}
