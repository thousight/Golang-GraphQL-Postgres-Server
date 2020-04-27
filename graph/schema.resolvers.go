package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golang-graphql-server/graph/generated"
	"golang-graphql-server/graph/generated/model"
)

func (r *mutationResolver) AddEducation(ctx context.Context, organization string, city string, state string, degree string, time string, order int, image string, desc []*string, primaryColor string, secondaryColor string, bannerImage *string) (*model.Education, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEducation(ctx context.Context, id string, organization *string, city *string, state *string, degree *string, time *string, order *int, image *string, desc []*string) (*model.Education, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteEducation(ctx context.Context, id string, styleID string) (*model.Education, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddExperience(ctx context.Context, organization string, city string, state string, title string, time string, order int, image string, desc []*string, primaryColor string, secondaryColor string, bannerImage *string) (*model.Experience, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateExperience(ctx context.Context, id string, organization *string, city *string, state *string, title *string, time *string, order *int, image *string, desc []*string) (*model.Experience, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteExperience(ctx context.Context, id string, styleID string) (*model.Experience, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateStyle(ctx context.Context, id string, primaryColor *string, secondaryColor *string, bannerImage *string) (*model.Style, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllEducations(ctx context.Context) ([]*model.Education, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetEducationByID(ctx context.Context, id string) (*model.Education, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllExperiences(ctx context.Context) ([]*model.Experience, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetExperienceByID(ctx context.Context, id string) (*model.Experience, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
