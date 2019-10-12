package task

import (
	"fmt"
	"testing"

	db "../../db"
	_ "github.com/mattn/go-sqlite3"
)

// Done!
func TestGetTasks(test *testing.T) {
	database := db.Open("sqlite3", "../../db/db.db")
	tasks := GetTasks(database)

	if len(tasks) != 5 {
		test.Error("List of tasks should be equal to 5!")
	}

	fmt.Println(tasks)
	for i := 0; i < 5; i++ {
		fmt.Println(*tasks[i])
		if (*tasks[i]).title != nil {
			fmt.Println("Title:", *tasks[i].title)
		}
		if (*tasks[i]).closed_at != nil {
			fmt.Println("Closed at:", *tasks[i].closed_at)
		}
		if (*tasks[i]).status != nil {
			fmt.Println("Status:", *tasks[i].status)
		}
		if (*tasks[i]).user != nil {
			fmt.Println("User:", *tasks[i].user)
		}
	}
}
