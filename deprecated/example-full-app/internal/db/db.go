package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var ddlStatements = []string{
	"DROP TABLE IF EXISTS message;",
	"DROP TABLE IF EXISTS channel;",
	"DROP TABLE IF EXISTS user_account;",
	`
CREATE TABLE IF NOT EXISTS user_account (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);`, `
CREATE TABLE IF NOT EXISTS channel (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  created_by_user_id INTEGER REFERENCES user_account(id),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);`, `
CREATE TABLE IF NOT EXISTS message (
  id SERIAL PRIMARY KEY,
  text_content TEXT NOT NULL,
  created_by_user_id INTEGER REFERENCES user_account(id),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
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

type Record interface {
	SetID(ID int)
	SetCreatedAt(t time.Time)
}

type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"created_at"`
}

var _ Record = &User{}

func (u *User) SetID(ID int) {
	u.ID = ID
}

func (u *User) SetCreatedAt(t time.Time) {
	u.CreatedAt = t
}

type Channel struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"created_at"`
}

var _ Record = &Channel{}

func (c *Channel) SetID(ID int) {
	c.ID = ID
}

func (c *Channel) SetCreatedAt(t time.Time) {
	c.CreatedAt = t
}

type Message struct {
	ID        int       `db:"id"`
	Text      string    `db:"text_content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"created_at"`
}

var _ Record = &Message{}

func (m *Message) SetCreatedAt(t time.Time) {
	m.CreatedAt = t
}

func (m *Message) SetID(ID int) {
	m.ID = ID
}

func (db *DB) AddFixtures() error {
	for i := 0; i != 5; i++ {
		err := db.CreateChannel(context.Background(), &Channel{
			Name: gofakeit.Company(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) doCreate(ctx context.Context, r Record) error {
	table := "channel"
	fields := "name"
	namedFields := ":name"

	stmt := fmt.Sprintf("INSERT INTO %s (%s, created_at, updated_at) VALUES (%s) RETURNING id", table, fields, namedFields)

	rows, err := db.conn.NamedQueryContext(ctx, stmt, r)
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
	r.SetID(i)
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
