package models

import (
	"time"
)

type PostMessage struct {
	Title     string    `json:"Title"`
	Timestamp time.Time `json:"Timestamp"`
}
