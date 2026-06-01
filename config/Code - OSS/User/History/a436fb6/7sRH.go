package models

import "time"

type Chat struct {
	ID         int64
	Chat_type  string
	Title      *string
	Created_by int64
}

type ChatUser struct {
	UserID   int64
	ChatID   int64
	Role     string
	JoinedAt time.Time
}
