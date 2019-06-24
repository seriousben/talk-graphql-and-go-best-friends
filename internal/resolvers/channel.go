package resolvers

import (
	"context"

	"github.com/seriousben/simple-graphql-chat/internal/models"
)

type ChannelResolver struct{ *RootResolver }

func (r *ChannelResolver) CreatedBy(ctx context.Context, obj *models.Channel) (*models.User, error) {
	panic("not implemented")
}
func (r *ChannelResolver) Messages(ctx context.Context, obj *models.Channel, after *string, first *int, sortKey *models.MessageSortKey) (*models.MessageConnection, error) {
	return &models.MessageConnection{
		Edges: []*models.MessageEdge{
			{
				Node: &models.Message{
					ID: "2",
				},
			},
		},
	}, nil
}
func (r *ChannelResolver) Members(ctx context.Context, obj *models.Channel, after *string, first *int, sortKey *models.ChannelMemberSortKey) (*models.UserConnection, error) {
	panic("not implemented")
}
