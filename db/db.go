package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/brianvoe/gofakeit"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const defaultDatabaseURL = "postgres://talk-graphql:pass@localhost:5432/talk-graphql?sslmode=disable"

var ddlStatements = []string{
	"DROP TABLE IF EXISTS message;",
	"DROP TABLE IF EXISTS channel;",
	"DROP TABLE IF EXISTS user_account;",
	`
CREATE TABLE IF NOT EXISTS user_account (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);`, `
CREATE TABLE IF NOT EXISTS channel (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  created_by_user_id INTEGER REFERENCES user_account(id)
);`, `
CREATE TABLE IF NOT EXISTS message (
  id SERIAL PRIMARY KEY,
  text_content TEXT NOT NULL,
  created_by_user_id INTEGER REFERENCES user_account(id),
  channel_id INTEGER REFERENCES channel(id)
);`,
}

type DB struct {
	conn *sqlx.DB
}

func Dial(connURL string) (*DB, error) {
	c, err := sqlx.Connect("postgres", connURL)
	if err != nil {
		return nil, err
	}

	for i, s := range ddlStatements {
		_, err := c.Exec(s)
		if err != nil {
			return nil, fmt.Errorf("error during ddl exec (%d): %v", i, err)
		}
	}

	return &DB{conn: c}, nil
}

func DialWithEnv() (*DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = defaultDatabaseURL
	}

	return Dial(dbURL)
}

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Channel struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	CreatedByID int    `db:"created_by_user_id"`
}

type Message struct {
	ID          int    `db:"id"`
	Text        string `db:"text_content"`
	ChannelID   int    `db:"channel_id"`
	CreatedByID int    `db:"created_by_user_id"`
}

func (db *DB) AddFixtures() error {
	gofakeit.Seed(0)
	u := User{
		Name: gofakeit.Name(),
	}
	err := db.CreateUser(context.Background(), &u)
	if err != nil {
		return err
	}

	for i := 0; i != 3; i++ {
		c := Channel{
			Name:        gofakeit.Company(),
			CreatedByID: u.ID,
		}
		err := db.CreateChannel(context.Background(), &c)
		if err != nil {
			return err
		}

		u2 := User{
			Name: gofakeit.Name(),
		}
		err = db.CreateUser(context.Background(), &u2)
		if err != nil {
			return err
		}

		for j := 0; j != 3; j++ {
			m := Message{
				Text:        gofakeit.HackerPhrase(),
				ChannelID:   c.ID,
				CreatedByID: u2.ID,
			}
			err := db.CreateMessage(context.Background(), &m)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (db *DB) CreateMessage(ctx context.Context, m *Message) error {
	table := "message"
	fields := "text_content, channel_id, created_by_user_id"
	namedFields := ":text_content, :channel_id, :created_by_user_id"

	stmt := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", table, fields, namedFields)

	rows, err := db.conn.NamedQueryContext(ctx, stmt, m)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return errors.New("no id returned")
	}

	var i int
	if err := rows.Scan(&i); err != nil {
		return err
	}
	m.ID = i
	return nil
}

func (db *DB) CreateChannel(ctx context.Context, c *Channel) error {
	table := "channel"
	fields := "name, created_by_user_id"
	namedFields := ":name, :created_by_user_id"

	stmt := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", table, fields, namedFields)

	rows, err := db.conn.NamedQueryContext(ctx, stmt, c)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return errors.New("no id returned")
	}

	var i int
	if err := rows.Scan(&i); err != nil {
		return err
	}
	c.ID = i
	return nil
}

func (db *DB) CreateUser(ctx context.Context, u *User) error {
	table := "user_account"
	fields := "name"
	namedFields := ":name"

	stmt := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", table, fields, namedFields)

	rows, err := db.conn.NamedQueryContext(ctx, stmt, u)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return errors.New("no id returned")
	}

	var i int
	if err := rows.Scan(&i); err != nil {
		return err
	}
	u.ID = i
	return nil
}

func (db *DB) ListChannels(ctx context.Context) ([]*Channel, error) {
	log.Println("db.ListChannels - SELECT id, name, created_by_user_id FROM channel")
	cs := []*Channel{}

	err := db.conn.Select(&cs, "SELECT id, name, created_by_user_id FROM channel")
	if err != nil {
		return nil, err
	}

	return cs, nil
}

func (db *DB) ListChannelsByID(ctx context.Context, ids []int) ([]*Channel, error) {
	log.Println("db.ListChannelsByID - SELECT id, name, created_by_user_id FROM channel WHERE id IN (?)", ids)
	cs := []*Channel{}

	query, args, err := sqlx.In("SELECT id, name, created_by_user_id FROM channel WHERE id IN (?)", ids)
	if err != nil {
		return nil, err
	}

	query = db.conn.Rebind(query)
	err = db.conn.Select(&cs, query, args...)
	if err != nil {
		return nil, err
	}

	return cs, nil
}

func (db *DB) ListUsers(ctx context.Context, ids []int) ([]*User, error) {
	log.Println("db.ListUsers - SELECT id, name FROM user_account WHERE id IN (?)", ids)
	us := []*User{}

	query, args, err := sqlx.In("SELECT id, name FROM user_account WHERE id IN (?)", ids)
	if err != nil {
		return nil, err
	}

	query = db.conn.Rebind(query)
	err = db.conn.Select(&us, query, args...)
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (db *DB) ListMessages(ctx context.Context, channelID int) ([]*Message, error) {
	log.Println("db.ListMessages - SELECT id, text_content, created_by_user_id, channel_id FROM message WHERE channel_id = $1", channelID)
	ms := []*Message{}

	err := db.conn.Select(&ms, "SELECT id, text_content, created_by_user_id, channel_id FROM message WHERE channel_id = $1", channelID)
	if err != nil {
		return nil, err
	}

	return ms, nil
}
