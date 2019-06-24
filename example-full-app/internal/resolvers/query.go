package resolvers

import (
	"context"
	"strconv"

	"github.com/seriousben/talk-graphql/internal/models"
)

type QueryResolver struct{ *RootResolver }

func (r *QueryResolver) Channels(ctx context.Context) ([]*models.Channel, error) {
	dCs, err := r.DB.ListChannels(ctx)
	if err != nil {
		return nil, err
	}

	c := make([]*models.Channel, 0, len(dCs))
	for _, dC := range dCs {
		c = append(c, &models.Channel{
			ID:   strconv.Itoa(dC.ID),
			Name: dC.Name,
		})
	}

	return c, nil
}
func (r *QueryResolver) Channel(ctx context.Context, id string) (*models.Channel, error) {
	panic("not implemented")
}
func (r *QueryResolver) Message(ctx context.Context, id string) (*models.Message, error) {
	panic("not implemented")
}
func (r *QueryResolver) User(ctx context.Context, id string) (*models.User, error) {
	panic("not implemented")
}
