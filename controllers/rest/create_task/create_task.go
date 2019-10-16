package create_task

import (
	"database/sql"
	"encoding/json"
	"net/http"

	task "../../../models/task"

	controller "../.."
)

func CreateTasks(Db *sql.DB) *controller.BindUrlHandler {
	handler := func(writer http.ResponseWriter, request *http.Request) {
		var data map[string]interface{}

		err := json.NewDecoder(request.Body).Decode(&data)
		if err != nil {
			panic(err)
		}

		task := task.CreateTask(Db, data)

	}
}
