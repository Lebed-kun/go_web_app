package db

import (
	"database/sql"
)

func Open(driverName string, dbPath string) *sql.DB {
	res, err := sql.Open(driverName, dbPath)
	if err != nil {
		panic(err)
	}
	return res
}

func Close(res *sql.DB) {
	err := res.Close()
	if err != nil {
		panic(err)
	}
}
