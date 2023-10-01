package types

import "time"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type Post struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	UserID    string    `json:"user_id"`
	PostedAt  time.Time `json:"posted_at"`
	ViewCount uint64    `json:"view_count"`
}
