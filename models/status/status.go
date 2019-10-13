package status

import "database/sql"

type Status struct {
	Id          int64
	Name        string
	Description *string
}

type StatusEntry struct {
	id          sql.NullInt64
	name        sql.NullString
	description sql.NullString
}

func GetStatus(Db *sql.DB, id int64) *Status {
	row := Db.QueryRow("SELECT * FROM statuses WHERE id = ?", id)
	var entry StatusEntry
	err := row.Scan(&entry.id, &entry.name, &entry.description)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}

	result := Status{
		Id:   entry.id.Int64,
		Name: entry.name.String,
	}
	if entry.description.Valid {
		result.Description = &entry.description.String
	}

	return &result
}
