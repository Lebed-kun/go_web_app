package user

import "database/sql"

type User struct {
	id   int32
	name string
	salt string

	// token *Token
}

type UserEntry struct {
	id   sql.NullInt32
	name sql.NullString
	salt sql.NullString

	token_id sql.NullInt32
}

func GetUser(Db *sql.DB, id int32) *User {
	row := Db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	var entry UserEntry
	err := row.Scan(&entry.id, &entry.name, &entry.salt, &entry.token_id)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(err)
	}

	result := User{
		id:   entry.id.Int32,
		name: entry.name.String,
		salt: entry.salt.String,
	}

	return &result
}
