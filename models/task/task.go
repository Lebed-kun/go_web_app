package task

import (
	"database/sql"
	"time"

	status "../status"
)

type Task struct {
	id          int32
	title       *string
	description string
	starts_at   time.Time
	closed_at   *time.Time

	status *status.Status
}

type TaskEntry struct {
	id          sql.NullInt32
	title       sql.NullString
	description sql.NullString
	starts_at   sql.NullTime
	closed_at   sql.NullTime

	status_id sql.NullInt32
}

func (task *Task) getTitle() string {
	if task.title != nil {
		return *task.title
	} else {
		return task.description[:20]
	}
}

func (task *Task) getShortDesc() string {
	return task.description[:100]
}

func GetTasks(Db *sql.DB) []*Task {
	var results []*Task

	rows, err := Db.Query("SELECT * FROM tasks")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		entry := TaskEntry{}
		err = rows.Scan(&entry.id, &entry.title, &entry.description, &entry.starts_at, &entry.closed_at, &entry.status_id)
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
			result.status = status.GetStatus(Db, entry.status_id.Int32)
		} else {
			result.status = nil
		}

		results = append(results, &result)
	}

	return results
}

func GetTask(Db *sql.DB, id int32) *Task {
	row := Db.QueryRow("SELECT * FROM tasks WHERE id = ?", id)

	var entry TaskEntry
	err := row.Scan(&entry.id, &entry.title, &entry.description, &entry.starts_at, &entry.closed_at, &entry.status_id)
	if err == sql.ErrNoRows {
		return nil
	}
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
		result.status = status.GetStatus(Db, entry.status_id.Int32)
	} else {
		result.status = nil
	}
	return &result
}
