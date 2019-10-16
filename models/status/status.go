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

func (status *Status) GetName() string {
	return status.Name
}

func GetStatuses(Db *sql.DB) ([]*Status, error) {
	var statuses []*Status

	rows, err := Db.Query("SELECT * FROM statuses")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		entry := StatusEntry{}
		err = rows.Scan(&entry.id, &entry.name, &entry.description)
		if err != nil {
			return nil, err
		}

		status := Status{
			Id:   entry.id.Int64,
			Name: entry.name.String,
		}

		if entry.description.Valid {
			status.Description = &entry.description.String
		}

		statuses = append(statuses, &status)
	}

	return statuses, nil
}

func GetStatus(Db *sql.DB, id int64) (*Status, error) {
	row := Db.QueryRow("SELECT * FROM statuses WHERE id = ?", id)
	var entry StatusEntry
	err := row.Scan(&entry.id, &entry.name, &entry.description)
	if err != nil {
		return nil, err
	}

	result := Status{
		Id:   entry.id.Int64,
		Name: entry.name.String,
	}
	if entry.description.Valid {
		result.Description = &entry.description.String
	}

	return &result, nil
}
