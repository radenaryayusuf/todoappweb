package models

import (
	"time"
)

type Task struct {
	ID       int
	Content  string
	IsDone   bool
	Assignee string
	Deadline string
	Created  time.Time
}
