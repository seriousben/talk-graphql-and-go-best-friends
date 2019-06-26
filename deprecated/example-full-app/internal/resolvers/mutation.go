package resolvers

import (
	"context"

	"github.com/seriousben/talk-graphql/internal/models"
)

type MutationResolver struct{ *RootResolver }

func (r *MutationResolver) CreateChannel(ctx context.Context, input models.NewChannel) (*models.Channel, error) {
	panic("not implemented")
}
func (r *MutationResolver) CreateMessage(ctx context.Context, input models.NewMessage) (*models.Message, error) {
	panic("not implemented")
}
