package resolvers

import (
	"context"

	"github.com/seriousben/simple-graphql-chat/internal/models"
)

type QueryResolver struct{ *RootResolver }

func (r *QueryResolver) Channels(ctx context.Context) ([]*models.Channel, error) {
	panic("not implemented")
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
