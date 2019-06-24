package resolvers

import (
	"context"

	"github.com/seriousben/simple-graphql-chat/internal/models"
)

type MessageResolver struct{ *RootResolver }

func (r *MessageResolver) CreatedBy(ctx context.Context, obj *models.Message) (*models.User, error) {
	panic("not implemented")
}
