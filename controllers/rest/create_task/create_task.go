package create_task

import (
	"database/sql"
	"encoding/json"
	"net/http"

	task "../../../models/task"

	controller "../.."
)

func CreateTask(Db *sql.DB) *controller.BindUrlHandler {
	bindHandler := controller.BindUrlHandler{}

	handler := func(writer http.ResponseWriter, request *http.Request) {
		var data map[string]interface{}

		err := json.NewDecoder(request.Body).Decode(&data)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		task, err := task.CreateTask(Db, data)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(task)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-type", "application/json")
		writer.Write(jsonData)
	}
	bindHandler.Handler = handler

	bindHandler.Db = Db

	return &bindHandler
}
