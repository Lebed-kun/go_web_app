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

		tasks := task.GetTasks(Db)

		tmpl, err := template.ParseFiles("./views/pages/tasks.html")
		if err != nil {
			panic(err)
		}

		pageData := PageData{
			Title: "Tasks",
			Tasks: tasks,
		}

		tmpl.Execute(writer, pageData)
	}
	bindHandler.Handler = handler

	bindHandler.Db = Db

	return &bindHandler
}
