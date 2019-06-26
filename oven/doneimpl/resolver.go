package doneimpl

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/seriousben/talk-graphql/db"
)

type Resolver struct {
	DB *db.DB
}

func (r *Resolver) Channel() ChannelResolver {
	return &channelResolver{r}
}
func (r *Resolver) Message() MessageResolver {
	return &messageResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type channelResolver struct{ *Resolver }

func (r *channelResolver) CreatedBy(ctx context.Context, obj *Channel) (*User, error) {
	uID, err := strconv.Atoi(obj.CreatedByID)
	if err != nil {
		return nil, err
	}
	us, err := r.DB.ListUsers(ctx, []int{uID})
	if err != nil {
		log.Printf("db error listing users: %#v", err)
		return nil, err
	}

	if len(us) != 1 {
		return nil, fmt.Errorf("unexpected number of users returned: %d", len(us))
	}

	u := us[0]

	return &User{
		ID:   strconv.Itoa(u.ID),
		Name: u.Name,
	}, nil
}
func (r *channelResolver) Messages(ctx context.Context, obj *Channel) ([]*Message, error) {
	cID, err := strconv.Atoi(obj.ID)
	if err != nil {
		return nil, err
	}

	ms, err := r.DB.ListMessages(ctx, cID)
	if err != nil {
		log.Printf("db error listing messages: %#v", err)
		return nil, err
	}

	gms := make([]*Message, 0, len(ms))
	for _, m := range ms {
		gms = append(gms, &Message{
			ID:          strconv.Itoa(m.ID),
			Text:        m.Text,
			CreatedByID: strconv.Itoa(m.CreatedByID),
			ChannelID:   strconv.Itoa(m.ChannelID),
		})
	}

	return gms, nil
}

type messageResolver struct{ *Resolver }

func (r *messageResolver) CreatedBy(ctx context.Context, obj *Message) (*User, error) {
	uID, err := strconv.Atoi(obj.CreatedByID)
	if err != nil {
		return nil, err
	}
	us, err := r.DB.ListUsers(ctx, []int{uID})
	if err != nil {
		log.Printf("db error listing users: %#v", err)
		return nil, err
	}

	if len(us) != 1 {
		return nil, fmt.Errorf("unexpected number of users returned: %d", len(us))
	}

	u := us[0]

	return &User{
		ID:   strconv.Itoa(u.ID),
		Name: u.Name,
	}, nil
}

func (r *messageResolver) Channel(ctx context.Context, obj *Message) (*Channel, error) {
	cID, err := strconv.Atoi(obj.ChannelID)
	if err != nil {
		return nil, err
	}
	ms, err := r.DB.ListChannelsByID(ctx, []int{cID})
	if err != nil {
		log.Printf("db error listing channel by ID: %#v", err)
		return nil, err
	}

	if len(ms) != 1 {
		return nil, fmt.Errorf("unexpected number of channels returned: %d", len(ms))
	}

	m := ms[0]

	return &Channel{
		ID:          strconv.Itoa(m.ID),
		Name:        m.Name,
		CreatedByID: strconv.Itoa(m.CreatedByID),
	}, nil

}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateMessage(ctx context.Context, input NewMessage) (*Message, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Channels(ctx context.Context) ([]*Channel, error) {
	cs, err := r.DB.ListChannels(ctx)
	if err != nil {
		log.Printf("db error listing channels: %#v", err)
		return nil, err
	}

	gcs := make([]*Channel, 0, len(cs))
	for _, c := range cs {
		gcs = append(gcs, &Channel{
			ID:          strconv.Itoa(c.ID),
			Name:        c.Name,
			CreatedByID: strconv.Itoa(c.CreatedByID),
		})
	}

	return gcs, nil
}
