package resolvers

import (
	"golang-graphql-server/graph/generated/model"

	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"
)

func (r *queryResolver) resolveAllEducations(ctx context.Context) ([]*model.Education, error) {
	fmt.Println(r.dbClient)
	fmt.Println(r.redisClient)
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) resolveAddEducation(ctx context.Context, education model.EducationInput, style model.StyleInput) (*model.Education, error) {
	var savedStyle model.Style
	styleErr := r.dbClient.
		Insert("styles").
		Columns("id", "primaryColor", "secondaryColor", "accentColor", "image_url").
		Values(uuid.Must(uuid.NewV4()).String(), style.PrimaryColor, style.SecondaryColor, style.AccentColor, style.ImageURL).
		Suffix("RETURNING *").
		Scan(&savedStyle.ID, &savedStyle.PrimaryColor, &savedStyle.SecondaryColor, &savedStyle.AccentColor, &savedStyle.ImageURL)

	if styleErr != nil {
		fmt.Fprintf(os.Stderr, "Failed creating style: %v\n", styleErr)
		panic(styleErr)
	}

	var savedEducation model.Education
	var ignored string
	currentTime := time.Now()
	fmt.Println(education)
	fmt.Println(r.dbClient.Insert("educations").
		Columns("id", "school", "city", "state", "degree", "month_start", "year_start", "month_end", "year_end", "order", "logo_url", "created_at", "updated_at", "style", "descriptions").
		Values(
			uuid.Must(uuid.NewV4()).String(),
			education.School,
			education.City,
			education.State,
			education.Degree,
			education.MonthStart,
			education.YearStart,
			education.MonthEnd,
			education.YearEnd,
			education.Order,
			education.LogoURL,
			currentTime,
			currentTime,
			savedStyle.ID,
			pq.Array(education.Descriptions),
		).
		Suffix("RETURNING *").ToSql())

	eduErr := r.dbClient.
		Insert("educations").
		Columns("id", "school", "city", "state", "degree", "month_start", "year_start", "month_end", "year_end", "order", "logo_url", "created_at", "updated_at", "style", "descriptions").
		Values(
			uuid.Must(uuid.NewV4()).String(),
			education.School,
			education.City,
			education.State,
			education.Degree,
			education.MonthStart,
			education.YearStart,
			education.MonthEnd,
			education.YearEnd,
			education.Order,
			education.LogoURL,
			currentTime,
			currentTime,
			savedStyle.ID,
			pq.Array(education.Descriptions),
		).
		Suffix("RETURNING *").
		Scan(
			&savedEducation.ID,
			&savedEducation.School,
			&savedEducation.City,
			&savedEducation.State,
			&savedEducation.Degree,
			&savedEducation.MonthStart,
			&savedEducation.YearStart,
			&savedEducation.MonthEnd,
			&savedEducation.YearEnd,
			&savedEducation.Order,
			&savedEducation.LogoURL,
			&savedEducation.CreatedAt,
			&savedEducation.UpdatedAt,
			&ignored,
			pq.Array(&savedEducation.Descriptions),
		)

	if eduErr != nil {
		fmt.Fprintf(os.Stderr, "Failed creating education: %v\n", eduErr)
		panic(eduErr)
	}

	savedEducation.Style = &savedStyle

	return &savedEducation, eduErr
}
