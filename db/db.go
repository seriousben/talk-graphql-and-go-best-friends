package db

import (
	"context"
	"errors"
	"fmt"
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
  name TEXT NOT NULL,
);`, `
CREATE TABLE IF NOT EXISTS channel (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  created_by_user_id INTEGER REFERENCES user_account(id),
);`, `
CREATE TABLE IF NOT EXISTS message (
  id SERIAL PRIMARY KEY,
  text_content TEXT NOT NULL,
  created_by_user_id INTEGER REFERENCES user_account(id),
  channel_id INTEGER REFERENCES channel(id),
);`,
}

type DB struct {
	conn *sqlx.DB
}

func Dial(connURL string) (*DB, error) {
	fmt.Println(connURL)
	c, err := sqlx.Connect("postgres", connURL)
	if err != nil {
		return nil, err
	}

	for _, s := range ddlStatements {
		_, err := c.Exec(s)
		if err != nil {
			return nil, err
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
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Message struct {
	ID        int    `db:"id"`
	Text      string `db:"text_content"`
	ChannelID int    `db:"channel_id"`
}

func (db *DB) AddFixtures() error {
	for i := 0; i != 5; i++ {
		c := Channel{
			Name: gofakeit.Company(),
		}
		err := db.CreateChannel(context.Background(), &c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) CreateMessage(ctx context.Context, m *Message) error {
	table := "message"
	fields := "text_content, channel_id"
	namedFields := ":text_content, :channel_id"

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
	fields := "name"
	namedFields := ":name"

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
	cs := []*Channel{}

	err := db.conn.Select(&cs, "SELECT id, name FROM channel")
	if err != nil {
		return nil, err
	}

	return cs, nil
}

func (db *DB) ListUsers(ctx context.Context) ([]*User, error) {
	us := []*User{}

	err := db.conn.Select(&us, "SELECT id, name FROM user_account")
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (db *DB) ListMessages(ctx context.Context, channelID string) ([]*Message, error) {
	ms := []*Message{}

	err := db.conn.Select(&ms, "SELECT id, text_content FROM message")
	if err != nil {
		return nil, err
	}

	return ms, nil
}
