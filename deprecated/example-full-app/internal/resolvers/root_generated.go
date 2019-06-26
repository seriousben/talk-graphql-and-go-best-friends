package resolvers

import (
	"github.com/seriousben/talk-graphql/internal/db"
	"github.com/seriousben/talk-graphql/internal/graph"
)

type RootResolver struct {
	DB *db.DB
}

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
