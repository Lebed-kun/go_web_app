package models

import "time"

type Task struct {
	id          int
	title       *string
	description string
	starts_at   *time.Time
	closed_at   *time.Time

	status *Status
	user   *User
}
