package domain

import "time"

type UserMessage struct {
	ID string `rethinkdb:"id,omitempty"`

	Messages []Message `rethinkdb:"messages"`

	CreatedAt time.Time `rethinkdb:"created_at"`
	UpdatedAt time.Time `rethinkdb:"updated_at"`
}

type Message struct {
	ID             string    `rethinkdb:"id,omitempty"`
	Text           string    `rethinkdb:"text"`
	FromUserID     string    `rethinkdb:"from_user_id"`
	ToUserID       string    `rethinkdb:"to_user_id"`
	ForwardUserID  string    `rethinkdb:"forward_user_id"`
	StatusToUserID int       `rethinkdb:"status_to_user_id"`
	CreatedAt      time.Time `rethinkdb:"created_at"`
	UpdatedAt      time.Time `rethinkdb:"updated_at"`
}
