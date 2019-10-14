package urls

import (
	"database/sql"
	"net/http"

	get_task_detail_ssr "../controllers/ssr/get_task_detail"
	get_tasks_ssr "../controllers/ssr/get_tasks"
)

func SetUrlHandlers(Db *sql.DB) {
	getTasksHandlerSSR := get_tasks_ssr.GetTasks(Db)
	http.HandleFunc("/", getTasksHandlerSSR.Handler)

	getTaskDetailHandlerSSR := get_task_detail_ssr.GetTaskDetail(Db, func(url string, request *http.Request) []string {
		id := request.URL.Path[len("/tasks/"):]
		params := []string{id}
		return params
	})
	http.HandleFunc("/tasks/", getTaskDetailHandlerSSR.Handler)
}
