package db

import (
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestOpen(test *testing.T) {
	db := Open("sqlite3", "./db.db")
	if db == nil {
		test.Error("Method db.Open should open database!")
	}
	fmt.Println(db)
}

func TestClose(test *testing.T) {
	db := Open("sqlite3", "./db.db")
	Close(db)
	fmt.Println(db)
}
