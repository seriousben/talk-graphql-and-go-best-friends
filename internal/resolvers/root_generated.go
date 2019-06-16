package resolvers

import (
	"github.com/seriousben/simple-graphql-chat/internal/graph"
)

type RootResolver struct{}

func (r *RootResolver) Channel() graph.ChannelResolver {
	return &ChannelResolver{r}
}
func (r *RootResolver) Message() graph.MessageResolver {
	return &MessageResolver{r}
}
func (r *RootResolver) Mutation() graph.MutationResolver {
	return &MutationResolver{r}
}
func (r *RootResolver) Query() graph.QueryResolver {
	return &QueryResolver{r}
}
func (r *RootResolver) User() graph.UserResolver {
	return &UserResolver{r}
}
