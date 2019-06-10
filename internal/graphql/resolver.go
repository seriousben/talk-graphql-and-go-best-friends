package graphql

import (
	"context"

	"github.com/seriousben/simple-graphql-chat/internal/graphql/graph"
	"github.com/seriousben/simple-graphql-chat/internal/graphql/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Channel() graph.ChannelResolver {
	return &channelResolver{r}
}
func (r *Resolver) Message() graph.MessageResolver {
	return &messageResolver{r}
}
func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() graph.UserResolver {
	return &userResolver{r}
}

type channelResolver struct{ *Resolver }

func (r *channelResolver) CreatedBy(ctx context.Context, obj *models.Channel) (*models.User, error) {
	panic("not implemented")
}
func (r *channelResolver) Messages(ctx context.Context, obj *models.Channel, after *string, first *int, sortKey *models.MessageSortKey) (*models.MessageConnection, error) {
	panic("not implemented")
}
func (r *channelResolver) Members(ctx context.Context, obj *models.Channel, after *string, first *int, sortKey *models.ChannelMemberSortKey) (*models.UserConnection, error) {
	panic("not implemented")
}

type messageResolver struct{ *Resolver }

func (r *messageResolver) CreatedBy(ctx context.Context, obj *models.Message) (*models.User, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateChannel(ctx context.Context, input models.NewChannel) (*models.Channel, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateMessage(ctx context.Context, input models.NewMessage) (*models.Message, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Channels(ctx context.Context) ([]*models.Channel, error) {
	panic("not implemented")
}
func (r *queryResolver) Channel(ctx context.Context, id string) (*models.Channel, error) {
	panic("not implemented")
}
func (r *queryResolver) Message(ctx context.Context, id string) (*models.Message, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) Messages(ctx context.Context, obj *models.User, after *string, first *int, sortKey *models.MessageSortKey) (*models.MessageConnection, error) {
	panic("not implemented")
}
