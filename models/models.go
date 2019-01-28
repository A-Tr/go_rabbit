package models

import (
	"time"
)

type PostMessage struct {
	Title     string    `json:"Title"`
	Subtitle  string    `json:"Subtitle"`
	Timestamp time.Time `json:"Timestamp"`
}
