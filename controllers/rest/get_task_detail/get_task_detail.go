package get_task_detail

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	task "../../../models/task"

	controller "../.."
)

func GetTaskDetail(Db *sql.DB, paramSplitter func(string, *http.Request) []string) *controller.BindUrlHandler {
	bindHandler := controller.BindUrlHandler{}

	handler := func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.ParseInt(paramSplitter(request.URL.Path, request)[0], 10, 64)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		taskEntry, err := task.GetTask(Db, id)
		if err == sql.ErrNoRows {
			http.Error(writer, err.Error(), http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(taskEntry)
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
