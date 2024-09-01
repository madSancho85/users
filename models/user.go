package models

import (
	"time"
)

type User struct {
	ID            string    `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Age           int       `json:"age"`
	RecordingDate time.Time `json:"recording_date"`
}
