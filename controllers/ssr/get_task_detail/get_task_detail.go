package get_task_detail

import (
	"database/sql"
	"net/http"
	"strconv"
	"text/template"

	task "../../../models/task"

	controller "../.."
)

func GetTaskDetail(Db *sql.DB, paramSplitter func(string, *http.Request) []string) *controller.BindUrlHandler {
	bindHandler := controller.BindUrlHandler{}

	handler := func(writer http.ResponseWriter, request *http.Request) {
		type PageData struct {
			Title string
			Task  *task.Task
		}

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

		tmpl, err := template.ParseFiles("./views/pages/task_detail.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		pageData := PageData{
			Title: taskEntry.GetTitle(),
			Task:  taskEntry,
		}

		err = tmpl.Execute(writer, pageData)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	bindHandler.Handler = handler

	bindHandler.Db = Db

	return &bindHandler
}
