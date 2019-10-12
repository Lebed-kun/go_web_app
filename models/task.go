package models

import (
	"database/sql"
	"time"
)

type Task struct {
	id          int32
	title       *string
	description string
	starts_at   time.Time
	closed_at   *time.Time

	status *Status
	user   *User
}

type TaskEntry struct {
	id          sql.NullInt32
	title       sql.NullString
	description sql.NullString
	starts_at   sql.NullTime
	closed_at   sql.NullTime

	status_id sql.NullInt32
	user_id   sql.NullInt32
}

func getTasks(Db *sql.DB) []*Task {
	var results []*Task

	rows, err := Db.Query("SELECT * FROM tasks")
	if err != nil {
		panic(err)
	}

	var entry TaskEntry
	for rows.Next() {
		err = rows.Scan(&entry.id, &entry.title, &entry.description, &entry.starts_at, &entry.closed_at, &entry.status_id, &entry.user_id)
		if err != nil {
			panic(err)
		}

		result := Task{
			id:          entry.id.Int32,
			description: entry.description.String,
			starts_at:   entry.starts_at.Time,
		}

		if entry.title.Valid {
			result.title = &entry.title.String
		} else {
			result.title = nil
		}

		if entry.closed_at.Valid {
			result.closed_at = &entry.closed_at.Time
		} else {
			result.closed_at = nil
		}

		if entry.status_id.Valid {
			result.status = getStatus(Db, entry.status_id.Int32)
		} else {
			result.status = nil
		}

		if entry.user_id.Valid {
			result.user = getUser(Db, entry.user_id.Int32)
		} else {
			result.user = nil
		}

		results = append(results, &result)
	}

	return results
}
