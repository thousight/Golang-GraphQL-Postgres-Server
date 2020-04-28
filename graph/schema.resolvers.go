package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golang-graphql-server/graph/generated"
	"golang-graphql-server/graph/generated/model"
)

func (r *mutationResolver) AddEducation(ctx context.Context, education model.EducationInput, style model.StyleInput) (*model.Education, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteEducation(ctx context.Context, id string, styleID string) (*model.Education, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddExperience(ctx context.Context, experience model.ExperienceInput, style model.StyleInput) (*model.Experience, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteExperience(ctx context.Context, id string, styleID string) (*model.Experience, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateStyle(ctx context.Context, id string, style model.StyleInput) (*model.Style, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllEducations(ctx context.Context) ([]*model.Education, error) {
	fmt.Println(r.dbClient)
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) UpdateEducation(ctx context.Context, id string, education model.EducationInput) (*model.Education, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) UpdateExperience(ctx context.Context, id string, experience model.ExperienceInput) (*model.Experience, error) {
	panic(fmt.Errorf("not implemented"))
}
