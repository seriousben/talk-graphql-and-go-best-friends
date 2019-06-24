package donehookwithsql

import (
	"context"
	"strconv"

	"github.com/seriousben/talk-graphql/db"
)

type Resolver struct {
	db *db.DB
}

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
	cs, err := r.db.ListChannels(ctx)
	if err != nil {
		return nil, err
	}

	gcs := make([]*Channel, 0, len(cs))
	for _, c := range cs {
		gcs = append(gcs, &donehookwithsql.Channel{
			ID:   strconv.Itoa(c.ID),
			Name: c.Name,
		})
	}

	return gcs, nil
}
