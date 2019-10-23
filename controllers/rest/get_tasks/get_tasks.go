package get_tasks

import (
	"database/sql"
	"encoding/json"
	"net/http"

	controller "../.."
	task "../../../models/task"
)

func GetTasks(Db *sql.DB) *controller.BindUrlHandler {
	bindHandler := controller.BindUrlHandler{}

	handler := func(writer http.ResponseWriter, request *http.Request) {
		tasks, err := task.GetTasks(Db)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		jsonData, err := json.Marshal(tasks)
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
