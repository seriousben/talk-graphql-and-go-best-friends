package resolvers

import (
	"context"

	"github.com/seriousben/talk-graphql/internal/models"
)

type UserResolver struct{ *RootResolver }

func (r *UserResolver) Messages(ctx context.Context, obj *models.User, after *string, first *int, sortKey *models.MessageSortKey) (*models.MessageConnection, error) {
	panic("not implemented")
}
