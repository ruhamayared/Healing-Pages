package main

import "time"

// Entry represents a journal entry.
type Entry struct {
	ID        int       `json:"id"`
	Entry     string    `json:"entry"`
	CreatedAt time.Time `json:"created_at"`
}
