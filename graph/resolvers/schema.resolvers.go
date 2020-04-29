package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golang-graphql-server/graph/generated"
	"golang-graphql-server/graph/generated/model"
)

func (r *mutationResolver) AddEducation(ctx context.Context, education model.EducationInput, style model.StyleInput) (*model.Education, error) {
	return r.resolveAddEducation(ctx, education, style)
}

func (r *mutationResolver) UpdateEducation(ctx context.Context, id string, education model.EducationInput) (*model.Education, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteEducation(ctx context.Context, id string, styleID string) (*model.Education, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateStyle(ctx context.Context, id string, style model.StyleInput) (*model.Style, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllEducations(ctx context.Context) ([]*model.Education, error) {
	return r.resolveAllEducations(ctx)
}

func (r *queryResolver) GetEducationByID(ctx context.Context, id string) (*model.Education, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
