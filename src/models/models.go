package models

import "gorm.io/gorm"

// Entry represents a journal entry.
type Entry struct {
	gorm.Model
	Entry string `json:"entry"`
}
