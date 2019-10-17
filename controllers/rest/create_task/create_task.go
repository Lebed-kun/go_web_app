package create_task

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

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

		if starts_at, ok := data["starts_at"]; ok {
			data["starts_at"], err = time.Parse("2006-01-02", starts_at.(string))
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if closed_at, ok := data["closed_at"]; ok {
			data["closed_at"], err = time.Parse("2006-01-02", closed_at.(string))
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
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
