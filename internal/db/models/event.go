package models

import "time"

type Event struct {
	Id          int64
	CreatedAt   time.Time
	Type        string
	Description string
	Severity    string
}
