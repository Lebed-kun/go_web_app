package get_tasks

import (
	"database/sql"
	"html/template"
	"net/http"

	controller "../.."
	task "../../../models/task"
)

func GetTasks(Db *sql.DB) *controller.BindUrlHandler {
	bindHandler := controller.BindUrlHandler{}

	handler := func(writer http.ResponseWriter, request *http.Request) {
		type PageData struct {
			Title string
			Tasks []*task.Task
		}

		tasks, err := task.GetTasks(Db)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		tmpl, err := template.ParseFiles("./views/pages/tasks.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		pageData := PageData{
			Title: "Tasks",
			Tasks: tasks,
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
