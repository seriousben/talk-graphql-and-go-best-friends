// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

type Channel struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	CreatedBy   *User  `json:"createdBy"`
	// The date and time (ISO 8601 format) when the channel was created.
	CreatedAt string `json:"createdAt"`
	// The date and time (ISO 8601 format) when the channel was last modified.
	UpdatedAt string             `json:"updatedAt"`
	Messages  *MessageConnection `json:"messages"`
	Members   *UserConnection    `json:"members"`
}

type Message struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	CreatedBy *User  `json:"createdBy"`
	// The date and time (ISO 8601 format) when the message was created.
	CreatedAt string `json:"createdAt"`
	// The date and time (ISO 8601 format) when the message was last modified.
	UpdatedAt string   `json:"updatedAt"`
	Channel   *Channel `json:"channel"`
}

type MessageConnection struct {
	// A list of edges.
	Edges []*MessageEdge `json:"edges"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
}

type MessageEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the Edge.
	Node *Message `json:"node"`
}

type NewChannel struct {
	DisplayName string `json:"displayName"`
}

type NewMessage struct {
	ChannelID string `json:"channelId"`
	Text      string `json:"text"`
}

// Information about pagination in a connection.
type PageInfo struct {
	// Indicates if there are more pages to fetch.
	HasNextPage bool `json:"hasNextPage"`
	// Indicates if there are any pages prior to the current page.
	HasPreviousPage bool `json:"hasPreviousPage"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// The date and time (ISO 8601 format) when the user was created.
	CreatedAt string `json:"createdAt"`
	// The date and time (ISO 8601 format) when the user was last modified.
	UpdatedAt string             `json:"updatedAt"`
	Messages  *MessageConnection `json:"messages"`
}

type UserConnection struct {
	// A list of edges.
	Edges []*UserEdge `json:"edges"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
}

type UserEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the Edge.
	Node *User `json:"node"`
}

// SpecifiesTheSortOrderForChannelMembers.
type ChannelMemberSortKey string

const (
	ChannelMemberSortKeyJoinedAsc  ChannelMemberSortKey = "JOINED_ASC"
	ChannelMemberSortKeyJoinedDesc ChannelMemberSortKey = "JOINED_DESC"
	ChannelMemberSortKeyNameAsc    ChannelMemberSortKey = "NAME_ASC"
	ChannelMemberSortKeyNameDesc   ChannelMemberSortKey = "NAME_DESC"
)

var AllChannelMemberSortKey = []ChannelMemberSortKey{
	ChannelMemberSortKeyJoinedAsc,
	ChannelMemberSortKeyJoinedDesc,
	ChannelMemberSortKeyNameAsc,
	ChannelMemberSortKeyNameDesc,
}

func (e ChannelMemberSortKey) IsValid() bool {
	switch e {
	case ChannelMemberSortKeyJoinedAsc, ChannelMemberSortKeyJoinedDesc, ChannelMemberSortKeyNameAsc, ChannelMemberSortKeyNameDesc:
		return true
	}
	return false
}

func (e ChannelMemberSortKey) String() string {
	return string(e)
}

func (e *ChannelMemberSortKey) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ChannelMemberSortKey(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ChannelMemberSortKey", str)
	}
	return nil
}

func (e ChannelMemberSortKey) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// SpecifiesTheSortOrderForMessages.
type MessageSortKey string

const (
	MessageSortKeyCreatedAsc  MessageSortKey = "CREATED_ASC"
	MessageSortKeyCreatedDesc MessageSortKey = "CREATED_DESC"
)

var AllMessageSortKey = []MessageSortKey{
	MessageSortKeyCreatedAsc,
	MessageSortKeyCreatedDesc,
}

func (e MessageSortKey) IsValid() bool {
	switch e {
	case MessageSortKeyCreatedAsc, MessageSortKeyCreatedDesc:
		return true
	}
	return false
}

func (e MessageSortKey) String() string {
	return string(e)
}

func (e *MessageSortKey) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MessageSortKey(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MessageSortKey", str)
	}
	return nil
}

func (e MessageSortKey) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}