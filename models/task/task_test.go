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

func TestGetTask(test *testing.T) {
	database := db.Open("sqlite3", "../../db/db.db")
	task := GetTask(database, 2)

	fmt.Println(*task)
	if (*task).title != nil {
		fmt.Println("Title:", *task.title)
	}
	if (*task).closed_at != nil {
		fmt.Println("Closed at:", *task.closed_at)
	}
	if (*task).status != nil {
		fmt.Println("Status:", *task.status)
	}
	if (*task).user != nil {
		fmt.Println("User:", *task.user)
	}
}
