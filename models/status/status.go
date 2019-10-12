package status

import "database/sql"

type Status struct {
	id          int32
	name        string
	description *string
}

type StatusEntry struct {
	id          sql.NullInt32
	name        sql.NullString
	description sql.NullString
}

func GetStatus(Db *sql.DB, id int32) *Status {
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
		id:   entry.id.Int32,
		name: entry.name.String,
	}
	if entry.description.Valid {
		result.description = &entry.description.String
	} else {
		result.description = nil
	}

	return &result
}
