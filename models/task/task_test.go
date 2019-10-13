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
		if (*tasks[i]).Title != nil {
			fmt.Println("Title:", *tasks[i].Title)
		}
		if (*tasks[i]).Closed_at != nil {
			fmt.Println("Closed at:", *tasks[i].Closed_at)
		}
		if (*tasks[i]).Status != nil {
			fmt.Println("Status:", *tasks[i].Status)
		}
	}

	db.Close(database)
}

func TestGetTask(test *testing.T) {
	database := db.Open("sqlite3", "../../db/db.db")
	task := GetTask(database, 2)

	fmt.Println(*task)
	if (*task).Title != nil {
		fmt.Println("Title:", *task.Title)
	}
	if (*task).Closed_at != nil {
		fmt.Println("Closed at:", *task.Closed_at)
	}
	if (*task).Status != nil {
		fmt.Println("Status:", *task.Status)
	}

	db.Close(database)
}
