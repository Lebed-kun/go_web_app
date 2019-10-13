package status

import (
	"fmt"
	"testing"

	db "../../db"
	_ "github.com/mattn/go-sqlite3"
)

// Done!
func TestGetStatuses(test *testing.T) {
	database := db.Open("sqlite3", "../../db/db.db")
	statuses := GetStatuses(database)

	fmt.Println(statuses)
	for i := 0; i < 3; i++ {
		fmt.Println(*statuses[i])
		if (*statuses[i]).Description != nil {
			fmt.Println("Description:", *statuses[i].Description)
		}
	}

	db.Close(database)
}
