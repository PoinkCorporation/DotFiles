package models

import "time"

type Chat struct {
	ID        int64
	ChatType  string
	CreatedBy int64
	Title     *string
}

type ChatUser struct {
	UserID   int64
	ChatID   int64
	Role     string
	JoinedAt time.Time
}
