package domain

import "time"

type UserConversation struct {
	ID               string             `rethinkdb:"id,omitempty"`
	UserId           string             `rethinkdb:"user_id"`
	ConversationList []ConversationList `rethinkdb:"conversation_list"`

	CreatedAt time.Time `rethinkdb:"created_at"`
	UpdatedAt time.Time `rethinkdb:"updated_at"`
}

type ConversationList struct {
	ID       string `rethinkdb:"id,omitempty"`
	Text     string `rethinkdb:"text"`
	ToUserID string `rethinkdb:"to_user_id"`

	StatusToUserID int `rethinkdb:"status_to_user_id"`

	CreatedAt time.Time `rethinkdb:"created_at"`
	UpdatedAt time.Time `rethinkdb:"updated_at"`
}
