package main

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	db "./db"
	urls "./urls"
)

func main() {
	database := db.Open("sqlite3", "./db/db.db")
	urls.SetUrlHandlers(database)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Server is running")
	}
}
