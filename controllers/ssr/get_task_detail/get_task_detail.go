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
			panic(err)
		}
		taskEntry := task.GetTask(Db, id)

		tmpl, err := template.ParseFiles("./views/task_detail.html")
		if err != nil {
			panic(err)
		}

		pageData := PageData{
			Title: taskEntry.GetTitle(),
			Task:  taskEntry,
		}

		tmpl.Execute(writer, pageData)
	}
	bindHandler.Handler = handler

	bindHandler.Db = Db

	return &bindHandler
}
