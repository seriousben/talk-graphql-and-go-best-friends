package donegen

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateMessage(ctx context.Context, input NewMessage) (*Message, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Channels(ctx context.Context) ([]*Channel, error) {
	panic("not implemented")
}
