package models

import (
	"time"
)

type PostMessage struct {
	Message   string    `json:"Message"`
	Publisher string    `json:"Publisher"`
	Timestamp time.Time `json:"Timestamp"`
}
