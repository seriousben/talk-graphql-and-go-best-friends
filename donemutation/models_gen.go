// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package donemutation

type Channel struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	CreatedBy   *User      `json:"createdBy"`
	CreatedByID string     `json:"createdById"`
	Messages    []*Message `json:"messages"`
}

type Message struct {
	ID          string   `json:"id"`
	Text        string   `json:"text"`
	CreatedBy   *User    `json:"createdBy"`
	CreatedByID string   `json:"createdById"`
	Channel     *Channel `json:"channel"`
	ChannelID   string   `json:"channelId"`
}

type NewMessage struct {
	ChannelID string `json:"channelId"`
	Text      string `json:"text"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
