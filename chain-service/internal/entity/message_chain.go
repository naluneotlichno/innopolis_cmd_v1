package entity

import "time"

type ChainStatus string

const (
	Created  ChainStatus = "created"
	Archived ChainStatus = "archived"
)

type MessageChain struct {
	ID        int         `json:"id"`
	UUID      string      `json:"uuid"`
	UserID    int         `json:"user_id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Status    ChainStatus `json:"status"`
	Title     string      `json:"title"`
}
