package resolvers

import (
	"context"

	"github.com/seriousben/simple-graphql-chat/internal/models"
)

type MutationResolver struct{ *RootResolver }

func (r *MutationResolver) CreateChannel(ctx context.Context, input models.NewChannel) (*models.Channel, error) {
	panic("not implemented")
}
func (r *MutationResolver) CreateMessage(ctx context.Context, input models.NewMessage) (*models.Message, error) {
	panic("not implemented")
}
