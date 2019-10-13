package task

import (
	"database/sql"
	"time"

	status "../status"

	maps "../../utils/maps"
	query "../../utils/query"
)

type Task struct {
	Id          int64
	Title       *string
	Description string
	Starts_at   time.Time
	Closed_at   *time.Time

	Status *status.Status
}

type TaskEntry struct {
	id          sql.NullInt64
	title       sql.NullString
	description sql.NullString
	starts_at   sql.NullTime
	closed_at   sql.NullTime

	status_id sql.NullInt64
}

func (task *Task) getTitle() string {
	if task.Title != nil {
		return *task.Title
	} else {
		return task.Description[:20]
	}
}

func (task *Task) getShortDesc() string {
	return task.Description[:100]
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
			Id:          entry.id.Int64,
			Description: entry.description.String,
			Starts_at:   entry.starts_at.Time,
		}

		if entry.title.Valid {
			result.Title = &entry.title.String
		}

		if entry.closed_at.Valid {
			result.Closed_at = &entry.closed_at.Time
		}

		if entry.status_id.Valid {
			result.Status = status.GetStatus(Db, entry.status_id.Int64)
		}

		results = append(results, &result)
	}

	return results
}

func GetTask(Db *sql.DB, id int64) *Task {
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
		Id:          entry.id.Int64,
		Description: entry.description.String,
		Starts_at:   entry.starts_at.Time,
	}

	if entry.title.Valid {
		result.Title = &entry.title.String
	}

	if entry.closed_at.Valid {
		result.Closed_at = &entry.closed_at.Time
	}

	if entry.status_id.Valid {
		result.Status = status.GetStatus(Db, entry.status_id.Int64)
	}
	return &result
}

func CreateTask(Db *sql.DB, data map[string]interface{}) *Task {
	dataCopy := maps.Copy(data)
	// Modify status field for query
	if stat, ok := data["status"]; ok {
		dataCopy["status_id"] = stat.(*status.Status).Id
		delete(dataCopy, "status")
	}
	// Modify starts_at field for query
	dataCopy["starts_at"] = data["starts_at"].(time.Time).Format("2006-01-02")

	query, values := query.PrepareInsertQuery("tasks", dataCopy)
	result, err := Db.Exec(query, values...)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	task := Task{
		Id:          id,
		Description: data["description"].(string),
		Starts_at:   data["starts_at"].(time.Time),
	}
	if title, ok := data["title"]; ok {
		task.Title = title.(*string)
	}
	if closed_at, ok := data["closed_at"]; ok {
		task.Closed_at = closed_at.(*time.Time)
	}
	if stat, ok := data["status"]; ok {
		task.Status = stat.(*status.Status)
	}

	return &task
}

func DeleteTask(Db *sql.DB, id int64) *Task {

}
