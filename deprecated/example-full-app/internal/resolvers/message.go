package resolvers

import (
	"context"

	"github.com/seriousben/talk-graphql/internal/models"
)

type MessageResolver struct{ *RootResolver }

func (r *MessageResolver) CreatedBy(ctx context.Context, obj *models.Message) (*models.User, error) {
	panic("not implemented")
}
