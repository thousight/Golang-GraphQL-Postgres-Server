package resolvers

import (
	"context"
	"fmt"
	"golang-graphql-server/graph/generated/model"
)

func (r *queryResolver) resolveAllEducations(ctx context.Context) ([]*model.Education, error) {
	fmt.Println(r.dbClient)
	fmt.Println(r.redisClient)
	panic(fmt.Errorf("not implemented"))
}
