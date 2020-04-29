package redis

import (
	"golang-graphql-server/graph/utils"

	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
)

// Connect connects server to Redis and returns a connection client
func Connect() redis.Conn {
	conn, err := redis.DialURL(utils.GetEnv("REDIS_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to redis: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Redis connected")
	}

	return conn
}
