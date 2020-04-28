package postgres

import (
	"golang-graphql-server/graph/utils"

	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

// Connect connects server to DB and returns a connection client
func Connect() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), utils.GetEnv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Database connected")
	}
	defer conn.Close(context.Background())

	return conn
}
